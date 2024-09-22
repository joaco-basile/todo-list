package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Todo struct {
	Title       string
	Description string
	Id          int
	IsComplete  bool
}

type Todos []Todo

func WriteTodo(args []string) {
	newTodo := Todo{Title: args[0], Description: args[1], Id: genereId(), IsComplete: false}

	file := openFile()
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(todoDecoder(newTodo))
	checkError(err)
}

func ReadTodos() Todos {
	file := openFile()
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkError(err)
	if records == nil {
		return nil
	}

	var todos Todos
	for _, record := range records {
		todos = append(todos, todoEncoder(record))
	}

	return todos
}

func UpdateTodo(id int, title, description string) {
	newTodo := Todo{Title: title, Description: description, Id: id, IsComplete: false}
	todoAsRecord := todoDecoder(newTodo)

	file := openFile()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkError(err)

	for key, record := range records {
		recordEncode := todoEncoder(record)
		if recordEncode.Id == id {
			records = slices.Delete(records, key, key+1)
			records = slices.Insert(records, key, todoAsRecord)
			break
		}
	}

	file.Close()
	fmt.Println(records)
	file, err = os.Create("prueba.csv")
	defer file.Close()
	checkError(err)
	writer := csv.NewWriter(file)
	writer.WriteAll(records)
}

func DeleteTodo(id int) {
	file := openFile()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	checkError(err)

	var newRecords [][]string

	for k, record := range records {
		record := todoEncoder(record)
		if record.Id == id {
			newRecords = append(records[:k], records[k+1:]...)
			break
		}
	}

	err = file.Close()
	checkError(err)

	file, err = os.Create("prueba.csv")
	checkError(err)

	fmt.Println(newRecords)
	writer := csv.NewWriter(file)
	writer.WriteAll(newRecords)
	defer writer.Flush()
}

func openFile() *os.File {
	file, err := os.OpenFile("prueba.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	checkError(err)
	return file
}

func genereId() int {
	file := openFile()
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkError(err)
	if records == nil {
		return 1
	}

	lastTodo := todoEncoder(records[len(records)-1])
	return lastTodo.Id + 1
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func todoEncoder(record []string) Todo {
	id, err := strconv.Atoi(record[2])
	if err != nil {
		log.Fatal(err)
	}

	completed, err := strconv.ParseBool(record[3])
	if err != nil {
		log.Fatal(err)
	}

	encodeTodo := Todo{
		Title:       record[0],
		Description: record[1],
		Id:          id,
		IsComplete:  completed,
	}
	return encodeTodo
}

func todoDecoder(t Todo) []string {
	records := []string{
		t.Title,
		t.Description,
		strconv.FormatInt(int64(t.Id), 10),
		strconv.FormatBool(t.IsComplete),
	}
	return records
}
