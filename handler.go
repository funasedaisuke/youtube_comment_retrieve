package main

import (
	"fmt"
	"main/api"
	"main/db"
	"net/http"
)

// Index returns simple response.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func patchDb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func postDb(w http.ResponseWriter, r *http.Request) {
	responseBodyArray := api.ApiClientCall()
	db.DbxInsert(dbx, responseBodyArray)
	fmt.Fprintln(w, "OK")
}

func getData(w http.ResponseWriter, r *http.Request) {
	query := "select * from jarujaruch"
	result := db.DbxSelect(dbx, query)
	// a, _ := json.Marshal(result)
	// fmt.Println(string(a))

	fmt.Fprintln(w, result)
}
