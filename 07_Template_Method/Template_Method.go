package main

import "fmt"

/*
	模式特点：通过吧不变的行为搬到父类，去除子类中的重复代码。
	程序实例：考试时使用同一种考卷（父类），不同学生上交自己填写的试卷（子类方法的实现）

	模版模式 ---用绑定实现方法的延迟实现
 */

type IAnswerPaper interface {
	AnswerQuestion1() string
	AnswerQuestion2() string
	AnswerQuestion3() string
}


type TestPaper struct {
	AnsQuestion IAnswerPaper
}

func(this *TestPaper) TestQuestion1(){
	fmt.Println("Question_01")
	fmt.Printf("Answer: %s\n", this.AnswerQuestion1())
}

func(this *TestPaper) TestQuestion2(){
	fmt.Println("Question_02")
	fmt.Printf("Answer: %s\n", this.AnswerQuestion2())
}

func(this *TestPaper) TestQuestion3(){
	fmt.Println("Question_03")
	fmt.Printf("Answer: %s\n", this.AnswerQuestion3())
}

func(this *TestPaper) AnswerQuestion1() string{
	if this.AnsQuestion == nil {
		return ""
	}
	return this.AnsQuestion.AnswerQuestion1()
}

func(this *TestPaper) AnswerQuestion2() string{
	if this.AnsQuestion == nil {
		return ""
	}
	return this.AnsQuestion.AnswerQuestion2()
}

func(this *TestPaper) AnswerQuestion3() string{
	if this.AnsQuestion == nil {
		return ""
	}
	return this.AnsQuestion.AnswerQuestion3()
}


type TestPaperA struct {
	TestPaper
}


func(this *TestPaperA) AnswerQuestion1() string {
	return "X"
}

func(this *TestPaperA) AnswerQuestion2() string {
	return "Y"
}

func(this *TestPaperA) AnswerQuestion3() string {
	return "Z"
}


type TestPaperB struct {
	TestPaper
}


func(this *TestPaperB) AnswerQuestion1() string {
	return "L"
}

func(this *TestPaperB) AnswerQuestion2() string {
	return "M"
}

func(this *TestPaperB) AnswerQuestion3() string {
	return "N"
}

func main() {
	fmt.Println("学生甲抄的试卷: ")
	var studentA = &TestPaper{}
	studentA.AnsQuestion = &TestPaperA{}
	studentA.TestQuestion1()
	studentA.TestQuestion2()
	studentA.TestQuestion3()

	fmt.Println("\n学生乙抄的试卷: ")
	studentB := &TestPaper{}
	studentB.AnsQuestion = &TestPaperB{}
	studentB.TestQuestion1()
	studentB.TestQuestion2()
	studentB.TestQuestion3()
}
