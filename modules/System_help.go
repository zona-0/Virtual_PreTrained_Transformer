package modules

import "fmt"

func Help() {
	Clear()
	fmt.Println("[System] W E L C O M E")
	fmt.Printf("\n\n")
	fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     C A R E E R  E D G E                               ║")
	fmt.Println("║       AI Assistant for Smart Resumes & Cover Letter Creation           ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
	fmt.Println(" [System] This is help menu for user!")

	Header()
	// TODO: Commands section
	fmt.Println(" [Skills]   Pada saat menambahkan skill, ketik 'done' untuk berhenti atau selesai")
	fmt.Println(" [Commands] In main menu you can see multiple number and option, you can select the option using the number")
	fmt.Println("            for example: type '1' for open Manage Skill or type '2' to open Manage Education")
	fmt.Println(" [Commands] Type 'l-ctrl + c' to stop the program")
	fmt.Println(" [Commands] Type 'cls' to clear terminal")
	fmt.Println(" [Commands] Type 'pback' for back to the main menu")

	endSec()

	fmt.Println()
}
