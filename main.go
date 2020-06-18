package main

import (
	"context"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/reseller/v1"
)

func main() {
ctx := context.Background()

filename := "creds/credentials.json"
js, err := ioutil.ReadFile(filename)

if err != nil {
	log.Println(err)
}
credentials, err := google.JWTConfigFromJSON(js,
	reseller.AppsOrderScope,
)

if err != nil {
	log.Println(err)
}

credentials.Subject = "soletrader@reseller.soletrader.com"
client := credentials.Client(ctx)

resellerService, err := reseller.New(client)

if err != nil {
	log.Println(err)
}

// result := getCustomer(resellerService, "1stchoice-glassky.com")

file, _ := ioutil.ReadFile("subscription.json")

var subscription subscription

json.Unmarshal(file, &subscription)

fmt.Println(subscription)

newSubscription, err := createSubscription(resellerService, "C01rsqodj", subscription)

if err != nil {
	fmt.Println(err)
}

fmt.Println(newSubscription)

}