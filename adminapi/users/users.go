package adminapi

import (
	"encoding/json"
	"log"

	admin "google.golang.org/api/admin/directory/v1"
)

type name struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

type user struct {
	Name                      name   `json: "name"`
	Password                  string `json: "password"`
	PrimaryEmail              string `json: "primaryEmail"`
	ChangePasswordAtNextLogin bool   `json: "changePasswordAtNextLogin"`
}

func CreateUser(conn *admin.Service, data []byte) *admin.User {

	var newUser admin.User
	json.Unmarshal(data, &newUser)

	result, err := conn.Users.Insert(&newUser).Do()

	if err != nil {
		log.Println(err)
	}

	return result

}

func UpdateUser(conn *admin.Service, data []byte) *admin.User {

	var user admin.User
	json.Unmarshal(data, &user)

	result, err := conn.Users.Update(&user)

	if err != nil {
		log.Println(err)
	}

	return result

}

func GetUser(conn *admin.Service, userKey string) *admin.User {

	result, err := conn.Users.Get(userKey).Do()

	if err != nil {
		log.Println(err)
	}

	return result

}
