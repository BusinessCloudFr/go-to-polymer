package main

import (
	//"appengine"
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"log"
	//"net/url"
	"time"
)

// API for User
type APIUser struct{}

// User type
type User struct {
	UID             *datastore.Key `json:"uid" datastore:"-"`
	Pseudo          string         `json:"pseudo"`
	DateInscription time.Time      `json:"dateInscription"`
}

// Goupe of User type
type Users struct {
	Users []User `json:"Users"`
}

// type that is use for the request
type UserToCreate struct {
	Pseudo string
}

// Create allow you to create new User and save them into the datastore
// Waiting for the context and Pseudo
// give back a User or an error
func (APIUser) Create(c endpoints.Context, r *UserToCreate) (*User, error) {

	k := datastore.NewIncompleteKey(c, "User", nil)

	u := &User{
		Pseudo:          r.Pseudo,
		DateInscription: time.Now(),
	}

	k, err := datastore.Put(c, k, u)

	if err != nil {
		return nil, err
	}
	u.UID = k

	return u, nil
}

type UserPseudo struct {
	Pseudo string
}

// GetbyPseudo is used for checking if the user is already in the database
// waiting for a context and a pseudo
// give back a User or an error
func (APIUser) GetbyPseudo(c endpoints.Context, r *UserPseudo) (*User, error) {

	users := []User{}

	keys, err := datastore.NewQuery("User").Filter("Pseudo =", r.Pseudo).GetAll(c, &users)

	if err != nil || len(keys) != 1 {
		return nil, err
	}

	var user User

	if err := datastore.Get(c, keys[0], &user); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	user.UID = keys[0]

	return &user, nil
}

type UserUID struct {
	UID *datastore.Key
}

// Get give all data from the User
// waiting for a context and a key
// give back a User or an error
func (APIUser) Get(c endpoints.Context, r *UserUID) (*User, error) {

	var user User

	if err := datastore.Get(c, r.UID, &user); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	user.UID = r.UID

	return &user, nil

}

// List give all Users that are stored into the datastore
// waiting for a context
// give back a list of Users or an error
func (APIUser) List(c endpoints.Context) (*Users, error) {

	users := []User{}

	keys, err := datastore.NewQuery("User").GetAll(c, &users)

	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		users[i].UID = k
	}

	return &Users{users}, nil
}

// Delete allow you to delete a User
// waiging for a context and a user key
// give back an error if something went wrong
func (APIUser) Delete(c endpoints.Context, r *UserUID) error {

	err := datastore.Delete(c, r.UID)

	if err != nil {
		return nil
	}

	return err

}
