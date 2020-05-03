package main

import "fmt"

/*
	模版模式 ---提炼方法
 */

type TestPaper struct {

}

func(this *TestPaper) TestQuestion1(){
	fmt.Println("Question_01")
}

func(this *TestPaper) TestQuestion2(){
	fmt.Println("Question_02")
}

func(this *TestPaper) TestQuestion3(){
	fmt.Println("Question_03")
}




type TestPaperA struct {
	TestPaper
}

func(this *TestPaperA) TestQuestion1(){
	this.TestPaper.TestQuestion1()
	fmt.Println("Answer: A")
}

func(this *TestPaperA) TestQuestion2(){
	this.TestPaper.TestQuestion2()
	fmt.Println("Answer: D")
}

func(this *TestPaperA) TestQuestion3(){
	this.TestPaper.TestQuestion3()
	fmt.Println("Answer: B")
}

type TestPaperB struct {
	TestPaper
}

func(this *TestPaperB) TestQuestion1(){
	this.TestPaper.TestQuestion1()
	fmt.Println("Answer: B")
}

func(this *TestPaperB) TestQuestion2(){
	this.TestPaper.TestQuestion2()
	fmt.Println("Answer: A")
}

func(this *TestPaperB) TestQuestion3(){
	this.TestPaper.TestQuestion3()
	fmt.Println("Answer: D")
}

func main() {
	fmt.Println("学生甲抄的试卷: ")
	studentA := &TestPaperA{}
	studentA.TestQuestion1()
	studentA.TestQuestion2()
	studentA.TestQuestion3()

	fmt.Println("\n学生乙抄的试卷: ")
	studentB := &TestPaperB{}
	studentB.TestQuestion1()
	studentB.TestQuestion2()
	studentB.TestQuestion3()
}
