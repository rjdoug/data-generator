package main

import (
	"fmt"

	"github.com/icrowley/fake"
	"github.com/lucasjones/reggen"
)

// CREATE TABLE PATIENT(
// 	username VARCHAR(30) NOT NULL,
// 	NHI VARCHAR(7) NOT NULL,
// 	address VARCHAR(40) NOT NULL,
// 	CONSTRAINT constraintFOrmatNHI CHECK(REGEXP_LIKE(NHI, ‘^[A-Z]{3}\d{4}$’))
// );
//

func GeneratePatients(usernames []string) error {

	username := Field[string]{
		name: "username",
	}

	nhi := Field[string]{
		name: "nhi",
	}

	address := Field[string]{
		name: "address",
	}

	lines := make([]string, len(usernames))

	for idx, user := range usernames {
		username.value = user

		var err error
		nhi.value, err = reggen.Generate("[A-Z]{3}\\d{4}", 0)
		if err != nil {
			return fmt.Errorf("generating patient: %v", err)
		}

		address.value = fake.StreetAddress()

		line := fmt.Sprintf("INSERT INTO patient (%s, %s, %s) VALUES ('%s', '%s', '%s')\n",
			username.name, nhi.name, address.name,
			username.value, nhi.value, address.value)

		lines[idx] = line
	}

	err := writeFile("sql_scripts/patient.sql", lines)
	if err != nil {
		return fmt.Errorf("generating patient: %v", err)
	}

	return nil
}
