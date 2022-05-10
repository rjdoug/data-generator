package main

import (
	"fmt"
	"math/rand"
)

func GenerateMedication(dataLength int) error {

	isSubsidised := Field[int]{
		name: "isSubsidised",
	}

	scientificName := Field[string]{
		name: "scientific_name",
	}

	brand := Field[string]{
		name: "brand",
	}

	cost := Field[float64]{
		name: "cost",
	}

	lines := make([]string, dataLength)

	const minMedicationCost = 5.0
	const maxMedicationCost = 100.0

	for i := 0; i < dataLength; i++ {
		isSubsidised.value = rand.Intn(2)
		scientificName.value = GetBabble(1)
		brand.value = GetBabble(1)
		cost.value = minMedicationCost + rand.Float64()*(maxMedicationCost-minMedicationCost)

		line := fmt.Sprintf("INSERT INTO medication (%s, %s, %s, %s) VALUES (%d, '%s', '%s', %.2f);\n",
			isSubsidised.name, scientificName.name, brand.name, cost.name,
			isSubsidised.value, scientificName.value, brand.value, cost.value)

		lines[i] = line
	}

	err := writeFile("sql_scripts/medication.sql", lines)
	if err != nil {
		return fmt.Errorf("generating medication: %s", err)
	}
	return nil
}
