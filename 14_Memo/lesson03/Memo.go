package main

import "fmt"

//角色状态存储
type RoleStateMemento struct {
	vit int
	atk int
	def int
}

func (r *RoleStateMemento)RoleStateMemento(vit, atk, def int){
	r.atk = atk
	r.vit = vit
	r.def = def
}

func(r *RoleStateMemento) Vitality(value int){
	r.vit = value
}
func(r *RoleStateMemento) GetVitality()int{
	return r.vit
}

func(r *RoleStateMemento) Attack(value int){
	r.atk = value
}
func(r *RoleStateMemento) GetAttack()int{
	return r.atk
}

func(r *RoleStateMemento)Denfense(value int){
	r.def = value
}
func(r *RoleStateMemento)GetDenfense()int{
	return r.def
}

//游戏角色
type GameRole struct {
	vit int
	atk int
	def int
}

func(g *GameRole) saveState() *RoleStateMemento{
	return &RoleStateMemento{
		vit: g.vit,
		atk: g.atk,
		def: g.def,
	}
}

func(g *GameRole) recoveryState(memento *RoleStateMemento){
	g.vit = memento.vit
	g.atk = memento.atk
	g.def = memento.def
}

func(g *GameRole)StateDisplay(){
	fmt.Println("角色当前状态:\n" )
	fmt.Printf("体力: %d\n", g.vit)
	fmt.Printf("攻击力: %d\n", g.atk)
	fmt.Printf("防御力: %d\n", g.def)
}

//获得初始状态
func(g *GameRole)GetInitState(){
	g.vit = 100
	g.atk = 100
	g.def = 100
}

//战斗
func(g *GameRole)Fight(){
	g.vit = 0
	g.atk = 0
	g.def = 0
}


//角色状态管理者
type RoleStateCaretaker struct {
	memento *RoleStateMemento
}

func(r *RoleStateCaretaker)set(memento *RoleStateMemento){
	r.memento = memento
}

func(r *RoleStateCaretaker)get() *RoleStateMemento{
	return r.memento
}



func main(){
	//大战boss之前
	var lixiaoyao = new(GameRole)
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()

	//保存进度
	stateAdmin := new(RoleStateCaretaker)
	stateAdmin.memento = lixiaoyao.saveState()

	//大战boss 损失严重
	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()

	//恢复之前状态
	lixiaoyao.recoveryState(stateAdmin.memento)
	lixiaoyao.StateDisplay()

}
