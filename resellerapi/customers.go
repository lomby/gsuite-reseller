package resellerapi

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

func GetCustomer(conn *reseller.Service, customerID string) customer {

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

func CreateCustomer(conn *reseller.Service, data []byte) (*reseller.Customer, error) {

	var newCustomer reseller.Customer
	json.Unmarshal(data, &newCustomer)

	result, err := conn.Customers.Insert(&newCustomer).Do()

	if err != nil {
		return nil, err
	}

	return result, nil

}