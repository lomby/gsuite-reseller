package resellerapi

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/reseller/v1"
)

// Establish a connection to the Google Reseller API
func New() *reseller.Service {
	ctx := context.Background()

	filename := "creds/credentials.json"
	js, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	credentials, err := google.JWTConfigFromJSON(js,
		reseller.AppsOrderScope,
	)

	if err != nil {
		log.Println(err)
	}

	credentials.Subject = "soletrader@reseller.soletrader.com"
	client := credentials.Client(ctx)

	resellerService, err := reseller.New(client)

	if err != nil {
		log.Println(err)
	}

	return resellerService
}
