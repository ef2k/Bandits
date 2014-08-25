package main

import (
	"Bandits/EpsilonGreedy"
	"log"
	"math/rand"
	"time"
)

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	epGreedy := epsilongreedy.New(.1, 3)

	for idx := 0; idx < 10000; idx++ {
		arm := epGreedy.SelectArm()
		reward := random.Intn(10)
		epGreedy.Update(arm, float32(reward))
		log.Print(arm)
	}

	log.Print(epGreedy.Values[0])
	log.Print(epGreedy.Values[1])
	log.Print(epGreedy.Values[2])

	// log.Print(epGreedy)
	// count100 := 0
	// count1000 := 0
	// countElse := 0
	// for idx := 0; idx < 10000; idx++ {
	// 	armChoice := epGreedy.SelectArm()
	// 	if armChoice == 100 {
	// 		count100 += 1
	// 	} else if armChoice == 10000 {
	// 		count1000 += 1
	// 	} else {
	// 		countElse += 1
	// 	}
	// }
	// log.Print("Arm choices: ")
	// log.Print(count100)
	// log.Print(count1000)
	// log.Print(countElse)
}
