package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lomby/gsuite/adminapi"
	"github.com/lomby/gsuite/resellerapi"
	"github.com/lomby/gsuite/verificationapi"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func appInfo() {
	app.Name = "Google Reseller cli"
	app.Usage = "Manage subscriptions, customers and admin for Google Resellers"
	app.Version = "1.0.0"
}

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appInfo()

	var customerID string
	var customerData string
	var subscriptionData string
	var userData string
	var userKey string
	var userAlias string
	var verificationDomain string

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

	userAliasFlag := &cli.StringFlag{
		Name:        "userAlias",
		Usage:       "alias to add to a user",
		Destination: &userAlias,
		Required:    true,
	}

	verificationDomainFlag := &cli.StringFlag{
		Name:        "domain",
		Usage:       "domain to verify",
		Destination: &verificationDomain,
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
				// Make user super admin
				{
					Name:        "make-admin",
					Usage:       "user make-admin --userKey USERKEY",
					Description: "make user a super admin",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						status, err := adminapi.MakeUserAdmin(adminService, userKey)
						if err != nil {
							return err
						}
						fmt.Println(status)
						return nil
					},
				},
				// Make user super admin
				{
					Name:        "delete",
					Usage:       "user delete --userKey USERKEY",
					Description: "delete a user",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						status, err := adminapi.DeleteUser(adminService, userKey)
						if err != nil {
							return err
						}
						fmt.Println(status)
						return nil
					},
				},
				// Create an alias for a user
				{
					Name:        "add-alias",
					Usage:       "user add-alias --userKey USERKEY --userAlias ALIAS (e.g. alias@primary-domain.com)",
					Description: "add a user alias",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag, userAliasFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						result, err := adminapi.CreateUserAlias(adminService, userKey, userAlias)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				// Delete an alias for a user
				{
					Name:        "delete-alias",
					Usage:       "user delete-alias --userKey USERKEY --userAlias ALIAS (e.g. alias@primary-domain.com)",
					Description: "add a user alias",
					Category:    "user",
					Flags:       []cli.Flag{userKeyFlag, userAliasFlag},
					Action: func(c *cli.Context) error {
						adminService := adminapi.New()
						result, err := adminapi.DeleteUserAlias(adminService, userKey, userAlias)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
			},
		},
		// Commands for Site Verification API
		{
			Name:        "verification",
			Usage:       "Verify commands for site verification api",
			Description: "verification for domains",
			Subcommands: []*cli.Command{
				// Get the verification token for use with verifying the domain
				{
					Name:        "get-token",
					Usage:       "verification get-token --domain DOMAIN",
					Description: "Get token for domain verification",
					Category:    "verification",
					Flags:       []cli.Flag{verificationDomainFlag},
					Action: func(c *cli.Context) error {
						verificationService := verificationapi.New()
						verify, err := verificationapi.GetToken(verificationService, verificationDomain)
						if err != nil {
							return err
						}
						fmt.Println(verify)
						return nil
					},
				},
				// Attempt to verify a domain
				{
					Name:        "verify",
					Usage:       "verification verify --domain DOMAIN",
					Description: "Attempt to verify domain after verification steps are completed",
					Category:    "verification",
					Flags:       []cli.Flag{verificationDomainFlag},
					Action: func(c *cli.Context) error {
						verificationService := verificationapi.New()
						verify, err := verificationapi.Verify(verificationService, verificationDomain)
						if err != nil {
							return err
						}
						fmt.Println(verify)
						return nil
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
