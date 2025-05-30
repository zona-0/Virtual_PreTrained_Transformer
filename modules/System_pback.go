package modules

import "fmt"

func BackToMenu() {
	var backMenu string
	fmt.Println()
	fmt.Println("[System] Type 'pback' to return to the main menu")
	for {
		fmt.Scanln(&backMenu)
		if backMenu == "pback" {
			break
		} else {
			fmt.Println()
		}
	}
	Clear()
}
