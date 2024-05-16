package model

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/Nurtay2/TSIS2-UFC/pkg/ufc/validator"
)

// UFC Fighters
type Fighter struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    WeightClass string  `json:"weight_class"`
    Reach       float64 `json:"reach"`
    Wins        int     `json:"wins"`
    Losses      int     `json:"losses"`
}

type FighterModel struct {
    DB       *sql.DB
    InfoLog  *log.Logger
    ErrorLog *log.Logger
}

func (m FighterModel) GetAll(name string, from, to int, filters Filters) ([]*Fighter, Metadata, error) {

    // Retrieve all fighters from the database.
    query := fmt.Sprintf(
        `
        SELECT count(*) OVER(), id, name, weight_class, reach, wins, losses
        FROM fighters
        WHERE (LOWER(name) = LOWER($1) OR $1 = '')
        AND (reach >= $2 OR $2 = 0)
        AND (reach <= $3 OR $3 = 0)
        ORDER BY %s %s, id ASC
        LIMIT $4 OFFSET $5
        `,
        filters.sortColumn(), filters.sortDirection())

    // Create a context with a 3-second timeout.
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // Organize our four placeholder parameter values in a slice.
    args := []interface{}{name, from, to, filters.limit(), filters.offset()}

    // Use QueryContext to execute the query. This returns a sql.Rows result set containing
    // the result.
    rows, err := m.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, Metadata{}, err
    }

    // Importantly, defer a call to rows.Close() to ensure that the result set is closed
    // before GetAll returns.
    defer func() {
        if err := rows.Close(); err != nil {
            m.ErrorLog.Println(err)
        }
    }()

    // Declare a totalRecords variable
    totalRecords := 0

    var fighters []*Fighter
    for rows.Next() {
        var fighter Fighter
        err := rows.Scan(&totalRecords, &fighter.ID, &fighter.Name, &fighter.WeightClass, &fighter.Reach, &fighter.Wins, &fighter.Losses)
        if err != nil {
            return nil, Metadata{}, err
        }

        // Add the Fighter struct to the slice
        fighters = append(fighters, &fighter)
    }

    // When the rows.Next() loop has finished, call rows.Err() to retrieve any error
    // that was encountered during the iteration.
    if err = rows.Err(); err != nil {
        return nil, Metadata{}, err
    }

    // Generate a Metadata struct, passing in the total record count and pagination parameters
    // from the client.
    metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

    // If everything went OK, then return the slice of the fighters and metadata.
    return fighters, metadata, nil
}

func (m FighterModel) Insert(fighter *Fighter) error {
    // Insert a new fighter into the database.
    query := `
        INSERT INTO fighters (name, weight_class, reach, wins, losses) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id
        `
    args := []interface{}{fighter.Name, fighter.WeightClass, fighter.Reach, fighter.Wins, fighter.Losses}
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    return m.DB.QueryRowContext(ctx, query, args...).Scan(&fighter.ID)
}

func (m FighterModel) Get(id int) (*Fighter, error) {
    // Return an error if the ID is less than 1.
    if id < 1 {
        return nil, ErrRecordNotFound
    }
    // Retrieve a specific fighter based on its ID.
    query := `
        SELECT id, name, weight_class, reach, wins, losses
        FROM fighters
        WHERE id = $1
        `
    var fighter Fighter
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    row := m.DB.QueryRowContext(ctx, query, id)
    err := row.Scan(&fighter.ID, &fighter.Name, &fighter.WeightClass, &fighter.Reach, &fighter.Wins, &fighter.Losses)
    if err != nil {
        return nil, fmt.Errorf("cannot retrieve fighter with id: %v, %w", id, err)
    }
    return &fighter, nil
}

func (m FighterModel) Update(fighter *Fighter) error {
    // Update a specific fighter in the database.
    query := `
        UPDATE fighters
        SET name = $1, weight_class = $2, reach = $3, wins = $4, losses = $5
        WHERE id = $6
        `
    args := []interface{}{fighter.Name, fighter.WeightClass, fighter.Reach, fighter.Wins, fighter.Losses, fighter.ID}
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := m.DB.ExecContext(ctx, query, args...)
    return err
}

func (m FighterModel) Delete(id int) error {
    // Return an error if the ID is less than 1.
    if id < 1 {
        return ErrRecordNotFound
    }

    // Delete a specific fighter from the database.
    query := `
        DELETE FROM fighters
        WHERE id = $1
        `
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := m.DB.ExecContext(ctx, query, id)
    return err
}

func ValidateFighter(v *validator.Validator, fighter *Fighter) {
    // Check if the name field is empty.
    v.Check(fighter.Name != "", "name", "must be provided")
    // Add more validation rules as needed.
}
