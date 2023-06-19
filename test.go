package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Question struct {
	qOrder              int
	qType               string
	questionDescription string
	rightAnswer         string
	maxScore            int
	numberOfAttempts    int
	isTimed             bool
	hasMedia            bool
}

func addNewQuestion(qOrder int, qType string, questionDescription string, rightAnswer string, maxScore int, numberOfAttempts int, isTimed bool, hasMedia bool) *Question {
	q := Question{qOrder: qOrder}
	q.questionDescription = questionDescription
	q.qType = qType
	q.maxScore = maxScore
	q.rightAnswer = rightAnswer
	q.numberOfAttempts = numberOfAttempts
	q.isTimed = isTimed
	q.hasMedia = hasMedia
	fmt.Println("Successfully added")
	return &q
}

func main() {

	a := app.New()
	w := a.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
	//var quizName = "My first quiz"
	//var numberOfQuestions = 2
	//timed := false
	var qu1 = addNewQuestion(1, "multipleChoice", "Number of planets in Solar system", "8", 10, 2, false, false)

	fmt.Println("Question number is:", qu1.qOrder)
	//fmt.Println(w32.ABM_GETSTATE)
	//syscall.AF_INET
	//i := unsafe.Sizeof
}
