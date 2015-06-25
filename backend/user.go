/*
 */
package main

/*
 */
import (
	//"appengine"
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"log"
	//"net/url"
	"time"
)

/*
 */
type APIUser struct{}

/*
 */
type User struct {
	UID             *datastore.Key `json:"uid" datastore:"-"`
	Pseudo          string         `json:"pseudo"`
	DateInscription time.Time      `json:"dateInscription"`
}

/*
 */
type Users struct {
	Users []User `json:"Users"`
}

/*
 */
type UserToCreate struct {
	Pseudo string
}

/*
 */
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

/*
 */
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

/*
 */
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

/*
 */
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

/*
 */
func (APIUser) Delete(c endpoints.Context, r *UserUID) error {

	err := datastore.Delete(c, r.UID)

	if err != nil {
		return nil
	}

	return err

}
