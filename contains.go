// CREATE TABLE contains(
// id NUMBER GENERATED ALWAYS AS IDENTITY,
// perscription_id NUMBER(20) NOT NULL,
// medication_id NUMBER(20) NOT NULL,
// dosage_mg NUMBER(20,2),
// instruction VARCHAR2(2000),
// repeats NUMBER(20),
// PRIMARY KEY(id, perscription_id, medication_id)
// --    FOREIGN KEY(perscription_id)
// --    REFERENCES perscription(id),
// --    FOREIGN KEY(medication_id)
// --    REFERENCES medication(medication_id)
// );
//

package main

import (
	"fmt"
	"math/rand"
)

func GenerateContains(dataLength, perscriptionLength, medicationLength int) error {

	const maxMG = 1000
	const maxRepeats = 5

	perscriptionID := Field[int]{
		name: "perscription_id",
	}

	medicationID := Field[int]{
		name: "medication_id",
	}

	dosageMG := Field[int]{
		name: "dosage_mg",
	}

	instruction := Field[string]{
		name: "instruction",
	}

	repeats := Field[int]{
		name: "repeats",
	}

	lines := make([]string, dataLength)

	for i := 0; i < dataLength; i++ {
		perscriptionID.value = rand.Intn(perscriptionLength)
		medicationID.value = rand.Intn(medicationLength)
		dosageMG.value = rand.Intn(maxMG)
		instruction.value = GetBabble(10)
		repeats.value = rand.Intn(maxRepeats)

		line := fmt.Sprintf("INSERT INTO contains (%s, %s, %s, %s, %s) VALUES (%d, %d, %d, '%s', %d)\n",
			perscriptionID.name, medicationID.name, dosageMG.name, instruction.name, repeats.name,
			perscriptionID.value, medicationID.value, dosageMG.value, instruction.value, repeats.value)

		lines[i] = line
	}

	err := writeFile("sql_scripts/contains.sql", lines)
	if err != nil {
		return fmt.Errorf("generating contains relationship: %s", err)
	}
	return nil

}
