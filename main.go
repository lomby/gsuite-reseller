package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/reseller/v1"
)

// Establish a connection to the Google Reseller API
func newResellerService() *reseller.Service {
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

// Establish a connetion to the Gogle Admin API
func newAdminService() *admin.Service {
	ctx := context.Background()

	filename := "creds/credentials.json"
	js, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	credentials, err := google.JWTConfigFromJSON(js,
		admin.CloudPlatformScope,
	)

	if err != nil {
		log.Println(err)
	}

	// credentials.Subject = "soletrader@reseller.soletrader.com"
	client := credentials.Client(ctx)

	adminService, err := admin.New(client)

	if err != nil {
		log.Println(err)
	}

	return adminService
}

func main() {

	// resellerService := newResellerService()
	adminService := newAdminService()

	// result := getCustomer(resellerService, "CUSTOMER_ID_HERE")

	fmt.Println(adminService)

}
