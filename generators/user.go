package generators

import (
	"fmt"
	"math/rand"

	"github.com/BlaviButcher/data-generator/generators/helper"
	"github.com/BlaviButcher/data-generator/io"
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

	for i := 0; i < dataLength; i++ {
		username.value = helper.GetBabble(1)

		pwd, err := helper.GetPassword()
		if err != nil {
			return nil, fmt.Errorf("generating user: %v", err)
		}
		password.value = pwd

		email.value = fmt.Sprintf("%s@%s.com", helper.GetBabble(1), helper.GetBabble(1))
		randDateTime := helper.GetRandDateString(1940, 2010)
		dob.value = randDateTime[:10]

		gender.value = genders[rand.Intn(3)]

		line := fmt.Sprintf("INSERT INTO users (%s, %s, %s, %s, %s) VALUES ('%s', '%s', '%s', TO_DATE('%s', 'YYYY/MM/DD'), '%s');\n",
			username.name, password.name, email.name, dob.name, gender.name,
			username.value, password.value, email.value, dob.value, gender.value)

		lines[i] = line
		usernames[i] = username.value

	}

	err := io.WriteFile("sql_scripts/users.sql", lines)
	if err != nil {
		return nil, fmt.Errorf("generating medication: %s", err)
	}
	return usernames, nil

}
