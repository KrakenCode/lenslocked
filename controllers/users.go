package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		SignUp Template
	}
}

func (u Users) SignUp(w http.ResponseWriter, r *http.Request) {
	u.Templates.SignUp.Execute(w, nil)
}
