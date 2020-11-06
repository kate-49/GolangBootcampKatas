// package controlflow

package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {

	for _, user := range users {
		sum := 0
		for i := 0; i < len(user); i++ {
			switch string(user[i]) {
			case "a", "A":
				sum += 1
			case "e", "E":
				sum += 1
			case "i":
				sum += 2
			case "o", "O":
				sum += 3
			case "u", "U":
				sum += 4
			}

		}
		if sum > 10 {
			sum = 10
		}
		distribution[user] = sum
		coins -= sum
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)

}
