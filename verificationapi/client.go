package verificationapi

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/siteverification/v1"
)

func New() *siteverification.Service {
	ctx := context.Background()

	filename := os.Getenv("CREDENTIALS_FILE")
	js, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	credentials, err := google.JWTConfigFromJSON(js,
		siteverification.SiteverificationScope,
	)

	if err != nil {
		log.Println(err)
	}

	// user to impersonate (should be a super admin user)
	credentials.Subject = os.Getenv("CREDENTIALS_SUBJECT")
	client := credentials.Client(ctx)

	siteVerificationService, err := siteverification.New(client)

	if err != nil {
		log.Println(err)
	}

	return siteVerificationService
}
