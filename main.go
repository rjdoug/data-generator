package main

import (
	"fmt"
	"math/rand"
)

func main() {
	generateMedication(1)
}

// id NUMBER GENERATED ALWAYS AS IDENTITY,
// scientific_name VARCHAR2(50),
// cost NUMBER(38,2) CHECK (cost >= 0),
// isSubsidised NUMBER(1,0) DEFAULT 0 NOT NULL CHECK (isSubsidised in (0,1)),
// brand VARCHAR2(50),
// PRIMARY KEY(id)

func generateMedication(dataLength int) []string {
	var id, isSubsidised int
	var scientificName, brand string
	var cost float32

	for i := 0; i < dataLength; i++ {
		id = i
		isSubsidised = rand.Intn(1)
		scientificName = "word"
		brand = "test"
		cost = 3.45
	}

	line := fmt.Sprintf(`INSER INTO medication (id, scientific_name, cost, isSubsidised, brand) VALUES %d, %d, %s, %s, %.2f\n`,
		id, isSubsidised, scientificName, brand, cost)
	fmt.Println(line)
	return nil
}
