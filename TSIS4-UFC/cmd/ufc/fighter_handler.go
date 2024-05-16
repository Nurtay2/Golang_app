package main

import (
    "errors"
    "log"
    "net/http"

    "github.com/Nurtay2/TSIS2-UFC/pkg/ufc/model"
    "github.com/Nurtay2/TSIS2-UFC/pkg/ufc/validator"


    
)

func (app *application) createFighterHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name          string  `json:"name"`
        WeightClass   string  `json:"weight_class"`
        Reach         float64 `json:"reach"`
        Wins          int     `json:"wins"`
        Losses        int     `json:"losses"`
    }

    err := app.readJSON(w, r, &input)
    if err != nil {
        log.Println(err)
        app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
        return
    }

    fighter := &model.Fighter{
        Name:        input.Name,
        WeightClass: input.WeightClass,
        Reach:       input.Reach,
        Wins:        input.Wins,
        Losses:      input.Losses,
    }

    err = app.models.Fighters.Insert(fighter)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    app.writeJSON(w, http.StatusCreated, envelope{"fighter": fighter}, nil)
}

func (app *application) getFightersList(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name               string
        ReachFrom          int
        ReachTo            int
        model.Filters
    }
    v := validator.New()
    qs := r.URL.Query()

    // Use our helpers to extract the name and reach range query string values.
    input.Name = app.readStrings(qs, "name", "")
    input.ReachFrom = app.readInt(qs, "reachFrom", 0, v)
    input.ReachTo = app.readInt(qs, "reachTo", 0, v)

    // Extract other query string values like page, page_size, and sort.
    input.Filters.Page = app.readInt(qs, "page", 1, v)
    input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
    input.Filters.Sort = app.readStrings(qs, "sort", "id")

    // Add the supported sort values for this endpoint to the sort safelist.
    input.Filters.SortSafeList = []string{
        // ascending sort values
        "id", "name", "reach",
        // descending sort values
        "-id", "-name", "-reach",
    }

    if model.ValidateFilters(v, input.Filters); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    fighters, metadata, err := app.models.Fighters.GetAll(input.Name, input.ReachFrom, input.ReachTo, input.Filters)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    app.writeJSON(w, http.StatusOK, envelope{"fighters": fighters, "metadata": metadata}, nil)
}

func (app *application) getFighterHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    fighter, err := app.models.Fighters.Get(id)
    if err != nil {
        switch {
        case errors.Is(err, model.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    app.writeJSON(w, http.StatusOK, envelope{"fighter": fighter}, nil)
}

func (app *application) updateFighterHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    fighter, err := app.models.Fighters.Get(id)
    if err != nil {
        switch {
        case errors.Is(err, model.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    var input struct {
        Name        *string  `json:"name"`
        WeightClass *string  `json:"weight_class"`
        Reach       *float64 `json:"reach"`
        Wins        *int     `json:"wins"`
        Losses      *int     `json:"losses"`
    }

    err = app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }

    if input.Name != nil {
        fighter.Name = *input.Name
    }

    if input.WeightClass != nil {
        fighter.WeightClass = *input.WeightClass
    }

    if input.Reach != nil {
        fighter.Reach = *input.Reach
    }

    if input.Wins != nil {
        fighter.Wins = *input.Wins
    }

    if input.Losses != nil {
        fighter.Losses = *input.Losses
    }

    v := validator.New()

    if model.ValidateFighter(v, fighter); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    err = app.models.Fighters.Update(fighter)
    if err != nil {
        switch {
        case errors.Is(err, model.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    app.writeJSON(w, http.StatusOK, envelope{"fighter": fighter}, nil)
}

func (app *application) deleteFighterHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }

    err = app.models.Fighters.Delete(id)
    if err != nil {
        switch {
        case errors.Is(err, model.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
    }

    app.writeJSON(w, http.StatusOK, envelope{"message": "success"}, nil)
}
