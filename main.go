package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lomby/gsuite/adminapi"
	"github.com/lomby/gsuite/resellerapi"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func appInfo() {
	app.Name = "Google Reseller cli"
	app.Usage = "Manage subscriptions, customers and admin for Google Resellers"
	app.Version = "1.0.0"
}

func main() {

	appInfo()

	var customerID string
	var customerData string
	var userData string
	var userKey string

	customerIDFlag := []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Usage:       "Sets the customer id (can be customerDomain or customerId)",
			Destination: &customerID,
			Required:    true,
		},
	}

	customerObjectFlag := []cli.Flag{
		&cli.StringFlag{
			Name:        "customer",
			Usage:       "customer object in ESCAPED json data (refer to Google Reseller api docs for required data)",
			Destination: &customerData,
			Required:    true,
		},
	}

	userObjectFlag := &cli.StringFlag{
		Name:        "user",
		Usage:       "user object in ESCAPED json data (refer to Google Admin SDK api docs for required data)",
		Destination: &userData,
		Required:    true,
	}

	userKeyFlag := &cli.StringFlag{
		Name:        "userKey",
		Usage:       "user key (userId or primary email address)",
		Destination: &userKey,
		Required:    true,
	}

	app.Commands = []*cli.Command{
		// Commands for Customers in Google Reseller API
		&cli.Command{
			Name:        "customer",
			Usage:       "Customer commands for Google reseller api",
			Description: "Manage google customer details",
			Subcommands: []*cli.Command{
				{
					Name:        "get",
					Usage:       "customer get --id C0*****",
					Description: "get customer details using customerId",
					Category:    "customer",
					Flags:       customerIDFlag,
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						customer := resellerapi.GetCustomer(resellerService, customerID)
						fmt.Println(customer)
						return nil
					},
				},
				{
					Name:        "create",
					Usage:       "customer create --customer {jsonData}",
					Description: "create a customer using ESCAPED json data",
					Category:    "customer",
					Flags:       customerObjectFlag,
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						customer, err := resellerapi.CreateCustomer(resellerService, []byte(customerData))

						if err != nil {
							return err
						}

						fmt.Println(customer)
						return nil
					},
				},
			},
		},
		// Commands for Subscritions in Google Reseller
		{
			Name:        "subscription",
			Usage:       "Subscription commands for Google reseller api",
			Description: "Manage google subscription details",
			Subcommands: []*cli.Command{
				{
					Name:        "get",
					Usage:       "subscription get --id CUSTOMERID",
					Description: "get subscription details using customerId",
					Category:    "subscription",
					Flags:       customerIDFlag,
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						subscription := resellerapi.FindSubscriptionByCustomerID(resellerService, customerID)
						fmt.Println(subscription)
						return nil
					},
				},
			},
		},
		// Commands for Subscritions in Google Reseller
		{
			Name:        "user",
			Usage:       "User commands for Google Admin SDK",
			Description: "Manage google users details",
			Subcommands: []*cli.Command{
				// Update a user
				{
					Name:        "update",
					Usage:       "user update --userKey USERKEY --user {jsonData}",
					Description: "update user details using the userKey and a user object",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag, userObjectFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						user, err := adminapi.UpdateUser(adminService, userKey, []byte(userData))
						if err != nil {
							return err
						}
						fmt.Println(user)
						return nil
					},
				},
				// Create a user
				{
					Name:        "create",
					Usage:       "user create --user {jsonData}",
					Description: "create a user a user object",
					Category:    "user",
					Flags:       []cli.Flag{userObjectFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						user, err := adminapi.CreateUser(adminService, []byte(userData))
						if err != nil {
							return err
						}
						fmt.Println(user)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
