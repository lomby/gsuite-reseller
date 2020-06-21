package resellerapi

import (
	"encoding/json"
	"log"

	"google.golang.org/api/reseller/v1"
)

type seats struct {
	Kind                  string `json: "kind"`
	LicensedNumberOfSeats int    `json: "licensedNumberOfSeats,omitempty"`
	MaximumNumberOfSeats  int    `json: "maximumNumberOfSeats"`
	NumberOfSeats         int    `json: "numberOfSeats,omitempty"`
}

type plan struct {
	IsCommitmentPlan bool   `json: "isCommitmentPlan,omitempty"`
	PlanName         string `json: "planName"`
}

type renewalSettings struct {
	RenewalType string `json: "renewalType,omitempty"`
}

type subscription struct {
	Kind              string          `json:"kind"`
	BillingMethod     string          `json: "billingMethod,omitempty"`
	CreationTime      string          `json: "creationTime,omitempty"`
	CustomerDomain    string          `json: "customerDomain"`
	CustomerID        string          `json: "customerId"`
	SubscriptionID    string          `json: "subscriptionId,omitempty"`
	Status            string          `json: "status,omitEmpty"`
	SuspensionReasons []string        `json:"suspensionReasons,omitempty"`
	PurchaseOrderID   string          `json:"purchaseOrderId"`
	Seats             seats           `json: "seats"`
	Plan              plan            `json: "plan"`
	RenewalSettings   renewalSettings `json: "renewalSettings,omitempty"`
}

// Shouldn't really need this?
func ListSubscriptions(conn *reseller.Service, maxResults int64) []subscription {

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

func FindSubscriptionByCustomerID(conn *reseller.Service, customerID string) subscription {

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

func CreateSubscription(conn *reseller.Service, customerID string, data []byte) (*reseller.Subscription, error) {

	var subscription reseller.Subscription
	json.Unmarshal(data, &subscription)

	result, err := conn.Subscriptions.Insert(customerID, &subscription).Do()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func SuspendSubscription(conn *reseller.Service, customerID string) subscription {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionID

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

func ActivateSubscription(conn *reseller.Service, customerID string) subscription {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionID

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

func DeleteSubscription(conn *reseller.Service, customerID string) (string, error) {

	subscriptionID := FindSubscriptionByCustomerID(conn, customerID).SubscriptionID

	err := conn.Subscriptions.Delete(customerID, subscriptionID, "transfer_to_direct").Do()

	if err != nil {
		return "", err
	}

	return "Success", nil

}
