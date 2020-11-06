// package controlflow

package main

import (
	"fmt"
	"strings"
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
		a := (strings.Count(user, "a"))
		e := (strings.Count(user, "e"))
		i := (strings.Count(user, "i"))
		o := (strings.Count(user, "o"))
		u := (strings.Count(user, "u"))

		sum := (a + e + (i * 2) + (o * 3) + (u * 4))
		if sum > 10 {
			sum = 10
		}
		distribution[user] = sum
		coins -= sum
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)

}
