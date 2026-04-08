package main

import "fmt"

/*
Step 1: Declare variables
		HP, Attack, Defense, Speed, Special Attack, Special Defense
		U	  V		   W	   X		  Y				 Z
Step 2: Check (if remainder of Stat / 4 equals 2 or 3):

		For each stat (Hp, Atk, Def, Speed, Sp.Atk, Sp.Def), check if the IV value divided by 4 has a remainder of 2 or 3 (modulo). If yes, the variable is used; otherwise, it is 0.
		HP - 1		Speed - 8
		ATK - 2		Sp.Atk - 16
		DEF - 4		Sp.Def - 32
		So, when the stat is divided by 4, if remainder == 2 || 3, use value above; if remainder != 2 || 3, use 0.  (Modulo operand)

	Formula:
		[(U + V + W + X + Y + Z) * 40/63 + 30]
*/

func statPrompt(prompt string) int {
	var stat int

	fmt.Printf("%s", prompt)
	fmt.Scanln(&stat)
	return stat
}

func remainderCheck(stat int, value int) int {
	if stat%4 == 2 || stat%4 == 3 {
		return value
	}
	return 0
}

func powerFormula(u, v, w, x, y, z int) float64 {
	sum := u + v + w + x + y + z
	return float64(sum)*40.0/63.0 + 30.0
}

func main() {
	HP := statPrompt("Enter the IV value of your HP: ")
	Atk := statPrompt("Enter the IV value of your Attack: ")
	Def := statPrompt("Enter the IV value of your Defense: ")
	Speed := statPrompt("Enter the IV value of your Speed: ")
	SpAtk := statPrompt("Enter the IV value of your Special Attack: ")
	SpDef := statPrompt("Enter the IV value of your Special Defense: ")
	// var confirm string

	fmt.Printf("You entered the following values for your IV's. HP: %d, Attack: %d, Defense: %d, Speed: %d, Special Attack: %d, Special Defense: %d\n", HP, Atk, Def, Speed, SpAtk, SpDef)
	// fmt.Printf("Do these look correct? (y/n)")
	// fmt.Scanln(&confirm)

	u := remainderCheck(HP, 1)
	v := remainderCheck(Atk, 2)
	w := remainderCheck(Def, 4)
	x := remainderCheck(Speed, 8)
	y := remainderCheck(SpAtk, 16)
	z := remainderCheck(SpDef, 32)

	fmt.Printf("The base power of your Hidden Power ability is: %.2f\n", powerFormula(u, v, w, x, y, z))
}
