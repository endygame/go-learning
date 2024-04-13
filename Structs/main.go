package main

import "fmt"

type human struct {
	age  uint8
	name string
}

type alien struct {
	age  uint8
	name string
}

type lifeform interface {
	getName() string
}

func (e human) getName() string {
	return e.name
}

func (e alien) getName() string {
	return e.name
}

func getGreeting(e lifeform) string {
	return "Hi! " + lifeform.getName(e)
}

func main() {
	var tommy human = human{19, "tommy"}
	fmt.Println(getGreeting(tommy))
}
