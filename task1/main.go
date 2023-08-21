package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct{
	Name string
	Sport string
}

type Action struct{
	Sportsman Human
}

func (h *Human) bestAction() string{
	return h.Sport
}

func (a *Action) mainActivity() {
	fmt.Printf("%s's main activity is %s",a.Sportsman.Name,a.Sportsman.bestAction())
}

func main(){
	john:= Human{Name: "John", Sport: "Footbal",}
	Action:= Action{Sportsman: john}
	Action.mainActivity()
}