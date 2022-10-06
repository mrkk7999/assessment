package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	PersonName string `json:"personName,omitempty"`
}

var names = []Name{
	{PersonName: "Kiran"},
	{PersonName: "Divya"},
}

func checkIfNamePresent(name string) (bool, error) {
	for i := 0; i < len(names); i++ {
		if names[i].PersonName == name {
			return true, nil
		}
	}
	return false, errors.New("person name not found")
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// custom error
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var name Name
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// custom err
		// here is need to add one more field to error as we cannot define it in status so adding
		// body section to error message where we will written in depth about error body and
		// we can skip the status code for that
		//return fmt.Errorf("error in parsing request body : %v", err.Eror())
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err = json.Unmarshal(body, &name); err != nil {
		// custom err
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(name)
	fmt.Println(name.PersonName)
	isFound, err := checkIfNamePresent(name.PersonName)
	if err != nil {
		// custom error
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if !isFound {
		// custom error
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome home kiran"))
	return
}
func main() {
	// registering our handle
	http.HandleFunc("/home", home)

	// creating server on port number 9090 and starts listening
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("Port is already in used")
		return
	}
}
