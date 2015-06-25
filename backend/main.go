package main

/*
 */
import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"log"
)

/*------------------------------------------------------------------------------------------------------------------*/

func init() {

	_, err := endpoints.RegisterService(&APIUser{}, "User", "v1", "Gestionnaire des utilisateurs", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}
	_, err = endpoints.RegisterService(&APIMatch{}, "Match", "v1", "Gestionnaire des matchs", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}
	_, err = endpoints.RegisterService(&APIPays{}, "Pays", "v1", "Gestionnaire des pays", true)
	if err != nil {
		log.Fatalf("Register service %v", err)
	}

	endpoints.HandleHTTP()
}
