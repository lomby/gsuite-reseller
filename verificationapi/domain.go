package verificationapi

import "google.golang.org/api/siteverification/v1"

func GetToken(conn *siteverification.Service, domain string) (*siteverification.SiteVerificationWebResourceGettokenResponse, error) {

site := siteverification.SiteVerificationWebResourceGettokenRequestSite{
	Type: "INET_DOMAIN",
	Identifier: domain,
}

request := siteverification.SiteVerificationWebResourceGettokenRequest{
	Site: &site,
	VerificationMethod: "DNS_TXT",
}

result, err := siteverification.NewWebResourceService(conn).GetToken(&request).Do()

if err != nil {
	return nil, err
}

return result, nil

}

func Verify(conn *siteverification.Service, domain string) (*siteverification.SiteVerificationWebResourceResource, error) {


site := siteverification.SiteVerificationWebResourceResourceSite{
	Type: "INET_DOMAIN",
	Identifier: domain,
}

request := siteverification.SiteVerificationWebResourceResource{
	Site: &site,
}

result, err := siteverification.NewWebResourceService(conn).Insert("DNS_TXT", &request).Do()

if err != nil {
	return nil, err
}

return result, nil

}