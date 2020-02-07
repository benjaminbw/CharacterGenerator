package main

import (
	L "benjaminbw/characterSheet/lib"
	"fmt"
	"math/rand"
	"time"
)

/*
func collapseOldCode() {
	fmt.Println()
	//commented out legacy code, making a seperate rollDice file for better

	//this was going to be a more ambitious project. Better save that for after I learn how to use structs...

	//seeding multiple times couln't hurt, right?
	//simulates rolling 1d6
	/*
	   func oneD6() int {
	   	//used to seed here, i think it did hurt. see comments below
	   	dsix := rand.Intn(6)
	   	dsix++
	   	return (dsix)
	   }

	   //does d6 3 times, printing each result

	   func threeD6verbose() int {
	   	fmt.Println()
	   	fmt.Println("Rolling 3d6...")
	   	fmt.Println()

	   	totalscore := 0

	   	//for loop, prints each result for immersion
	   	count := 0
	   	for count < 3 {
	   		dsix := oneD6()
	   		fmt.Println(dsix)
	   		totalscore += dsix
	   		count++
	   	}

	   	fmt.Println()
	   	fmt.Println("You rolled a total of:", totalscore)
	   	return (totalscore)
	   }


	   //does d6 3 times quietly (without printing)
	   func threeD6() int {
	   	totalscore := 0

	   	//for loop, prints each result for immersion
	   	count := 0
	   	for count < 3 {
	   		dsix := oneD6()
	   		totalscore += dsix
	   		count++
	   	}
	   	//there used to be a couple prints here, but when I removed them it started printing out IDENTICAL NUMBERS????
	   	//WTF
	   	//my theory is that the time it took to print gave the seed more time to change?????
	   	//so I replaced the prints with this quicker sleep, and removed the second "seeding" in oneD6

	   	return (totalscore)
	   }


	   //Pass stat to this and output stat bonus. how to do this with idfferent variables????
	   func statBonus() int {
	   	//subtract 10 from the ability score and then divide the total by 2 (round down).
	   	stat =
	   	bonus = (stat-10)/2
	   	return bonus
	   }


}
*/

var strn, dext, cons, intl, wisd, char int

var strbn, dexbn, conbn, intbn, wisbn, chabn int

var hp, hpmax, ac int

func setLowStat() {
	strn = L.Xdx(3, 6)
	dext = L.Xdx(3, 6)
	cons = L.Xdx(3, 6)
	intl = L.Xdx(3, 6)
	wisd = L.Xdx(3, 6)
	char = L.Xdx(3, 6)
}

func setHighStat() {
	strn = L.DropLowest()
	dext = L.DropLowest()
	cons = L.DropLowest()
	intl = L.DropLowest()
	wisd = L.DropLowest()
	char = L.DropLowest()
}

func setStatbonus() {
	strbn = (strn / 2) - 5
	dexbn = (dext / 2) - 5
	conbn = (cons / 2) - 5
	intbn = (intl / 2) - 5
	wisbn = (wisd / 2) - 5
	chabn = (char / 2) - 5
}

func setSubstats() {
	hpmax = L.Dx(6) + conbn
	hp = hpmax
	ac = 10 + dexbn

}

func printSheet() {

	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("     Strength =", strn, "(", strbn, ")")
	fmt.Println("    Dexterity =", dext, "(", dexbn, ")")
	fmt.Println(" Constitution =", cons, "(", conbn, ")")
	fmt.Println(" Intelligence =", intl, "(", intbn, ")")
	fmt.Println("       Wisdom =", wisd, "(", wisbn, ")")
	fmt.Println("     Charisma =", char, "(", chabn, ")")
	fmt.Println()
	fmt.Println(" HP=(", hp, "/", hpmax, ")")
	fmt.Println(" AC =", ac)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//Uses fmt.Scanln as a way to select high or low.
	hilo := 0
	fmt.Println("Press 1 for normal stats (3d6) or 2 for higher stats (4d6 drop lowest)")
	fmt.Scanln(&hilo)

	if hilo == 2 {
		fmt.Println("Higher stats selected.")
		fmt.Println("Rolling dice...")
		setHighStat()
	} else {
		fmt.Println("Normal stats selected.")
		fmt.Println("Rolling dice...")
		setLowStat()
	}

	setStatbonus()
	setSubstats()

	printSheet()

	//this makes the program wait for input before closing
	fmt.Scanln(&hilo)
}
