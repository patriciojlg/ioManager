package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/xuri/excelize/v2"
)

func parserByteDatoTo(arrayBytesData [][]byte) ([]map[string]interface{}, error) {
	var allMaps []map[string]interface{}

	for _, item := range arrayBytesData {
		var obj map[string]interface{}
		if err := json.Unmarshal(item, &obj); err != nil {
			log.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}
		allMaps = append(allMaps, obj)
	}
	return allMaps, nil
}

func setColumns(f *excelize.File, allMaps []map[string]interface{}) []string {
	keySet := make(map[string]struct{})
	for _, obj := range allMaps {
		for key := range obj {
			keySet[key] = struct{}{}
		}
	}
	var keys []string
	for k := range keySet {
		keys = append(keys, k)
	}

	sort.Strings(keys) // Opcional: ordena columnas alfabéticamente

	for i, key := range keys {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Sheet1", cell, key)
	}
	return keys
}

func setData(f *excelize.File, allMaps []map[string]interface{}, keys []string) {
	for i, obj := range allMaps {
		for j, key := range keys {
			if val, ok := obj[key]; ok {
				cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
				f.SetCellValue("Sheet1", cell, val)
			}
		}
	}
}

// 👉 Nueva versión
func ImplodeOutputToXlsx(rawData [][]byte) ([]byte, error) {
	f := excelize.NewFile()

	allMaps, err := parserByteDatoTo(rawData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON data: %w", err)
	}

	keys := setColumns(f, allMaps)
	setData(f, allMaps, keys)

	// 🧠 Aquí viene el cambio: guardar en memoria, no en disco
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("error writing Excel to buffer: %w", err)
	}

	return buffer.Bytes(), nil
}
