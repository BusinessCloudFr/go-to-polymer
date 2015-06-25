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
type APIMatch struct{}

/*
 */
type Match struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	UIDPaysA     *datastore.Key `json:"uidPaysA" datastore:"-"`
	UIDPaysB     *datastore.Key `json:"uidPaysB" datastore:"-"`
	UIDWinner    *datastore.Key `json:"uidWinner" datastore:"-"`
	UIDUser      *datastore.Key `json:"uidUser" datastore:"-"`
	Round        int            `json:"round"`
	DateCreation time.Time      `json:"dateCreation"`
}

/*
 */
type Matchs struct {
	Matchs []Match `json:"matchs"`
}

/*
 */
type MatchToCreate struct {
	UIDPaysA *datastore.Key
	UIDPaysB *datastore.Key
	UIDUser  *datastore.Key
}

/*
 */
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

/*
 */
type MatchToUpgrate struct {
	UID      *datastore.Key
	UIDPaysA *datastore.Key
	UIDPaysB *datastore.Key
}

/*
 */
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
