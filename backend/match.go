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
	UIDCountryA  *datastore.Key `json:"uidCountryA" `
	UIDCountryB  *datastore.Key `json:"uidCountryB" `
	UIDWinner    *datastore.Key `json:"uidWinner" `
	UIDUser      *datastore.Key `json:"uidUser" `
	Round        int            `json:"round"`
	DateCreation time.Time      `json:"dateCreation"`
}

// List of Match
type Matchs struct {
	Matchs []Match `json:"matchs"`
}

// Type that is use for creating a Match
type MatchToCreate struct {
	UIDCountryA *datastore.Key
	UIDCountryB *datastore.Key
	UIDUser     *datastore.Key
	UIDWinner   *datastore.Key
	Round       int
}

// Create allow you to create a match
// waiting for a context and all data that are editable by the client side
// give back the current created Match or an error
func (APIMatch) Create(c endpoints.Context, r *MatchToCreate) (*Match, error) {

	k := datastore.NewIncompleteKey(c, "Match", nil)

	m := &Match{
		UIDCountryA:  r.UIDCountryA,
		UIDCountryB:  r.UIDCountryB,
		UIDUser:      r.UIDUser,
		UIDWinner:    r.UIDWinner,
		Round:        r.Round,
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
type MatchToUpdate struct {
	UID         *datastore.Key
	UIDCountryA *datastore.Key
	UIDCountryB *datastore.Key
	UIDWinner   *datastore.Key
	UIDUser     *datastore.Key
	Round       int
}

// Update allow you to update a Match
// waiting for a context and all data that are editable for the client side
// give back the updated Match or an error
func (APIMatch) Update(c endpoints.Context, r *MatchToUpdate) (*Match, error) {

	var m Match

	if err := datastore.Get(c, r.UID, &m); err == datastore.ErrNoSuchEntity {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	m.UIDCountryA = r.UIDCountryA
	m.UIDCountryB = r.UIDCountryB
	m.UIDWinner = r.UIDWinner
	m.UIDUser = r.UIDUser
	m.Round = r.Round

	_, err := datastore.Put(c, r.UID, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil

}

// type for getting a Match by UID
type MatchToGet struct {
	UID *datastore.Key
}

// Get give all data from the Match
// waiting for a context and a key
// give back a Match or an error
func (APIMatch) Get(c endpoints.Context, r *MatchToGet) (*Match, error) {

	var m Match

	if err := datastore.Get(c, r.UID, &m); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	m.UID = r.UID

	return &m, nil

}

// List give all Matchs that are stored into the datastore
// waiting for a context
// give back a list of Matchs or an error
func (APIMatch) List(c endpoints.Context) (*Matchs, error) {

	matchs := []Match{}

	keys, err := datastore.NewQuery("Match").GetAll(c, &matchs)

	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		matchs[i].UID = k
	}

	return &Matchs{matchs}, nil
}

// type for getting a Match by UID
type ListByRoundRequest struct {
	Round int
}

// List give all Matchs that are stored into the datastore
// waiting for a context
// give back a list of Matchs or an error
func (APIMatch) ListByRound(c endpoints.Context, r *ListByRoundRequest) (*Matchs, error) {

	matchs := []Match{}

	q := datastore.NewQuery("Match").Filter("Round =", r.Round)
	keys, err := q.GetAll(c, &matchs)

	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		matchs[i].UID = k
	}

	return &Matchs{matchs}, nil
}

// type for deleting a Match by UID
type MatchToDelete struct {
	UID *datastore.Key
}

// Delete allow you to delete a User
// waiging for a context and a user key
// give back an error if something went wrong
func (APIMatch) Delete(c endpoints.Context, r *MatchToDelete) error {

	err := datastore.Delete(c, r.UID)

	if err != nil {
		return nil
	}

	return err

}

// Delete allow you to delete a Match
// waiging for a context and a Match key
// give back an error if something went wrong

/*
func (APIMatch) RandomCreate(c endpoints) Matchs {

}
*/
