package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

/*
by working with pointers we are going to be able to manipulate each users from
various places within the application without having to copy the user around
At package level I also don't need to provide the colon operator to get the implicit initialization
*/

var (
	// users  []*User
	nextID = 1
)

var users = []*User{
	{
		ID:        1,
		FirstName: "Carlos",
		LastName:  "Jafet",
	},
	{
		ID:        2,
		FirstName: "Carlos",
		LastName:  "Neto",
	},
}

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User already exists")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", u.ID)
}

func DeleteUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' not found", id)
}
