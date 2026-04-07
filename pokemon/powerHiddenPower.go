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

func statPrompt(prompt string) float64 {
	var U float64

	fmt.Printf("%s", prompt)
	fmt.Scanln(&U)
	return U
}

func main() {
	HP := statPrompt("Enter the IV value of your HP: ")
	Atk := statPrompt("Enter the IV value of your Attack: ")
	Def := statPrompt("Enter the IV value of your Defense: ")
	Speed := statPrompt("Enter the IV value of your Speed: ")
	SpAtk := statPrompt("Enter the IV value of your Special Attack: ")
	SpDef := statPrompt("Enter the IV value of your Special Defense: ")

	// Next we need to perform our Modulo check on each stat...

	// Test
	fmt.Printf("HP: %.2f, Attack: %.2f, Defense: %.2f, Speed: %.2f, Special Attack: %.2f, Special Defense: %.2f\n", HP, Atk, Def, Speed, SpAtk, SpDef)
}
