package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func WriteInFile(file *excelize.File, nextRow int, task []string) error {
	for i, v := range task {
		cell := fmt.Sprintf("%c%d", 'A'+i, nextRow)
		err := file.SetCellValue("tasks", cell, v)
		if err != nil {
			return fmt.Errorf("error while setting cell value: %s", err)
		}
	}
	err := file.Save()
	if err != nil {
		return fmt.Errorf("error while saving file: %s", err)
	}
	return nil
}

func GetTask(id string, records [][]string) []string {
	for _, record := range records {
		if record[0] == id {
			return record
		}
	}
	return []string{}
}
