package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ioManager/models"
	"log"

	"github.com/xuri/excelize/v2"
)

func ReadExcelFromBytes(data []byte) ([][]string, error) {
	// Open the Excel file from the byte slice
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get all the rows in the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func RowsToArrayJsonStruct(rows [][]string) ([]models.JSONFile, error) {
	headers := rows[0]
	var result []models.JSONFile

	for i, row := range rows[1:] {
		if len(row) == 0 || row[0] == "" {
			log.Printf("Saltando fila %d sin identificador", i+2)
			continue
		}

		record := make(map[string]string)
		for j, val := range row {
			if j < len(headers) {
				record[headers[j]] = val
			}
		}

		jsonData, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Printf("Error convirtiendo fila %d a JSON: %v", i+2, err)
			continue
		}

		result = append(result, models.JSONFile{
			Filename: fmt.Sprintf("%s.json", row[0]),
			Data:     jsonData,
		})
	}

	return result, nil
}
