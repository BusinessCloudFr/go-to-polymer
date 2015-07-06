package main

import (
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"time"
)

// API for Match
type APIStat struct{}

type Stat struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	Countries    *Countries     `json:"countries"`
	Results      []float64
	Round        int       `json:"round"`
	DateCreation time.Time `json:"dateCreation"`
}

// Set create Stats
func (APIStat) Set() {

}

// Get retrive all vots
func (APIStat) Get(c endpoints.Context) (*Matchs, error) {
	q := datastore.NewQuery("Match")

	var ms Matchs

	_, err := q.GetAll(c, &ms)

	return &ms, err
}
