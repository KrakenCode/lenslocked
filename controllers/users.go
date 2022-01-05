package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		SignUp Template
	}
}

func (u Users) SignUp(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.SignUp.Execute(w, data)
}

func (u Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Email: ", r.FormValue("email"), " Password: ", r.FormValue("password"))
}
