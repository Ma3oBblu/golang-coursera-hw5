package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	error    ApiError
	response interface{}
}

func (srv *MyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//todo: parse r to in params (structure)
	//todo: validate parsed params
	switch r.URL.Path {
	case "/user/profile":
		srv.profile(w, r)
	case "/user/create":
		//srv.create(w, r)
	default:
	}
}

func (srv *MyApi) profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var login string

	switch r.Method {
	case "GET":
		login = r.URL.Query().Get("login")
		if login == "" {
			js, _ := json.Marshal(&Response{ApiError{http.StatusBadRequest, fmt.Errorf("empty login")}, nil})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(js)
			return
		}

	case "POST":
		_ = r.ParseForm()
		login = r.Form.Get("login")
		if login == "" {
			js, _ := json.Marshal(&Response{ApiError{http.StatusBadRequest, fmt.Errorf("empty login")}, nil})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(js)
			return
		}
	}

	profileParams := ProfileParams{login}
	user, err := srv.Profile(ctx, profileParams)

	if err != nil {
		switch err.(type) {
		case ApiError:
			js, _ := json.Marshal(&Response{err.(ApiError), nil})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(err.(ApiError).HTTPStatus)
			w.Write(js)
			return
		default:
			js, _ := json.Marshal(&Response{ApiError{http.StatusInternalServerError, fmt.Errorf("unknown error")}, nil})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(js)
			return
		}
	}

	js, _ := json.Marshal(&Response{nil, user})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (srv *OtherApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//todo: parse r to in params (structure)
	//todo: validate parsed params
	switch r.URL.Path {
	case "/user/profile":
		//srv.profile(w, r)
	case "/user/create":
		//srv.create(w, r)
	default:
	}
}
