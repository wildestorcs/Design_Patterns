package main

import "fmt"

type Memento struct {
	state string
}

func(m *Memento)SetMemento(state string){
	m.state = state
}

func(m *Memento)GetMemento() string{
	return m.state
}


type Originator struct {
	state string
}

func(o *Originator)CreateMemento() *Memento{
	return &Memento{state: o.state}
}

func (o *Originator)SetMemento(memento *Memento){
	o.state = memento.state
}

func (o *Originator)Show(){
	fmt.Printf("State = %s",o.state)
}


type Caretaker struct {
	memento *Memento
}

func(c *Caretaker)setState(state string){
	c.memento.state = state
}

func(c *Caretaker)getState() string{
	return c.memento.state
}

func main(){
	 o:= new(Originator)
	 o.state = "On\n"
	 o.Show()

     caretaker := new(Caretaker)
     caretaker.memento = o.CreateMemento()

     o.state = "Off\n"
     o.Show()

     o.SetMemento(caretaker.memento)
     o.Show()




}
