package db

import (
	"fmt"
	"log"
	"os"

	exl "github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	*exl.File
}

func ReadFile(fileName, sheetName string) (*exl.File, [][]string, error) {
	file, err := GetFile(fileName, sheetName)
	if err != nil {
		return nil, nil, fmt.Errorf("error while getting file: %s", err)
	}
	defer file.Close()
	records, err := file.GetRows(sheetName)
	if err != nil {
		return nil, nil, fmt.Errorf("error while getting rows: %s", err)
	}
	return file, records, nil
}

func GetRecords(fileName, sheetName string) ([][]string, error) {
	file, err := GetFile(fileName, sheetName)
	if err != nil {
		return nil, fmt.Errorf("error while getting file: %s", err)
	}
	defer file.Close()
	return file.GetRows(sheetName)
}

// Check if a sheet exists
func (file *ExcelFile) SheetExists(sheetName string) bool {
	index, err := file.GetSheetIndex(sheetName)
	if err != nil {
		log.Fatal("error while getting sheet index: ", err)
	}
	return index >= 0
}

func (file *ExcelFile) GetSheet(sheetName string) error {
	if !file.SheetExists(sheetName) {
		_, err := file.NewSheet(sheetName)
		if err != nil {
			return fmt.Errorf("error while creating sheet: %s", err)
		}
		err = file.Save()
		if err != nil {
			return fmt.Errorf("error while saving file: %s", err)
		}
		fmt.Println("sheet created successfully")
	}
	return nil
}

// Check if a file exists
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}

// open file if exist else create
func OpenFile(fileName string) (*exl.File, error) {
	if FileExists(fileName) {
		file := exl.NewFile()
		err := file.SaveAs(fileName)
		if err != nil {
			return nil, fmt.Errorf("error while saving file: %s", err)
		}
		fmt.Println("file created successfully")
		return file, nil
	}
	return exl.OpenFile(fileName)
}

func GetFile(fileName, sheetName string) (*exl.File, error) {
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error while opening file: %s", err)
	}
	exFile := &ExcelFile{file}
	err = exFile.GetSheet(sheetName)
	if err != nil {
		return nil, fmt.Errorf("error while getting sheet: %s", err)
	}
	return file, nil
}
