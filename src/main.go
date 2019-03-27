package main

// @TODO caching of questions
// @TODO fix "better" logic
// @TODO split the code

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Question 's structure
type Question struct {
	Question        string    `json:"question"`
	Answers         [3]string `json:"answers"`
	correctAnswerIx int       `json:"-"`
}

// A list of questions to send to the client
var questionsDataSet = map[int][3]Question{
	1: [3]Question{
		Question{
			Question:        "What is the name of the network of computers from which the Internet has emerged?",
			Answers:         [3]string{"Internet", "Arpanet", ".Net"},
			correctAnswerIx: 1,
		},
		Question{
			Question:        "Which unit is an indication for the sound quality of MP3?",
			Answers:         [3]string{"Kbps", "Kg", "Km"},
			correctAnswerIx: 0,
		},
		Question{
			Question:        "In what year was Google launched on the web?",
			Answers:         [3]string{"1898", "2098", "1998"},
			correctAnswerIx: 2,
		},
	},
	2: [3]Question{
		Question{
			Question:        "When did the French Revolution end?",
			Answers:         [3]string{"1799", "1999", "Never"},
			correctAnswerIx: 0,
		},
		Question{
			Question:        "What was Louis Armstrong's chosen form of music?",
			Answers:         [3]string{"Jazz", "Rock", "Pop"},
			correctAnswerIx: 0,
		},
		Question{
			Question:        "What is the Italian word for pie?",
			Answers:         [3]string{"Pizza", "Pizzzzzza", "Ravioli"},
			correctAnswerIx: 0,
		},
	},
}

/*
	How many users correctly answered this amount of questions. Index in array
	represents amount of correctly answered questions, and value - users.
*/
var usersAnsweredCorrectly [4]int

// Total amount of users taken the quiz
var totalAnsweredUsers = 0

// Main
func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/questions/{id:\\d+}", QuestionsHandler).Methods("GET")
	subRouter.HandleFunc("/questions/{id:\\d+}/answers", AnswersHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}

// QuestionsHandler encodes  questions into JSON format and puts to the response.
func QuestionsHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("NEW REQUEST", "QuestionsHandler")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		returnStatusBadRequest(response, "Invalid ID.")
		return
	}

	questions, questionsExist := questionsDataSet[id]
	if !questionsExist {
		returnStatusNotFound(response)
		return
	}

	json.NewEncoder(response).Encode(questions)
}

// AnswersHandler saves client's answers.
func AnswersHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("NEW REQUEST", "AnswersHandler")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		returnStatusBadRequest(response, "")
		return
	}

	questions, questionsExist := questionsDataSet[id]
	if !questionsExist {
		returnStatusNotFound(response)
		return
	}

	// parsing JSON
	var answers []struct {
		AnswerIx int `json:"answerIx"`
	}
	err = json.NewDecoder(request.Body).Decode(&answers)
	if err != nil {
		returnStatusBadRequest(response, "")
		return
	}

	totalQuestions := len(questions)

	// checking all questions are answered and there are no "extra" answers
	if totalQuestions != len(answers) {
		returnStatusBadRequest(response, "Invalid amount of answers.")
		return
	}

	// calculate total correct answers
	correctAnswers := 0
	for questionIx, answer := range answers {
		if answer.AnswerIx == questions[questionIx].correctAnswerIx {
			correctAnswers++
		}
	}

	// update general statistics
	usersAnsweredCorrectly[correctAnswers]++
	totalAnsweredUsers++

	// respond with user's results
	userIsBetterThanOthers := getUserIsBetterThanOthersPerc(correctAnswers)
	result := struct {
		CorrectAnswers    int    `json:"correctAnswers"`
		ComparingToOthers string `json:"comparingToOthers"`
	}{
		CorrectAnswers:    correctAnswers,
		ComparingToOthers: "You were better than " + strconv.Itoa(userIsBetterThanOthers) + "% of all quizers",
	}
	json.NewEncoder(response).Encode(result)
}

// HTTP 400 status response with custom message
func returnStatusBadRequest(response http.ResponseWriter, body string) {
	response.WriteHeader(http.StatusBadRequest)
	response.Header().Set("Content-Type", "plain/text")
	if body != "" {
		io.WriteString(response, body+"\n")
	}
}

// HTTP 404 status response
func returnStatusNotFound(response http.ResponseWriter) {
	response.WriteHeader(http.StatusNotFound)
}

// Returns by how many percentage current user has answered better than the rest
func getUserIsBetterThanOthersPerc(correctAnswers int) int {
	usersAmountAnsweredWorse := 0
	for i := 0; i < correctAnswers; i++ {
		usersAmountAnsweredWorse += usersAnsweredCorrectly[i]
	}
	var res int
	if totalAnsweredUsers > 0 {
		res = int(math.Ceil(float64(usersAmountAnsweredWorse) / float64(totalAnsweredUsers) * 100.0))
	} else {
		res = 100
	}
	return res
}
