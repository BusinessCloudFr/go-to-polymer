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
type APIPays struct{}

/*
 */
type Pays struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	Label        string         `json:"label"`
	UrlFlag      string         `json:"urlflag"`
	IsoCode      string         `json:"isocode"`
	DateCreation time.Time      `json:"dateCreation"`
}

/*
 */
type Payss struct {
	Payss []Pays `json:"payss"`
}

/*
 */
type PaysToCreate struct {
	Label   string
	UrlFlag string
	IsoCode string
}

/*
 */
func (APIPays) Create(c endpoints.Context, r *PaysToCreate) (*Pays, error) {

	k := datastore.NewIncompleteKey(c, "Pays", nil)

	p := &Pays{
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

type PaysUID struct {
	UID *datastore.Key
}

/*
 */
func (APIPays) Get(c endpoints.Context, r *PaysUID) (*Pays, error) {

	var pays Pays

	if err := datastore.Get(c, r.UID, &pays); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	pays.UID = r.UID

	return &pays, nil

}

type PaysIso struct {
	IsoCode string
}

/*
 */
func (APIPays) GetbyIso(c endpoints.Context, r *PaysIso) (*Pays, error) {

	payss := []Pays{}

	keys, err := datastore.NewQuery("Pays").Filter("IsoCode =", r.IsoCode).GetAll(c, &payss)

	if err != nil || len(keys) != 1 {
		return nil, err
	}

	var pays Pays

	if err := datastore.Get(c, keys[0], &pays); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("User not found")
	} else if err != nil {
		return nil, err
	}

	pays.UID = keys[0]

	return &pays, nil
}

/*
 */
type PaysToEdit struct {
	UID     *datastore.Key
	Label   string
	UrlFlag string
	IsoCode string
}

/*
 */
func (APIPays) Edit(c endpoints.Context, r *PaysToEdit) (*Pays, error) {

	var pays Pays

	if err := datastore.Get(c, r.UID, &pays); err == datastore.ErrNoSuchEntity {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	pays.Label = r.Label
	pays.UrlFlag = r.UrlFlag
	pays.IsoCode = r.IsoCode

	k, err := datastore.Put(c, r.UID, &pays)

	if err != nil {
		return nil, err
	}

	pays.UID = k

	return &pays, nil
}

/*
 */
func (APIPays) List(c endpoints.Context) (*Payss, error) {

	payss := []Pays{}
	keys, err := datastore.NewQuery("Pays").GetAll(c, &payss)

	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		payss[i].UID = k
	}

	return &Payss{payss}, nil
}
