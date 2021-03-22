package main

import (
	"html/template"
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = AccountGET(w, r)
		break
	case "POST":
		break
	case "PUT":
		break
	case "DELETE":
		break
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AccountGET(w http.ResponseWriter, r *http.Request) (err error) {
	var accounts []*Account
	for _, account := range AccountById {
		accounts = append(accounts, account)
	}
	t, _ := template.ParseFiles("tmp.html")
	t.Execute(w, accounts)
	return nil
}
