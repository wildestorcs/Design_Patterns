package main

import "fmt"

type Player struct {
	Name string
	Play Play
}

func(p *Player)PlayerName(name string){
	p.Name = name
}

type Play interface {
	Attack()
	Defense()
}

type Forwords struct {
	Player
}

func(f *Forwords)PlayerName(name string){
	f.Name = name
}

func(f *Forwords)Attack() {
	fmt.Printf("前锋: %s 发起进攻\n", f.Name)
}

func(f *Forwords)Defense(){
	fmt.Printf("前锋: %s 发起防守\n", f.Name)
}


type Center struct {
	Player
}

func(f *Center)PlayerName(name string){
	f.Name = name
}

func(f *Center)Attack() {
	fmt.Printf("中锋: %s 发起进攻\n", f.Name)
}

func(f *Center)Defense(){
	fmt.Printf("中锋: %s 发起防守\n", f.Name)
}



type Guards struct {
	Player
}

func(f *Guards)PlayerName(name string){
	f.Name = name
}

func(f *Guards)Attack() {
	fmt.Printf("后卫: %s 发起进攻\n", f.Name)
}

func(f *Guards)Defense(){
	fmt.Printf("后卫: %s 发起防守\n", f.Name)
}


type ForeignCenter struct {
	Name string
}

func(f *ForeignCenter)进攻(){
	fmt.Printf("外籍中锋: %s 进攻\n", f.Name)
}

func (f *ForeignCenter)防御(){
	fmt.Printf("外籍中锋: %s 防御\n", f.Name)
}

type Translator struct {
	ForeignCenter
}

func (t *Translator)PlayerName(name string){
	t.Name = name
}

func (t *Translator)Attack() {
	t.进攻()
}

func (t *Translator)Defense() {
	t.防御()
}


func main(){
	playerb := new(Forwords)
	playerb.PlayerName("巴蒂尔")
	playerb.Attack()

	playerm := new(Guards)
	playerm.PlayerName("麦克格雷迪")
	playerm.Defense()

	playerym := new(Translator)
	playerym.PlayerName("姚明")
	playerym.Attack()
	playerym.Defense()


}
