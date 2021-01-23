package main

import "fmt"

type State interface {
	WriteProgram(work *Work)
}

type ForenoonState struct {
}

func (f *ForenoonState)WriteProgram(work *Work){
	if work.Hour < 12{
		fmt.Printf("当前时间: %d 点,上午工作,精神百倍\n", work.Hour)
	}else{
		work.setState(new(NoonState))
		work.WriteProgram()
	}
}

type NoonState struct {
}

func (n *NoonState)WriteProgram(work *Work){
	if work.Hour < 13{
		fmt.Printf("当前时间: %d 点, 饿了 午饭;犯困 午休\n", work.Hour)
	}else{
		work.setState(new(AfternoonState))
		work.WriteProgram()
	}
}

type AfternoonState struct {
}

func(a *AfternoonState)WriteProgram(work *Work){
	if work.Hour < 17{
		fmt.Printf("当前时间: %d 点,下午状态不错,继续努力\n", work.Hour)
	}else{
		work.setState(new(EvenningState))
		work.WriteProgram()
	}
}

type EvenningState struct {
}

func(e *EvenningState)WriteProgram(work *Work){
	if work.Finish == true{
		work.setState(new(ResetState))
		work.WriteProgram()
	}else{
		if work.Hour < 21{
			fmt.Printf("当前时间: %d 点,加班,疲惫至极\n", work.Hour)
		}else{
			work.setState(new(SleepingState))
			work.WriteProgram()
		}
	}
}

type SleepingState struct {
}

func (s *SleepingState)WriteProgram(work *Work){
	fmt.Printf("当前时间: %d 点,睡着了...\n", work.Hour)
}


type ResetState struct {
}

func (r *ResetState) WriteProgram(work *Work){
	fmt.Printf("当前时间: %d点,下班了", work.Hour)
}


type Work struct {
	Current State
	Hour int
	Finish bool
}

func(w *Work) Work(){
	w.Current = new(ForenoonState)
}

func(w *Work)getHour() int{
	return w.Hour
}

func(w *Work)setHour(value int){
	w.Hour = value
}

func(w *Work)getTaskFinished() bool{
	return w.Finish
}

func(w *Work)setTaskFinished(value bool){
	w.Finish = value
}

func(w *Work)setState(state State){
	w.Current = state
}

func (w *Work)WriteProgram(){
	w.Current.WriteProgram(w)
}


func main(){
	//紧急项目
	emergencyProjects := new(Work)
	emergencyProjects.Hour = 9
	context := new(ForenoonState)
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 12
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 13
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 14
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 17
	context.WriteProgram(emergencyProjects)

	//emergencyProjects.Finish = true
	emergencyProjects.Finish = false
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 19
	context.WriteProgram(emergencyProjects)
	emergencyProjects.Hour = 22
	context.WriteProgram(emergencyProjects)

}
