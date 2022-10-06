package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type myError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}

type HttpError struct {
	Eror    error
	Status  int
	Message string
}

func (h HttpError) Error() string {
	if h.Status == 0 {
		return fmt.Sprintf(h.Message)
	}
	return fmt.Sprintf("status " + " : " + strconv.Itoa(h.Status) + " " + "\n" + h.Message)
}

func (h *HttpError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("not able to parse the body: %v", err)
	}
	return body, nil
}

func (h *HttpError) ResponseHeaders() (int, map[string]string) {
	return h.Status, map[string]string{
		"Content-Type": "application/json;",
	}
}

type errHandlerWrapper func(http.ResponseWriter, *http.Request) error

func (fn errHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// calls handler function(here home)
	err := fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	myError, ok := err.(HttpError)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := myError.ResponseBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	status, headers := myError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

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

func home(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		// custom error
		return &HttpError{
			Eror:    nil,
			Status:  http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		}
	}
	var name Name
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// custom err
		// here is need to add one more field to error as we cannot define it in status so adding
		// body section to error message where we will written in depth about error body and
		// we can skip the status code for that
		//return fmt.Errorf("error in parsing request body : %v", err.Eror())
		return &HttpError{
			Eror:    err,
			Message: "Failed to parse body",
		}
	}
	if err = json.Unmarshal(body, &name); err != nil {
		// custom err
		return &HttpError{
			Eror:    err,
			Message: "Failed to unmarshal json data into go object",
		}
	}
	fmt.Println(name)
	fmt.Println(name.PersonName)
	isFound, err := checkIfNamePresent(name.PersonName)
	if err != nil {
		// custom error
		return &HttpError{
			Eror:    err,
			Status:  401,
			Message: "Not authorized person to enter home",
		}
	}
	if !isFound {
		// custom error
		return &HttpError{
			Eror:    err,
			Status:  401,
			Message: "Not authorized person to enter home",
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome home kiran"))
	return nil
}
func main() {
	// handle func only accepts handler with signature func(ResponseWriter, *Request)
	//http.HandleFunc("/home", home)

	// registering our handle
	http.Handle("/home", errHandlerWrapper(home))

	// creating server on port number 9090 and starts listening
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("Port is already in used")
		return
	}
}
