package main

import "fmt"

//import "github.com/piyush1146115/Learn-Go-Design-Pattern/Creational/Builder"
import "github.com/piyush1146115/Learn-Go-Design-Pattern/Creational/Factory"

func main() {
	ak47, _ := Factory.GetGun("ak47")
	musket, _ := Factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g Factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
