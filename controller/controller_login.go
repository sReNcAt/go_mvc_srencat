package mvc_controller

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"srencat/mvc_model"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jsonData := responseStatus{"error","400"}

	queryValues := r.URL.Query()

	username := queryValues.Get("username")
	password := queryValues.Get("password")

	if username == "" || password == "" {

	}else{
		username,password := mvc_model.Login(username,password)
		jsonData = responseStatus{"ok","11"}
		fmt.Println(username)
		fmt.Println(password)
		
	}
	fmt.Println(jsonData)

	resJson, err := json.Marshal(jsonData)

	fmt.Println(err)
	fmt.Println(resJson)
	fmt.Println(string(resJson))

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}