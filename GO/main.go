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
	idlCombi, totalPoints, matchScore := logic.GetIdealConfig(entities.Config{
		Power:      335,
		Aero:       355,
		LigtWeight: 305,
		Grip:       375,
		Cost:       53,
	})
	if idlCombi == nil {
		fmt.Println("No Ideal Combination Found")
		return
	}

	fmt.Println("Total Points: ", totalPoints+(28+36+31))
	fmt.Println("Match Score: ", matchScore)
	fmt.Println("Ideal Combination: ", idlCombi.CombiCode)
	fmt.Println("Power: ", idlCombi.Power)
	fmt.Println("Aero: ", idlCombi.Aero)
	fmt.Println("LigtWeight: ", idlCombi.LigtWeight)
	fmt.Println("Grip: ", idlCombi.Grip)
	fmt.Println("Cost: ", idlCombi.Cost)
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

	err = logic.CsvToXlsx("results.csv", "results.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
