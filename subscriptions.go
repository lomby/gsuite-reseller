package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/api/reseller/v1"
)

type seats struct {
	LicensedNumberOfSeats int `json: "licensedNumberOfSeats,omitempty"`
	MaximumNumberOfSeats  int `json: "maximumNumberOfSeats"`
	NumberOfSeats         int `json: "numberOfSeats"`
}

type plan struct {
	IsCommitmentPlan bool   `json: "isCommitmentPlan"`
	PlanName         string `json: "planName"`
}

type renewalSettings struct {
	RenewalType string `json: "renewalType"`
}

type subscription struct {
	BillingMethod     string          `json: "billingMethod,omitempty"`
	CreationTime      string          `json: "creationTime,omitempty"`
	CustomerDomain    string          `json: "customerDomain"`
	CustomerID        string          `json: "customerId"`
	SubscriptionID    string          `json: "subscriptionId,omitempty"`
	Status            string          `json: "status,omitEmpty"`
	SuspensionReasons []string        `json:"suspensionReasons,omitempty"`
	Seats             seats           `json: "seats"`
	SKUID             string          `json: "skuId"`
	Plan              plan            `json: "plan"`
	RenewalSettings   renewalSettings `json: "renewalSettings"`
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

func createSubscription(conn *reseller.Service, customerID string, subscription subscription) (*reseller.Subscription, error) {

	js, err := json.Marshal(subscription)

	if err != nil {

	}

	var newSubscription reseller.Subscription

	json.Unmarshal(js, &newSubscription)

	fmt.Println(newSubscription)

	result, err := conn.Subscriptions.Insert(customerID, &newSubscription).Do()

	if err != nil {
		return nil, err
	}

	return result, nil
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
