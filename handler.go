package main

import (
	"fmt"
	"main/db"
	"net/http"
)

// Index returns simple response.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func postDb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func getData(w http.ResponseWriter, r *http.Request) {
	query := "select * from jarujaru"
	result := db.DbxSelect(dbx, query)
	fmt.Fprintln(w, result)
}
