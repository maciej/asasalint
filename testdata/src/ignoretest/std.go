package main

import (
	"fmt"
	"log"
)

func std() {
	var a = []any{1, 2, 3}

	fmt.Println(a)
	fmt.Printf("%v", a)
	fmt.Print(a)

	log.Println(a)
	log.Printf("%v", a)
	log.Print(a)

	log.Fatalln(a)
	log.Fatalf("%v", a)
	log.Fatal(a)

	log.Panicln(a)
	log.Panicf("%v", a)
	log.Panic(a)
}
