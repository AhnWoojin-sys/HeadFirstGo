package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) +1
	fmt.Println("I`ve chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Make a guess : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	for i:=0;i<10;i++{
		fmt.Println("remains ",10 - i)
		if guess < target {
			fmt.Println("Oops. Your guess was LOW. ")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH. ")
		} else {
			fmt.Println("Oh good!")
			break;
		}
	}
}