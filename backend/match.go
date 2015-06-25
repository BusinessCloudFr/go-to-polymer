package main

import (
	//"appengine"
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"log"
	//"net/url"
	"time"
)

// API for Match
type APIMatch struct{}

// Type of Match
type Match struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	UIDPaysA     *datastore.Key `json:"uidPaysA" datastore:"-"`
	UIDPaysB     *datastore.Key `json:"uidPaysB" datastore:"-"`
	UIDWinner    *datastore.Key `json:"uidWinner" datastore:"-"`
	UIDUser      *datastore.Key `json:"uidUser" datastore:"-"`
	Round        int            `json:"round"`
	DateCreation time.Time      `json:"dateCreation"`
}

// List of Match
type Matchs struct {
	Matchs []Match `json:"matchs"`
}

// Type that is use for creating a Match
type MatchToCreate struct {
	UIDPaysA *datastore.Key
	UIDPaysB *datastore.Key
	UIDUser  *datastore.Key
}

// Create allow you to create a match
// waiting for a context and all data that are editable by the client side
// give back the current created Match or an error
func (APIMatch) Create(c endpoints.Context, r *MatchToCreate) (*Match, error) {

	k := datastore.NewIncompleteKey(c, "Match", nil)

	m := &Match{
		UIDPaysA:     r.UIDPaysA,
		UIDPaysB:     r.UIDPaysB,
		UIDUser:      r.UIDUser,
		DateCreation: time.Now(),
	}

	k, err := datastore.Put(c, k, m)

	if err != nil {
		return nil, err
	}

	m.UID = k

	return m, nil
}

// type for upgrading a Match
type MatchToUpgrate struct {
	UID      *datastore.Key
	UIDPaysA *datastore.Key
	UIDPaysB *datastore.Key
}

// Update allow you to update a Match
// waiting for a context and all data that are editable for the client side
// give back the updated Match or an error
func (APIMatch) Update(c endpoints.Context, r *MatchToUpgrate) (*Match, error) {

	var m Match

	if err := datastore.Get(c, r.UID, &m); err == datastore.ErrNoSuchEntity {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	m.UIDPaysA = r.UIDPaysA
	m.UIDPaysB = r.UIDPaysB

	_, err := datastore.Put(c, r.UID, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil

}

/*
func (APIMatch) RandomCreate(c endpoints) Matchs {

}
*/
