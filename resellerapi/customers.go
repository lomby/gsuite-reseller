package resellerapi

import (
	"encoding/json"

	"google.golang.org/api/reseller/v1"
)

func GetCustomer(conn *reseller.Service, customerID string) (string, error) {

	result, err := conn.Customers.Get(customerID).Do()

	if err != nil {
		return "", nil
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func CreateCustomer(conn *reseller.Service, data []byte) (string, error) {

	var newCustomer reseller.Customer
	json.Unmarshal(data, &newCustomer)

	result, err := conn.Customers.Insert(&newCustomer).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}
