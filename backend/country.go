package main

import (
	//"appengine"
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	//"log"
	//"net/url"
	"time"
)

// API for Country
type APICountry struct{}

// Country type
type Country struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	Label        string         `json:"label"`
	UrlFlag      string         `json:"urlflag"`
	IsoCode      string         `json:"isocode"`
	DateCreation time.Time      `json:"dateCreation"`
}

// Goupe of Country type
type Countrys struct {
	Countrys []Country `json:"countries"`
}

// Type that are used for creating a country
type CountryToCreate struct {
	Label   string
	UrlFlag string
	IsoCode string
}

// Create allow you to create an new Country
// waiting for a context and all datas for creating a country
// give back a Country or an error
func (APICountry) Create(c endpoints.Context, r *CountryToCreate) (*Country, error) {

	k := datastore.NewIncompleteKey(c, "Country", nil)

	p := &Country{
		Label:        r.Label,
		UrlFlag:      r.UrlFlag,
		IsoCode:      r.IsoCode,
		DateCreation: time.Now(),
	}

	k, err := datastore.Put(c, k, p)

	if err != nil {
		return nil, err
	}

	p.UID = k

	return p, nil
}

type CountryUID struct {
	UID *datastore.Key `json:"uid"`
}

// Get let you get all data form a Country with the Country key
// waiging for a context and a key
// give back a Country or an error
func (APICountry) Get(c endpoints.Context, r *CountryUID) (*Country, error) {

	var country Country

	err := datastore.Get(c, r.UID, &country)

	if err != nil {
		return nil, err
	}

	country.UID = r.UID

	return &country, nil

}

type CountryIso struct {
	IsoCode string
}

// GetbyIso allow you to get a country with it's ISO code
// waiting for a context and an ISO code
// give back a Country or an error
func (APICountry) GetbyIso(c endpoints.Context, r *CountryIso) (*Country, error) {

	countries := []Country{}

	keys, err := datastore.NewQuery("Country").Filter("IsoCode =", r.IsoCode).GetAll(c, &countries)

	if err != nil || len(keys) != 1 {
		return nil, err
	}

	var country Country

	if err := datastore.Get(c, keys[0], &country); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	country.UID = keys[0]

	return &country, nil
}

// Type that is used for editing the Country type
type CountryToEdit struct {
	UID     *datastore.Key
	Label   string
	UrlFlag string
	IsoCode string
}

// Edit allow to edit a Country
// waiting for a context and a all editables infos form the Country and the Country key
// give back the current modified Country or an error
func (APICountry) Edit(c endpoints.Context, r *CountryToEdit) (*Country, error) {

	var country Country

	if err := datastore.Get(c, r.UID, &country); err == datastore.ErrNoSuchEntity {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	country.Label = r.Label
	country.UrlFlag = r.UrlFlag
	country.IsoCode = r.IsoCode

	k, err := datastore.Put(c, r.UID, &country)

	if err != nil {
		return nil, err
	}

	country.UID = k

	return &country, nil
}

// List let you list all Counties that are stored into the datastore
// waiging for a context
// give back a list of Countries or an error
func (APICountry) List(c endpoints.Context) (*Countrys, error) {

	countries := []Country{}
	keys, err := datastore.NewQuery("Country").GetAll(c, &countries)

	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		countries[i].UID = k
	}

	return &Countrys{countries}, nil
}

// Delete allow you to delete a Country
// waiging for a context and a Country key
// give back an error if something went wrong
func (APICountry) Delete(c endpoints.Context, r *CountryUID) error {

	err := datastore.Delete(c, r.UID)

	if err != nil {
		return nil
	}

	return err

}
