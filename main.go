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
	var subscriptionData string
	var userData string
	var userKey string

	// Flags for use within all commands and subcommands

	customerIDFlag := &cli.StringFlag{
		Name:        "id",
		Usage:       "Sets the customer id (can be customerDomain or customerId)",
		Destination: &customerID,
		Required:    true,
	}

	customerObjectFlag := &cli.StringFlag{
		Name:        "customer",
		Usage:       "customer object in ESCAPED json data (refer to Google Reseller api docs for required data)",
		Destination: &customerData,
		Required:    true,
	}

	subscriptionObjectFlag := &cli.StringFlag{
		Name:        "subscription",
		Usage:       "subscription object in ESCAPED json data (refer to Google Reseller api docs for required data)",
		Destination: &subscriptionData,
		Required:    true,
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

	// Commands and Subcommands

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
					Flags:       []cli.Flag{customerIDFlag},
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
					Flags:       []cli.Flag{customerObjectFlag},
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
				// Create a subscription using a subscription object
				{
					Name:        "create",
					Usage:       "subscription create --id CUSTOMERID --subscription {jsonData}",
					Description: "create a subscription using a customerId and a subscrption Object",
					Category:    "subscription",
					Flags:       []cli.Flag{customerIDFlag, subscriptionObjectFlag},
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						subscription, err := resellerapi.CreateSubscription(resellerService, customerID, []byte(subscriptionData))
						if err != nil {
							return err
						}
						fmt.Println(subscription)
						return nil
					},
				},
				// Get a subscription using a customerId
				{
					Name:        "get",
					Usage:       "subscription get --id CUSTOMERID",
					Description: "get subscription details using customerId",
					Category:    "subscription",
					Flags:       []cli.Flag{customerIDFlag},
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						subscription := resellerapi.FindSubscriptionByCustomerID(resellerService, customerID)
						fmt.Println(subscription)
						return nil
					},
				},
				// Suspend a subscription using customerId
				{
					Name:        "suspend",
					Usage:       "subscription suspend --id CUSTOMERID",
					Description: "suspend a subscription using customerId",
					Category:    "subscription",
					Flags:       []cli.Flag{customerIDFlag},
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						subscription := resellerapi.SuspendSubscription(resellerService, customerID)
						fmt.Println(subscription)
						return nil
					},
				},
				// Activate a subscription using customerId
				{
					Name:        "activate",
					Usage:       "subscription activate --id CUSTOMERID",
					Description: "activate a subscription using customerId",
					Category:    "subscription",
					Flags:       []cli.Flag{customerIDFlag},
					Action: func(c *cli.Context) error {
						resellerService := resellerapi.New()
						subscription := resellerapi.ActivateSubscription(resellerService, customerID)
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
				// Get a user
				{
					Name:        "get",
					Usage:       "user get --userKey USERKEY",
					Description: "get a user using a userKey",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						user, err := adminapi.GetUser(adminService, userKey)
						if err != nil {
							return err
						}
						fmt.Println(user)
						return nil
					},
				},
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
