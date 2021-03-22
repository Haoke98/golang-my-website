package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Account struct {
	Id       uint64
	Name     string
	Username string
	Password string
	Url      string
	Tel      string
	Email    string
}

var CsvDataFileName string = "accounts.csv"
var IdMax uint64 = 0
var AccountById map[uint64]*Account = make(map[uint64]*Account)
var AccountByName map[string]*Account = make(map[string]*Account)

func (account *Account) store() {
	AccountById[account.Id] = account
	AccountByName[account.Name] = account
}
func (account *Account) save() {
	IdMax++
	account.Id = IdMax
	account.store()
	csvDataFile, err := os.Create(CsvDataFileName)
	if err != nil {
		panic(err)
	}
	defer csvDataFile.Close()
	writer := csv.NewWriter(csvDataFile)
	for _, value := range AccountById {
		line := []string{strconv.Itoa(int(IdMax)), strconv.Itoa(int(value.Id)), value.Name, value.Username, value.Password, value.Tel, value.Email, value.Url}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	csvDataFile.Close()
}

func load() {
	file, err := os.Open(CsvDataFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	for true {
		row, err := reader.Read()
		if err == nil {
			IdMax, err = strconv.ParseUint(row[0], 0, 64)
			if err != nil {
				panic(err)
			} else {
				id, err := strconv.ParseUint(row[1], 0, 64)
				if err != nil {
					panic(err)
				} else {
					account := Account{Id: id, Name: row[2], Username: row[3], Password: row[4], Tel: row[5], Email: row[6], Url: row[7]}
					account.store()
				}
				fmt.Println(row)

			}
		} else if err == io.EOF {
			//	说明已经读到尾了，读完了
			break
		} else {
			panic(err)
		}
	}
}
