package main

import "fmt"

/*
	模版模式
 */

type TestPaperA struct {
}

func(this *TestPaperA) TestQuestion1(){
	fmt.Println("Question_01")
	fmt.Println("Answer: A")
}

func(this *TestPaperA) TestQuestion2(){
	fmt.Println("Question_02")
	fmt.Println("Answer: D")
}

func(this *TestPaperA) TestQuestion3(){
	fmt.Println("Question_03")
	fmt.Println("Answer: B")
}

type TestPaperB struct {
}

func(this *TestPaperB) TestQuestion1(){
	fmt.Println("Question_01")
	fmt.Println("Answer: B")
}

func(this *TestPaperB) TestQuestion2(){
	fmt.Println("Question_02")
	fmt.Println("Answer: A")
}

func(this *TestPaperB) TestQuestion3(){
	fmt.Println("Question_03")
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
