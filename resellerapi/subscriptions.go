package resellerapi

import (
	"encoding/json"
	"log"

	"google.golang.org/api/reseller/v1"
)

// Shouldn't really need this?
func ListSubscriptions(conn *reseller.Service, maxResults int64) (string, error) {

	result, err := conn.Subscriptions.List().MaxResults(maxResults).Do()

	if err != nil {
		return "", nil
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func FindSubscriptionByCustomerID(conn *reseller.Service, customerID string) reseller.Subscription {

	result, err := conn.Subscriptions.List().CustomerId(customerID).Do()

	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(result.Subscriptions[0])

	if err != nil {
		log.Println(err)
	}

	var subscription reseller.Subscription
	json.Unmarshal(js, &subscription)

	return subscription

}

func CreateSubscription(conn *reseller.Service, customerID string, data []byte) (string, error) {

	var subscription reseller.Subscription
	json.Unmarshal(data, &subscription)

	result, err := conn.Subscriptions.Insert(customerID, &subscription).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil
}

func SuspendSubscription(conn *reseller.Service, customerID string) (string, error) {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionId

	result, err := conn.Subscriptions.Suspend(customerID, subscriptionID).Do()

	if err != nil {
		return "", nil
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil
}

func ActivateSubscription(conn *reseller.Service, customerID string) (string, error) {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionId

	result, err := conn.Subscriptions.Activate(customerID, subscriptionID).Do()

	if err != nil {
		return "", nil
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil
}

func DeleteSubscription(conn *reseller.Service, customerID string) (string, error) {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionId

	err := conn.Subscriptions.Delete(customerID, subscriptionID, "transfer_to_direct").Do()

	if err != nil {
		return "", err
	}

	return "Success", nil

}
