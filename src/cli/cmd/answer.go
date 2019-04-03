package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Question 's structure
type Question struct {
	Question string    `json:"question"`
	Answers  [3]string `json:"answers"`
}

// Answer 's structure
type Answer struct {
	AnswerID int `json:"answerId"`
}

// Answers 's structure
type Answers [3]Answer

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer questions",
	Long:  "Answer questions...",
	Run: func(cmd *cobra.Command, args []string) {
		// @TODO split into functions

		// read ID from the argument
		questionID, _ := cmd.Flags().GetString("id")

		// fetch the question from API
		questionReader, err := FetchQuestion(questionID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer questionReader.Close()

		// parsing JSON
		var questions [3]Question
		err = json.NewDecoder(questionReader).Decode(&questions)
		if err != nil {
			fmt.Println("Failed to parse question's JSON")
			os.Exit(1)
		}

		// go through the questions and output to the user with options

		var answer string
		var answers [3]Answer

		// lets try strings.Builder for string concat.
		var questionStr strings.Builder
		for i, question := range questions {
			questionStr.WriteString(question.Question + ":\n")
			answersLen := len(question.Answers)
			for j, answer := range question.Answers {
				questionStr.WriteString("[")
				questionStr.WriteString(strconv.Itoa(j + 1))
				questionStr.WriteString("] ")
				questionStr.WriteString(answer)
				questionStr.WriteString("\n")
			}
			questionStr.WriteString("Enter the right answer's number: ")
			fmt.Println(questionStr.String())
			questionStr.Reset()

			// read user input
			fmt.Scanln(&answer)

			// validate user input
			answerNum, err := strconv.Atoi(answer)
			if err != nil {
				fmt.Println("Answer must be a number")
				os.Exit(1)
			}
			if answerNum <= 0 || answerNum > answersLen {
				answersLenStr := strconv.Itoa(answersLen)
				fmt.Println("Answer must be in the range 1-" + answersLenStr)
				os.Exit(1)
			}

			answers[i] = Answer{
				AnswerID: answerNum,
			}
		}

		// sending our answers and outputting results to the CLI
		answersResultReader, err := sendAnswer(questionID, answers)
		if err != nil {
			fmt.Println("Failed to send answers.", err)
			os.Exit(1)
		}
		defer answersResultReader.Close()

		respBody, err := ioutil.ReadAll(answersResultReader)
		if err != nil {
			fmt.Println("Error reading response.", err)
			os.Exit(1)
		}

		fmt.Println("Response:\n", string(respBody))
	},
}

// Send our answers to the API for processing
func sendAnswer(questionID string, answers Answers) (io.ReadCloser, error) {
	answersJSON, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("Failed to JSON encode our answers.", err)
		os.Exit(1)
	}
	// json.NewEncoder(response).Encode(&answers)

	answerURL := apiStartURL + "/questions/" + questionID + "/answers"
	fmt.Println(answerURL)
	resp, err := http.Post(answerURL, "application/json", bytes.NewBuffer(answersJSON))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP status is not OK")
	}

	return resp.Body, nil
}

// init :)
func init() {
	rootCmd.AddCommand(answerCmd)

	answerCmd.Flags().StringP("id", "i", "", "Question ID")
}
