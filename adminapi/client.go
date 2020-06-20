package adminapi

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
)

// Establish a connetion to the Gogle Admin API
func New() *admin.Service {

	ctx := context.Background()

	filename := "creds/credentials.json"
	js, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	credentials, err := google.JWTConfigFromJSON(js,
		admin.AdminDirectoryUserScope,
		admin.AdminDirectoryGroupScope,
	)

	if err != nil {
		log.Println(err)
	}

	// user to impersonate
	credentials.Subject = "soletrader@reseller.soletrader.com"
	client := credentials.Client(ctx)

	adminService, err := admin.New(client)

	if err != nil {
		log.Println(err)
	}

	return adminService
}
