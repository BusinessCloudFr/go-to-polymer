package main

// API for Match
type APIStat struct{}

type Stat struct {
	UID *datastore.Key `json:"uid" datastore:"-"`

	Countrys
	Round        int       `json:"round"`
	DateCreation time.Time `json:"dateCreation"`
}

// Get retrive all vots
// get
func (APIStat) Get(r int) (stat, error) {

}
