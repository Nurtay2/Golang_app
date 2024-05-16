package model

import "time"

// Fighter represents the UFC fighters table model
type Fighter struct {
  ID          int     `json:"id"`
  Name        string  `json:"name"`
  WeightClass string  `json:"weight_class"`
  Reach       float64 `json:"reach"`
  Wins        int     `json:"wins"`
  Losses      int     `json:"losses"`
}

// FightingStyle represents the UFC fighting_styles table model
type FightingStyle struct {
  ID          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}

// Match represents the UFC matches table model
type Match struct {
  ID              int           `json:"id"`
  Date            time.Time     `json:"date"`
  Duration        time.Duration `json:"duration"`
  WinnerFighterID int           `json:"winner_fighter_id"`
}

// Gym represents the UFC gyms table model
type Gym struct {
  ID       int    `json:"id"`
  Name     string `json:"name"`
  Location string `json:"location"`
}

// Promotion represents the UFC promotions table model
type Promotion struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  FounderID int    `json:"founder_id"`
}
