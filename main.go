package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/tjarratt/babble"
)

func main() {
	generateMedication(10)
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

	var lines []string

	babbler := babble.NewBabbler()
	babbler.Count = 1

	for i := 1; i <= dataLength; i++ {
		id = i
		isSubsidised = rand.Intn(2)
		scientificName = babbler.Babble()
		brand = babbler.Babble()
		cost = 3.45

		line := fmt.Sprintf("INSER INTO medication (id, scientific_name, cost, isSubsidised, brand) VALUES (%d, %d, %s, %s, %.2f)\n",
			id, isSubsidised, scientificName, brand, cost)

		lines = append(lines, line)
	}

	fmt.Print(lines)
	return nil
}

func writeFile(file string, lines []string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("")
	}
}
