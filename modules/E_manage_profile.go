package modules

import "fmt"

func ManageProfile() {
	var input int
	var running bool = true

	Clear()
	for running {
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                         [Profile Manager]                          ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Manage Education                                                ║")
		fmt.Println("║ 2. Manage Skills                                                   ║")
		fmt.Println("║ 3. Manage Experience                                               ║")
		fmt.Println("║ 4. Show User Data                                                  ║")
		fmt.Println("║ 5. Done                                                            ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&input)

		if input == 1 {
			Clear()
			ManageEducation()
		} else if input == 2 {
			Clear()
			ManageSkill()
		} else if input == 3 {
			Clear()
			ManageExperience()
		} else if input == 4 {
			Clear()
			ShowAllUserData()
		} else if input == 5 {
			running = false
		} else {
			fmt.Println("[System] Pilihan anda tidak ada di menu. Pilihlah menu yang tersedia dari angka 1 sampai 5")
		}
	}
}
