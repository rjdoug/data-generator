package main

import (
	"fmt"
	"math/rand"
)

func GeneratePractioners(usernames []string) error {

	username := Field[string]{
		name: "username",
	}

	role := Field[string]{
		name: "role",
	}

	hourlySalary := Field[float64]{
		name: "hourly_salary",
	}

	roles := []string{"gp", "dermitologist", "radiologist", "pyschiatrist", "neurologist"}

	lines := make([]string, len(usernames))

	for idx, user := range usernames {
		username.value = user

		role.value = roles[rand.Intn(len(roles))]

		hourlySalary.value = rand.Float64() * 100

		line := fmt.Sprintf("INSERT INTO medical_practitioner (%s, %s, %s) VALUES ('%s', '%s', %.2f)\n",
			username.name, role.name, hourlySalary.name,
			username.value, role.value, hourlySalary.value)

		lines[idx] = line

	}

	err := writeFile("sql_scripts/practitioner.sql", lines)
	if err != nil {
		return fmt.Errorf("generating patient: %v", err)
	}

	return nil

}
