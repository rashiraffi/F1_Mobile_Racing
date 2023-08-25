package main

import (
	"encoding/csv"
	"f1/internal/entities"
	"f1/internal/logic"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	tim := time.Now()
	idlCombi, totalPoints := logic.GetIdealConfig(entities.Config{
		Power:      340,
		Aero:       320,
		LigtWeight: 320,
		Grip:       340,
		Cost:       44,
	})
	if idlCombi == nil {
		fmt.Println("No Ideal Combination Found")
		return
	}

	fmt.Println("Total Points: ", totalPoints+(26+28+33))
	fmt.Println("Ideal Combination: ", idlCombi)
	fmt.Println("Time: ", time.Since(tim))

	file, err := os.OpenFile("results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := []string{fmt.Sprintf("%.0f", totalPoints), fmt.Sprintf("%d", int(idlCombi.Power)), fmt.Sprintf("%d", int(idlCombi.Aero)), fmt.Sprintf("%d", int(idlCombi.LigtWeight)), fmt.Sprintf("%d", int(idlCombi.Grip)), fmt.Sprintf("%d", int(idlCombi.Cost))}
	err = writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
