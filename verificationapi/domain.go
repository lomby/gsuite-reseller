package verificationapi

import (
	"encoding/json"

	"google.golang.org/api/siteverification/v1"
)

func GetToken(conn *siteverification.Service, domain string) (string, error) {

	site := siteverification.SiteVerificationWebResourceGettokenRequestSite{
		Type:       "INET_DOMAIN",
		Identifier: domain,
	}

	request := siteverification.SiteVerificationWebResourceGettokenRequest{
		Site:               &site,
		VerificationMethod: "DNS_TXT",
	}

	result, err := siteverification.NewWebResourceService(conn).GetToken(&request).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func Verify(conn *siteverification.Service, domain string) (string, error) {

	site := siteverification.SiteVerificationWebResourceResourceSite{
		Type:       "INET_DOMAIN",
		Identifier: domain,
	}

	request := siteverification.SiteVerificationWebResourceResource{
		Site: &site,
	}

	result, err := siteverification.NewWebResourceService(conn).Insert("DNS_TXT", &request).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}
