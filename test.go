package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

func showAddForm() {

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
	var questionsArray []*Question
	qSum := 0
	myApp := app.New()
	myWindow := myApp.NewWindow("Questions")
	//myCanvas := myWindow.Canvas()
	//grey := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	//rect := canvas.NewRectangle(grey)
	//Make add button
	addBtn := widget.NewButton("Add question", func() {
		addNewQuestion(1, "multipleChoice", "Number of planets in Solar system", "8", 10, 2, false, false)
	})
	//Sum all the questions:
	qSumStr := fmt.Sprintln("Questions added:", qSum)
	//content := widget.NewLabel(qSumStr)
	//content := canvas.NewText("1", color.White)
	content := widget.NewLabel(qSumStr)
	//myCanvas.SetContent(rect)
	grid := container.New(layout.NewGridLayout(2), content, addBtn)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.SetContent(grid)
	//myWindow.SetContent(addBtn)
	myWindow.ShowAndRun()

	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")

	// w.ShowAndRun()

	//var quizName = "My first quiz"
	//var numberOfQuestions = 2
	//timed := false
	var newQ = addNewQuestion(1, "multipleChoice", "Number of planets in Solar system", "8", 10, 2, false, false)
	questionsArray = append(questionsArray, newQ)
}
