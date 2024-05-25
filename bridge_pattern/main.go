package main

import "fmt"

func main() {

	hpPrinter := &Hp{}
	esponPrinter := &Espon{}

	macComputer := &Mac{}
	windowsComputer := &Windows{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(esponPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(esponPrinter)
	windowsComputer.Print()
	fmt.Println()

}
