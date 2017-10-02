package main

import (
    "log"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
    "time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/trivium", GetTrivium).Methods("GET")
    router.HandleFunc("/trivium", CreateTrivium).Methods("POST")
    router.HandleFunc("/trivium", DeleteTrivium).Methods("DELETE")

    trivia = append(trivia, Trivium{Prompt: "What is Dr. Seuss's real name?", Answer: "Theodore Geisel"})
    trivia = append(trivia, Trivium{Prompt: "In what country is the region of Andalusia located?", Answer: "Spain"})
    trivia = append(trivia, Trivium{Prompt: "In what theater was Lincoln killed?", Answer: "Ford"})
    var t Trivium
    t.Prompt = "Test Prompt!!"
    trivia = append(trivia, t)



    log.Fatal(http.ListenAndServe(":8000", router))

}

func GetTrivium(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(len(trivia))
	json.NewEncoder(w).Encode(trivia[rand])
	return
}

func CreateTrivium(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
	decoder := json.NewDecoder(r.Body)

	var t Trivium
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	
    trivia = append(trivia, t)
    json.NewEncoder(w).Encode(trivia)
    defer r.Body.Close()

    return
}

func DeleteTrivium(w http.ResponseWriter, r *http.Request) {
    // TODO(CHLOE): Implement.
    return
}

type Trivium struct {
    Prompt        string   `json:"prompt"`
    Answer string   `json:"answer"`
    AnswerDetails  string   `json:"answer_details"`
    Attribution   string `json:"attribution"`
}

var trivia []Trivium
