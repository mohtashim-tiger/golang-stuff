package main

import "github.com/google/uuid"

func main() {
	for i := 0; i < 10; i++ {
		uuid := uuid.New()
		println(uuid)
	}
}
