package verificationapi

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/siteverification/v1"
)

func New() *siteverification.Service {
	ctx := context.Background()

	filename := "creds/credentials.json"
	js, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	credentials, err := google.JWTConfigFromJSON(js,
		siteverification.SiteverificationScope,
	)

	if err != nil {
		log.Println(err)
	}

	credentials.Subject = "soletrader@reseller.soletrader.com"
	client := credentials.Client(ctx)

	siteVerificationService, err := siteverification.New(client)

	if err != nil {
		log.Println(err)
	}

	return siteVerificationService
}
