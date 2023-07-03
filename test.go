package main

import (
	"fmt"
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

func checkForDoubles(q1, q2 Question) bool {
	var isDouble bool
	if q1 == q2 {
		isDouble = true
	} else {
		isDouble = false
	}
	return isDouble
}

func countAddedQuestions(questList []*Question) int {
	numQ := len(questList)
	return numQ
}

func main() {
	var questionsArray []*Question
	//qSum := 0
	//myApp := app.New()
	//myWindow := myApp.NewWindow("Questions")
	//myCanvas := myWindow.Canvas()
	//grey := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	//rect := canvas.NewRectangle(grey)
	//Make add button
	/*addBtn := widget.NewButton("Add question", func() {
		//var voidance struct
		//qSum, voidance  = addNewQuestion(qSum, 1, "multipleChoice", "Number of planets in Solar system", "8", 10, 2, false, false)
		headText := "Type your question here"
		hTxt := canvas.NewText(headText, color.Gray)
		qForm := widget.NewEntry()
		newQuestion := container.New(layout.NewBoxLayout(), hTxt, layout.NewSpacer(), qForm)
	})*/
	//Sum all the questions:
	//qSumStr := fmt.Sprintln("Questions added:", qSum)
	//content := widget.NewLabel(qSumStr)
	//content := canvas.NewText("1", color.White)
	////content := widget.NewLabel(qSumStr)
	//myCanvas.SetContent(rect)
	////grid := container.New(layout.NewGridLayout(2), content, addBtn)
	////myWindow.Resize(fyne.NewSize(800, 600))
	////myWindow.SetContent(grid)
	//myWindow.SetContent(addBtn)
	////myWindow.ShowAndRun()

	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")

	// w.ShowAndRun()

	//var quizName = "My first quiz"
	//var numberOfQuestions = 2
	//timed := false
	var newQ *Question
	newQ = addNewQuestion(1, "multipleChoice", "Number of planets in Solar system", "8", 10, 2, false, false)
	questionsArray = append(questionsArray, newQ)
	fmt.Println("Total number of questions:")
	questionsArray = append(questionsArray, newQ)
	//time.Sleep(2 * time.Second)
	numQ := countAddedQuestions(questionsArray)
	fmt.Println(numQ)
	isDoubleQuestion := checkForDoubles(*questionsArray[numQ-1], *questionsArray[numQ-2])
	if isDoubleQuestion {
		fmt.Println("Last 2 questions are doubles")
	} else {
		fmt.Println("Unique questions, proceeding")
	}
}
