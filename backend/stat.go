package main

import (
	"appengine/datastore"
	"errors"
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

	matchs := []Match{}

	keys, err := datastore.NewQuery("Match").GetAll(c, &matchs)

	if err != nil {
		return nil, err
	}

	// all matchs
	var r1 Matchs
	var r2 Matchs
	var r3 Matchs
	var r4 Matchs

	var r int // current round
	for i, k := range keys {
		matchs[i].UID = k

		r = matchs[1].Round
		switch {
		case r == 1:
			r1.Matchs = append(r1.Matchs, matchs[i])
		case r == 2:
			r2.Matchs = append(r2.Matchs, matchs[i])
		case r == 3:
			r3.Matchs = append(r3.Matchs, matchs[i])
		case r == 4:
			r4.Matchs = append(r4.Matchs, matchs[i])
		}
	}

	l1 := len(r1.Matchs)
	l2 := len(r2.Matchs)
	l3 := len(r3.Matchs)
	l4 := len(r4.Matchs)

	d1, err := distinctCountries(r1)
	d2, err := distinctCountries(r2)
	d3, err := distinctCountries(r3)
	d4, err := distinctCountries(r4)

	c.Infof("Distinct Countires : ", d1, " / len 1 : ", l1)
	c.Infof("Distinct Countires : ", d2, " / len 1 : ", l2)
	c.Infof("Distinct Countires : ", d3, " / len 1 : ", l3)
	c.Infof("Distinct Countires : ", d4, " / len 1 : ", l4)

	return &Matchs{matchs}, nil

}

func distinctCountries(m Matchs) (*map[datastore.Key]int, error) {
	if len(m.Matchs) == 0 {
		return nil, errors.New("no Match in list")
	}
	//var uids UIDs
	elem := map[datastore.Key]int{}
	for i := 0; i < len(m.Matchs); i++ {
		for j := 0; j < len(m.Matchs); j++ {
			if m.Matchs[i].UIDWinner == m.Matchs[j].UIDWinner {
				elem[*m.Matchs[i].UIDWinner]++
			}
		}
	}
	return &elem, nil
}
