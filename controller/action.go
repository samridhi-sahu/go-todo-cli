package controllers

import (
	"fmt"
	"log"

	db "github.com/samridhi-sahu/go-cli/database"
	"github.com/samridhi-sahu/go-cli/utils"
)

func UpdateTasks(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide the task details")
		return
	}
	file, records, err := db.ReadFile("workbook.xlsx", "tasks")
	if err != nil {
		log.Fatal("error while getting file: ", err)
	}
	for i, record := range records {
		if record[0] == args[0] {
			record = args[1:]
			err = utils.WriteInFile(file, i+1, record)
			if err != nil {
				log.Fatal("error while updating task: ", err)
			}
			fmt.Printf("Task updated successfully with id %s", args[0])
		}
	}
}

func DeleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide the task details")
		return
	}
	file, records, err := db.ReadFile("workbook.xlsx", "tasks")
	if err != nil {
		log.Fatal("Error while getting records from file: ", err)
	}
	var updatedRecords [][]string
	for _, record := range records {
		if len(record) > 0 && record[0] != args[0] {
			updatedRecords = append(updatedRecords, record)
		}
	}
	file.DeleteSheet("tasks")
	file.NewSheet("tasks")
	for i, row := range updatedRecords {
		err := utils.WriteInFile(file, i+1, row)
		if err != nil {
			log.Fatal("Error while updating task: ", err)
		}
	}
	err = file.Save()
	if err != nil {
		log.Fatal("Error while saving file: ", err)
	}
	fmt.Println("Task deleted successfully")
}

func AddTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide the task details")
		return
	}
	file, records, err := db.ReadFile("workbook.xlsx", "tasks")
	if err != nil {
		log.Fatal("error while getting file: ", err)
	}
	nextRow := len(records) + 1
	err = utils.WriteInFile(file, nextRow, args)
	if err != nil {
		log.Fatal("Error while adding task: ", err)
	}
	fmt.Printf("Added task with id %s", args[0])
}

func GetTasks(args []string) {
	records, err := db.GetRecords("workbook.xlsx", "tasks")
	if err != nil {
		log.Fatal("Error while getting tasks from file: ", err)
	}
	if len(records) == 0 {
		fmt.Println("No task found, empty records")
	} else if len(args) == 0 {
		fmt.Println("providing all the stored task details...")
		for _, record := range records {
			fmt.Println(record)
		}
	} else if len(args) == 1 {
		fmt.Printf("providing the task details for id: %s\n", args[0])
		record := utils.GetTask(args[0], records)
		fmt.Println(record)
	} else {
		for _, id := range args {
			fmt.Printf("providing the task details for id: %s\n", id)
			record := utils.GetTask(id, records)
			fmt.Println(record)
		}
	}
}
