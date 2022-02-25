package main

import (
	"fmt"
)

var globalvar string = "World" //a global variable

func main() {
	localvar := "Hello" // a local variable
	fmt.Println(localvar + " " + globalvar)
	localvar = "Hi" //re - assigned a variable
	fmt.Println(localvar + " " + globalvar)

}
