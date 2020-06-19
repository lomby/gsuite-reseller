package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
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

func main() {

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	// resellerService := newResellerService()
	// adminService := newAdminService()

	// result := getCustomer(resellerService, "CUSTOMER_ID_HERE")

	// data, _ := ioutil.ReadFile("user.json")
	// newUser := adminapi.CreateUser(adminService, data)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
