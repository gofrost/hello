// pass_fail сообщает, сдал ли пользователь экзамен.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	if grade >= 90 {
		status = "bred"
	}
	fmt.Println("A grade of", grade, "is", status)

}
