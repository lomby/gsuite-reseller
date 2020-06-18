package main

import (
	"encoding/json"
	"log"

	"google.golang.org/api/reseller/v1"
)

type address struct {
	ContactName      string `json: "contactName"`
	OrganizationName string `json: "organizationName"`
	AddressLine1     string `json: "addressLine1"`
	CountryCode      string `json: "countryCode"`
	PostalCode       string `json: "postcalCode"`
}

type customer struct {
	CustomerID             string  `json: "customerId"`
	CustomerDomain         string  `json: "customerDomain"`
	CustomerDomainVerified bool    `json: "customerDomainVerified"`
	AlternateEmail         string  `json: "alternateEmail"`
	PostalAddress          address `json: "postalAddress"`
}

func getCustomer(conn *reseller.Service, customerID string) customer {

	result, err := conn.Customers.Get(customerID).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
	}

	var customer customer

	json.Unmarshal(js, &customer)

	return customer

}

func createCustomer(conn *reseller.Service, customer customer) (*reseller.Customer, error) {

	js, err := json.Marshal(customer)

	var newCustomer reseller.Customer

	json.Unmarshal(js, &newCustomer)

	if err != nil {

	}

	result, err := conn.Customers.Insert(&newCustomer).Do()

	if err != nil {
		return nil, err
	}

	return result, nil

}
