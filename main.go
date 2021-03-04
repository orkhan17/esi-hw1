package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bradfitz/slice"
	"github.com/gorilla/mux"
)

// ToDo ...
type ToDo struct {
	ID           string `json:"Id"`
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	Date         string `json:"Date"`
	Time         string `json:"Time"`
	HighPriority bool   `json:"HighPriority"`
	Completed    bool   `json:"Completed"`
}

// ToDos ...
var ToDos []ToDo

func GenHomePage() []byte {
	return []byte("Welcome to the HomePage!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func GenSingleToDo(id string) []byte {
	buf := new(bytes.Buffer)
	var check = false
	for _, todo := range ToDos {
		if todo.ID == id {
			check = true
			json.NewEncoder(buf).Encode(todo)
		}
	}
	if check == false {
		fmt.Println("There is no any Todo with this ID")
	}
	return buf.Bytes()
}

func AddNewToDo(todo ToDo) {
	ToDos = append(ToDos, todo)
}

func DeleteToDo(id string) {
	var check = false
	for index, todo := range ToDos {
		if todo.ID == id {
			check = true
			ToDos = append(ToDos[:index], ToDos[index+1:]...)
		}
	}

	if check == false {
		fmt.Println("There is no any Todo with this ID")
	}
}

func GenAllToDos() []byte {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(ToDos)
	return buf.Bytes()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/todos", returnAllToDo).Methods("GET")
	myRouter.HandleFunc("/todo/{id}", returnSingleToDo).Methods("GET")
	myRouter.HandleFunc("/todo", createNewToDo).Methods("POST")
	myRouter.HandleFunc("/todo/{id}", deleteToDo).Methods("DELETE")
	myRouter.HandleFunc("/todo/{id}", updateToDo).Methods("PUT")
	myRouter.HandleFunc("/todos/{id}", updatePriorityToDo).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnAllToDo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTodos")
	slice.Sort(ToDos[:], func(i, j int) bool {
		return ToDos[i].ID < ToDos[j].ID
	})
	json.NewEncoder(w).Encode(ToDos)
}

func returnSingleToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var check = false
	for _, todo := range ToDos {
		if todo.ID == key {
			check = true
			json.NewEncoder(w).Encode(todo)
		}
	}

	if check == false {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There is no any Todo with this ID")
	}
}

func updateToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	key := 0
	var check = false
	for index, todo := range ToDos {
		if todo.ID == id {
			check = true
			key = index
		}
	}

	var td ToDo = ToDos[key]

	td.Completed = true

	for index, todo := range ToDos {
		if todo.ID == id {
			ToDos = append(ToDos[:index], ToDos[index+1:]...)
		}
	}

	ToDos = append(ToDos, td)
	if check == true {
		json.NewEncoder(w).Encode(td)
	}
	if check == false {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There is no any Todo with this ID")
	}
}

func updatePriorityToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	key := 0
	var check = false
	for index, todo := range ToDos {
		if todo.ID == id {
			check = true
			key = index
		}
	}

	var td ToDo = ToDos[key]

	td.HighPriority = true

	for index, todo := range ToDos {
		if todo.ID == id {
			ToDos = append(ToDos[:index], ToDos[index+1:]...)
		}
	}

	ToDos = append(ToDos, td)
	if check == true {
		json.NewEncoder(w).Encode(td)
	}
	if check == false {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There is no any Todo with this ID")
	}
}

func createNewToDo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo ToDo
	json.Unmarshal(reqBody, &todo)

	slice.Sort(ToDos[:], func(i, j int) bool {
		return ToDos[i].ID < ToDos[j].ID
	})

	var l = len(ToDos) - 1
	var t ToDo = ToDos[l]

	x, err := strconv.Atoi(t.ID)

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	todo.ID = strconv.Itoa(x + 1)

	ToDos = append(ToDos, todo)

	json.NewEncoder(w).Encode(todo)
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var check = false
	for index, todo := range ToDos {
		if todo.ID == id {
			check = true
			ToDos = append(ToDos[:index], ToDos[index+1:]...)
		}
	}

	if check == false {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("There is no any Todo with this ID")
	}

}

func main() {
	ToDos = []ToDo{
		ToDo{ID: "1",
			Title:        "Meeting with friends",
			Description:  "Meeting with my friends in city center",
			Date:         "2021-03-06",
			Time:         "2:00 pm",
			HighPriority: true,
			Completed:    false},
		ToDo{ID: "2",
			Title:        "Shopping",
			Description:  "Don't forget to buy clothes",
			Date:         "2021-03-06",
			Time:         "4:00 pm",
			HighPriority: true,
			Completed:    true},
		ToDo{ID: "3",
			Title:        "Watch football",
			Description:  "Watch Champions League Matches",
			Date:         "2021-03-09",
			Time:         "10:00 pm",
			HighPriority: false,
			Completed:    true},
	}
	handleRequests()
}
