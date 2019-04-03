package main

// Question 's structure
type Question struct {
	Question        string    `json:"question"`
	Answers         [3]string `json:"answers"`
	CorrectAnswerID int       `json:"-"`
}

// QuestionsDataSet is a list of questions to send to the client. DB emulation
var QuestionsDataSet = map[int][3]Question{
	1: [3]Question{
		Question{
			Question:        "What is the name of the network of computers from which the Internet has emerged?",
			Answers:         [3]string{"Internet", "Arpanet", ".Net"},
			CorrectAnswerID: 1,
		},
		Question{
			Question:        "Which unit is an indication for the sound quality of MP3?",
			Answers:         [3]string{"Kbps", "Kg", "Km"},
			CorrectAnswerID: 0,
		},
		Question{
			Question:        "In what year was Google launched on the web?",
			Answers:         [3]string{"1898", "2098", "1998"},
			CorrectAnswerID: 2,
		},
	},
	2: [3]Question{
		Question{
			Question:        "When did the French Revolution end?",
			Answers:         [3]string{"1799", "1999", "Never"},
			CorrectAnswerID: 0,
		},
		Question{
			Question:        "What was Louis Armstrong's chosen form of music?",
			Answers:         [3]string{"Jazz", "Rock", "Pop"},
			CorrectAnswerID: 0,
		},
		Question{
			Question:        "What is the Italian word for pie?",
			Answers:         [3]string{"Pizza", "Pizzzzzza", "Ravioli"},
			CorrectAnswerID: 0,
		},
	},
}

/* UsersAnsweredCorrectly contains how many users correctly answered this
amount of questions. Index in array represents amount of correctly answered
questions, and value - users. */
var UsersAnsweredCorrectly [4]int

// TotalAnsweredUsers has total amount of users taken the quiz.
var TotalAnsweredUsers = 0
