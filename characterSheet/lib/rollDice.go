package lib

import (
	"math/rand"
	"sort"
)

//! IMPORTANT! Use seed:
//! rand.Seed(time.Now().UnixNano())
//! once before this used. It does not seed itself.

//Dx returns a value between 1 and x - used for simulation of dice.
func Dx(x int) int {
	d := rand.Intn(x)
	d++
	return (d)
}

//Xdx simulates d amount of x-sided dice
func Xdx(d int, x int) int {
	totalscore := 0
	count := 0
	for count < d {
		dx := Dx(x)
		totalscore += dx
		count++
	}
	return (totalscore)
}

//DropLowest takes 4 random numbers between 1&6, removes the lowest and adds the rest.
func DropLowest() int {
	rolls := []int{Dx(6), Dx(6), Dx(6), Dx(6)}
	sort.Ints(rolls)
	droplowest := rolls[1] + rolls[2] + rolls[3]
	return droplowest
}
//rolls is a slice filled with 4*1d6. sort.Ints(rolls) orders them from lowest to highest. then you simply add the higher values.
//This isn't flexible at all (like Xdx), but it works.



/* 
func collapseOldCode(){

   *Lots of old code here. I just copy-pasted the same code with slight variations over and over again,
   *Instead of writing more flexible functions.
   ?Maybe re-use the old function names (like D6 etc.) for use as shortcuts? 
   *Nah, the current ones are readable enough.

	//D4 returns a value between 1 and 4
	func D4() int {
	dfour := rand.Intn(4)
	dfour++
	return (dfour)


	//D6 returns a value between 1 and 6
	func D6() int {

	dsix := rand.Intn(6)
	dsix++
	return (dsix)
	}

	//D8 returns a value between 1 and 8
		func D8() int {

		deight := rand.Intn(8)
		deight++
		return (deight)
	}

	//D10 returns a value between 1 and 10
		func D10() int {

		dten := rand.Intn(10)
		dten++
		return (dten)
	}

	//D12 returns a value between 1 and 12
		func D12() int {

		dtwelve := rand.Intn(12)
		dtwelve++
		return (dtwelve)
	}

	//D100 returns a value between 1 and 100
		func D100() int {

		dhundred := rand.Intn(100)
		dhundred++
		return (dhundred)
	}

	//Xd6 returns xd6 - input amount of dice to add
	func Xd6(d int) int {
		totalscore := 0

		count := 0
		for count < d {
			dsix := D6()
			totalscore += dsix
			count++
		}

		return (totalscore)
	}

	Now obsolete, replacing this with a function that utilizes Xd6
	func threeD6() int {
	totalscore := 0

	count := 0
	for count < 3 {
		dsix := d6()
		totalscore += dsix
		count++
	}

	return (totalscore)

	//ThreeD6 - legacy fix using Xd6
	func ThreeD6() int {
		threeD6 := Xd6(3)
		return threeD6
	}
}
*/


