package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("hello")

	portNumber := os.Getenv("PORT")
	fmt.Println(portNumber)
}
