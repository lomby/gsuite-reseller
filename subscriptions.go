package main

import (
	"encoding/json"
	"log"

	"google.golang.org/api/reseller/v1"
)

type seats struct {
	LicensedNumberOfSeats int `json: "licensedNumberOfSeats"`
}

type subscription struct {
	BillingMethod     string   `json: "billingMethod"`
	CreationTime      string   `json: "creationTime"`
	CustomerDomain    string   `json: "customerDomain"`
	CustomerID        string   `json: "customerId"`
	SubscriptionID    string   `json: "subscriptionId"`
	Status            string   `json: "status"`
	SuspensionReasons []string `json:"suspensionReasons"`
	Seats             seats    `json: "seats"`
}

// Shouldn't really need this?
func listSubscriptions(conn *reseller.Service, maxResults int64) []subscription {

	result, err := conn.Subscriptions.List().MaxResults(maxResults).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result.Subscriptions)

	if err != nil {
		log.Println(err)
	}

	var subscriptions []subscription

	json.Unmarshal(js, &subscriptions)

	return subscriptions

}

func findSubscriptionByCustomerID(conn *reseller.Service, customerID string) subscription {

	result, err := conn.Subscriptions.List().CustomerId(customerID).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result.Subscriptions[0])

	if err != nil {
		log.Println(err)
	}

	var subscription subscription

	json.Unmarshal(js, &subscription)

	return subscription

}

func suspendSubscription(conn *reseller.Service, customerID string) subscription {

	subscriptionID := findSubscriptionByCustomerID(conn, customerID).SubscriptionID

	result, err := conn.Subscriptions.Suspend(customerID, subscriptionID).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
	}

	var subscription subscription

	json.Unmarshal(js, &subscription)

	return subscription
}

func activateSubscription(conn *reseller.Service, customerID string) subscription {

	subscriptionID := findSubscriptionByCustomerID(conn, customerID).SubscriptionID

	result, err := conn.Subscriptions.Activate(customerID, subscriptionID).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
	}

	var subscription subscription

	json.Unmarshal(js, &subscription)

	return subscription
}
