package logic

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

// CsvToXlsx reads a CSV file and creates a corresponding xlsx file
func CsvToXlsx(csvFilePath string, xlsxFilePath string) error {
	// Open the CSV file
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Print(err)
		return err
	}
	defer csvFile.Close()

	// Read the CSV file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Print("Error reading CSV file: ", err)
		return err
	}

	// Create a new xlsx file
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("Sheet1")
	if err != nil {
		log.Print("Error creating sheet: ", err)
		return err
	}

	// Write the records to the xlsx file
	for _, record := range records {
		row := sheet.AddRow()
		for _, value := range record {
			cell := row.AddCell()
			cell.Value = value
		}
	}

	// Save the xlsx file
	err = xlsxFile.Save(xlsxFilePath)
	if err != nil {
		log.Print("Error saving xlsx file: ", err)
		return err
	}

	return nil
}
