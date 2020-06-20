package adminapi

import (
	"encoding/json"
	"fmt"

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

func CreateUser(conn *admin.Service, data []byte) (*admin.User, error) {

	var newUser admin.User
	json.Unmarshal(data, &newUser)

	result, err := conn.Users.Insert(&newUser).Do()

	if err != nil {
		return nil, err
	}

	return result, nil

}

func UpdateUser(conn *admin.Service, userKey string, data []byte) (*admin.User, error) {

	fmt.Println(string(data))

	var user admin.User
	json.Unmarshal(data, &user)

	result, err := conn.Users.Update(userKey, &user).Do()

	if err != nil {
		return nil, err
	}

	return result, nil

}

func GetUser(conn *admin.Service, userKey string) (*admin.User, error) {

	result, err := conn.Users.Get(userKey).Do()

	if err != nil {
		return nil, err
	}

	return result, nil

}
