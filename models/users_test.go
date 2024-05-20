package models

import "testing"

func TestGetUserByID(t *testing.T) {
	// Arrange
	expect := &User{
		ID:        1,
		FirstName: "Carlos",
		LastName:  "Jafet",
	}
	users = []*User{expect}

	// Act
	user, err := GetUserByID(expect.ID)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if user != *expect {
		t.Errorf("Got user '%v', expect '%v", user, expect)

	}
}

func TestDeleteUserByID(t *testing.T) {
	// Arrange
	user := &User{
		ID:        1,
		FirstName: "Carlos",
		LastName:  "Jafet",
	}
	expect := &User{}

	users = []*User{user}

	// Act
	err := DeleteUserByID(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	u, e := GetUserByID(user.ID)
	if e == nil {
		t.Fatal(err)
	}

	// Assert
	if u != *expect {
		t.Errorf("Got user '%v', expect '%v", u, expect)

	}

}

func TestGetUserByIDTableDriven(t *testing.T) {
	scenarios := []struct {
		input  int
		expect User
	}{
		{input: 1, expect: User{ID: 1, FirstName: "Carlos", LastName: "Jafet"}},
		{input: 2, expect: User{ID: 2, FirstName: "Carlos", LastName: "Neto"}},
		{input: 7777, expect: User{}},
	}

	for _, s := range scenarios {
		got, err := GetUserByID(s.input)
		if err != nil {
			t.Logf("Fail getting request for userID: %v", s.input)
		}
		if got != s.expect {
			t.Errorf("Got %v, but was expecting %v", got, s.expect)
		}
	}
}
