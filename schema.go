package main

import (
	"time"

	bandit "github.com/alextanhongpin/go-bandit"
	"github.com/rs/xid"
)

// Arm represents a choice in the multi-arm bandit experiment
type Arm struct {
	Arm            int       `json:"arm"` // Don't omit empty - 0 is a valid arm
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Reward         float64   `json:"reward,omitempty"`
	Feature        string    `json:"feature"`
	IsActionTaken  bool      `json:"-"`
	IsCompleted    bool      `json:"-"`
	ExperimentName string    `json:"-"`
	ExperimentID   string    `json:"-"`
}

func NewGuid() string {
	return xid.New().String()
}

func NewUTCDate() time.Time {
	return time.Now().UTC()
}

// NewArm returns a new bandit
func NewArm(arm int) *Arm {
	return &Arm{
		Arm:       arm,
		ID:        NewGuid(),
		CreatedAt: NewUTCDate(),
		UpdatedAt: NewUTCDate(),
	}
}

func NewDefaultEpsilonGreedy() (*bandit.EpsilonGreedy, error) {
	return bandit.NewEpsilonGreedy(0.1, nil, nil)
}

type ResponseError struct {
	Error string `json:"message"`
	Code  int    `json:"code"`
}

type GetStatsResponse struct {
	Counts         []int     `json:"counts"`
	Rewards        []float64 `json:"rewards"`
	Features       []string  `json:"features"`
	ExperimentName string    `json:"experimentName"`
}
