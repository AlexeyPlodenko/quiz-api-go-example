# Fasttrack Code Test

## To run API server

CD to ./src/apiserver/ and exec: go run main.go helpers.go handlers.go data.go

## To run CLI client

CD to ./src/cli/ and exec: go install . And after that: ./bin/cli.exe answer --id=1

There are 2 question sets with ID 1 and 2 for testing.

## Quiz

The task is to build a super simple quiz with a few questions and a few alternatives for each question. With one correct answer.

## Preferred Stack

* Backend - Golang
* Database - Just in-memory , so no database

## Preferred Components
* REST API or gRPC
* CLI that talks with the API, preferably using https://github.com/spf13/cobra ( as cli framework )

## User stories

* User should be able to get questions with answers
* User should be able to select just one answer per question.
* User should be able to answer all the questions and then post his/hers answers and get back how many correct answer there was. and that should be displayed to the user.
* User should see how good he/she did compared to others that have taken the quiz , "You where better then 60% of all quizer"