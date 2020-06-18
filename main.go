package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/reseller/v1"
)

func main() {
ctx := context.Background()

filename := "creds/credentials.json"
json, err := ioutil.ReadFile(filename)

if err != nil {
	log.Println(err)
}
credentials, err := google.JWTConfigFromJSON(json,
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

result := getCustomer(resellerService, "1stchoice-glassky.com")

fmt.Println(result)

}