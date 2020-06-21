package adminapi

import (
	"encoding/json"
	"fmt"

	admin "google.golang.org/api/admin/directory/v1"
)

func CreateUser(conn *admin.Service, data []byte) (string, error) {

	var newUser admin.User
	json.Unmarshal(data, &newUser)

	result, err := conn.Users.Insert(&newUser).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func UpdateUser(conn *admin.Service, userKey string, data []byte) (string, error) {

	fmt.Println(string(data))

	var user admin.User
	json.Unmarshal(data, &user)

	result, err := conn.Users.Update(userKey, &user).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func GetUser(conn *admin.Service, userKey string) (string, error) {

	result, err := conn.Users.Get(userKey).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func MakeUserAdmin(conn *admin.Service, userKey string) (string, error) {

	var makeAdmin admin.UserMakeAdmin
	makeAdmin.Status = true

	err := conn.Users.MakeAdmin(userKey, &makeAdmin).Do()

	if err != nil {
		return "", err
	}

	return "Success", nil
}

func DeleteUser(conn *admin.Service, userKey string) (string, error) {

	err := conn.Users.Delete(userKey).Do()

	if err != nil {
		return "", err
	}

	return "Successfully Deleted", nil
}

func CreateUserAlias(conn *admin.Service, userKey string, alias string) (string, error) {

	newAlias := &admin.Alias{
		Alias: alias,
	}

	result, err := conn.Users.Aliases.Insert(userKey, newAlias).Do()

	if err != nil {
		return "", err
	}

	toJSON, err := json.MarshalIndent(result, "", " ")

	return string(toJSON), nil

}

func DeleteUserAlias(conn *admin.Service, userKey string, alias string) (string, error) {

	err := conn.Users.Aliases.Delete(userKey, alias).Do()

	if err != nil {
		return "", err
	}

	return "Success", nil

}
