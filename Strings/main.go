package main

import (
	"fmt"
	"strings"
)


func main(){
	switchCharacter := false;
	var strBuilder strings.Builder

	for i := 0; i < 100; i++ {
		if switchCharacter {
			strBuilder.WriteString("b")
		} else {
			strBuilder.WriteString("a")
		}
		switchCharacter = !switchCharacter
	}

	fmt.Println(strBuilder.String())
}