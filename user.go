package main

import (
	"fmt"
	"math/rand"

	"github.com/tjarratt/babble"
)

func GenerateUser(dataLength int) ([]string, error) {

	username := Field[string]{
		name: "username",
	}

	password := Field[string]{
		name: "password",
	}

	email := Field[string]{
		name: "email",
	}
	dob := Field[string]{
		name: "dob",
	}
	gender := Field[string]{
		name: "gender",
	}

	genders := []string{"male", "female", "other"}

	lines := make([]string, dataLength)

	// to be returned for use when making generalization data
	usernames := make([]string, dataLength)

	babbler := babble.NewBabbler()
	babbler.Count = 1

	for i := 0; i < dataLength; i++ {
		username.value = GetBabble()

		pwd, err := GetPassword()
		if err != nil {
			return nil, fmt.Errorf("generating user: %v", err)
		}
		password.value = pwd

		email.value = fmt.Sprintf("%s@%s.com", GetBabble(), GetBabble())
		dob.value = GetRandDateString(1940, 2010)
		gender.value = genders[rand.Intn(3)]

		line := fmt.Sprintf("INSERT INTO user (%s, %s, %s, %s, %s) VALUES ('%s', '%s', '%s', TO_DATE('%s'), '%s')\n",
			username.name, password.name, email.name, dob.name, gender.name,
			username.value, password.value, email.value, dob.value, gender.value)

		lines[i] = line
		usernames[i] = username.value

	}

	err := writeFile("users.sql", lines)
	if err != nil {
		return nil, fmt.Errorf("generating medication: %s", err)
	}
	return usernames, nil

}
