package modules

import "fmt"

func ManageJob() {
	var input int
	var running bool = true

	Clear()
	for running {
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                           [Job Manager]                            ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Menu]   Setting data lowongan pekerjaan                        ║")
		fmt.Println("║ 2. [Search] Cari data lowongan pekerjaan                           ║")
		fmt.Println("║ 3. [Show]   List data pekerjaan                                    ║")
		fmt.Println("║ 4. [Done]   Selesai dan simpan                                     ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print(">> Pilih menu: ")
		fmt.Scan(&input)

		if input == 1 {
			Clear()
			JobMenu()
		} else if input == 2 {
			Clear()
			listJobListings()
			SearchJob()
		} else if input == 3 {
			Clear()
			listJobListings()
		} else if input == 4 {
			running = false
		} else {
			fmt.Println("[System] Pilihan anda tidak ada di menu. Pilihlah menu yang tersedia dari angka 1 sampai 5")
		}
	}
}

// Git
