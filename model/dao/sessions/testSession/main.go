package main

import (
	sessions2 "PSHOP/model/dao/sessions"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	sessions2.Open(sessions2.NewRDSOptions("127.0.0.1", 6379, ""))
}

type UserInfo struct {
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Age      uint8  `json:"age,omitempty"`
}

func main() {
	http.HandleFunc("/set", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := sessions2.GetSession(writer, request)

		session.Values["user"] = &UserInfo{
			UserName: "Leon Ding",
			Email:    "ding@ibyte.me",
			Age:      21,
		}
		session.Sync()

		fmt.Fprintln(writer, "set value successful.")
	})

	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := sessions2.GetSession(writer, request)
		jsonstr, _ := json.Marshal(session.Values["user"])
		fmt.Fprintln(writer, string(jsonstr))
	})

	http.HandleFunc("/clean", func(rw http.ResponseWriter, request *http.Request) {
		session, _ := sessions2.GetSession(rw, request)
		// clean session data
		session.Values = make(sessions2.Values)
		//gws.Malloc(&session.Values)
		// sync session modify
		session.Remove()
		fmt.Fprintf(rw, "clean session data successful.")
	})
	http.HandleFunc("/migrate", func(writer http.ResponseWriter, request *http.Request) {
		var (
			session *sessions2.Session
			err     error
		)

		session, _ = sessions2.GetSession(writer, request)
		log.Printf("old session %p \n", session)

		if session, err = sessions2.Migrate(writer, session); err != nil {
			fmt.Fprintln(writer, err.Error())
			return
		}

		log.Printf("old session %v \n", session)
		jsonstr, _ := json.Marshal(session.Values["user"])
		fmt.Fprintln(writer, string(jsonstr))
	})
	log.Println(http.ListenAndServe(":8083", nil))
}
