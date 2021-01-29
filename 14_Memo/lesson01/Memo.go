package main

import "fmt"

/*
	备忘录模式:
 */

type GameRole struct {
	vit int
	atk int
	def int
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

func main()  {
	//大战boss前
	var lixiaoyao = new(GameRole)
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()

	//保存进度
	var backup = new(GameRole)
	backup.def = lixiaoyao.def
	backup.atk = lixiaoyao.atk
	backup.vit = lixiaoyao.vit

	//大战boss 损耗严重
	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()

	//恢复之前状态
	lixiaoyao.vit = backup.vit
	lixiaoyao.def = backup.def
	lixiaoyao.atk = backup.atk
	lixiaoyao.StateDisplay()


}
