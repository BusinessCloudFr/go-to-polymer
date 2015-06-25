package main

import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"log"
)

// Register all necessary services :
// User
// Match
// Country
func init() {

	_, err := endpoints.RegisterService(&APIUser{}, "User", "v1", "Gestionnaire des utilisateurs", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}
	_, err = endpoints.RegisterService(&APIMatch{}, "Match", "v1", "Gestionnaire des matchs", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}
	_, err = endpoints.RegisterService(&APICountry{}, "Country", "v1", "Gestionnaire des country", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}

	endpoints.HandleHTTP()
}
