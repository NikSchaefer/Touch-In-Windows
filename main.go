package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not Enough Arguments Provided")
	}
	dir, err := os.Getwd()
	check(err)
	fmt.Println(dir)
	
	file, err := os.Create(os.Args[1])
	check(err)
	defer file.Close()
	fmt.Println(file)


}

func check(e error) {
    if e != nil {
        panic(e)
    }
}