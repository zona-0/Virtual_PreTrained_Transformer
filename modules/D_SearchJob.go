package modules

import "fmt"

func SearchJob() {
	var choose int
	var running bool = true

	for running {
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                   [MENU] Pencarian Lowongan                  ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Search]  Cari berdasarkan posisi pekerjaan               ║")
		fmt.Println("║ 2. [Search]  Cari berdasarkan kata kunci industri            ║")
		fmt.Println("║ 3. [Search]  Cari posisi pekerjaan berdasarkan gaji          ║")
		fmt.Println("║ 4. [Search]  Cari posisi pekerjaan berdasarkan relevansi     ║")
		fmt.Println("║ 5. [Show]    List data pekerjaan                             ║")
		fmt.Println("║ 6. [Done]    Selesai dan Kembali                             ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("Pilih menu: ")
		fmt.Scan(&choose)

		if choose == 1 {
			listJobListings()
			searchJobBinary()
		} else if choose == 2 {
			listJobListings()
			searchByIndustry()
		} else if choose == 3 {
			SortBySalary()
			listJobListings()
		} else if choose == 4 {
			listJobListings()
		} else if choose == 5 {
			listJobListings()
		} else if choose == 6 {
			running = false
		} else {
			fmt.Println("[System] Pilihan anda tidak ada di menu. Pilihlah menu yang tersedia dari angka 1 sampai 5")
		}
	}
}

func searchByIndustry() {
	if jobCount == 0 {
		fmt.Println("[System] Belum ada data lowongan untuk dicari")
	}

	var industryCode int
	fmt.Print("Masukkan Kode Industri yang dicari: ")
	fmt.Scan(&industryCode)

	var found bool = false
	fmt.Println("╔═════════╦════════════════════════╦═════════════════╦══════════════╗")
	fmt.Println("║   No    ║        Title           ║    Industri     ║     Gaji     ║")
	fmt.Println("╠═════════╬════════════════════════╬═════════════════╬══════════════╣")
	for i := 0; i < jobCount; i++ {
		if jobListings[i].Industry == industryCode {
			fmt.Printf("║  %-6d ║ %-22s ║ %-15d ║ Rp %-9d ║\n",
				i+1,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			found = true
		}
	}
	if !found {
		fmt.Println("║                Tidak ditemukan lowongan dengan kode industri ini              ║")
	}
	fmt.Println("╚═════════╩════════════════════════╩═════════════════╩══════════════╝")
}

func searchJobBinary() {
	if jobCount == 0 {
		fmt.Println("[System] Tidak ada lowongan yang tersedia untuk dicari")
		fmt.Println()
		BackToMenu()
	}

	sortJobTitle()

	var keyword string
	fmt.Println("Masukan judul lowongan yang dicari: ")
	fmt.Scan(&keyword)

	var left int
	var right int = jobCount - 1
	var found bool = false
	var mid int = -1

	for left <= right {
		var m int = (right + left) / 2
		if jobListings[m].Title == keyword {
			mid = m
			found = true
			left = right + 1
		} else if jobListings[m].Title < keyword {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	if found {
		fmt.Println("╔═════════╦════════════════════════╦═════════════════╦══════════════╗")
		fmt.Println("║   No    ║        Title           ║    Industri     ║     Gaji     ║")
		fmt.Println("╠═════════╬════════════════════════╬═════════════════╬══════════════╣")

		var i int = mid
		var no int = 1
		for i >= 0 && jobListings[i].Title == keyword {
			fmt.Printf("╠ %-7d ╠ %-22s ╠ %-15d ╠ Rp %-9d ╠\n",
				no,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			i--
			no++
		}

		i = mid + 1
		for i < jobCount && jobListings[i].Title == keyword {
			fmt.Printf("╠ %-7d ╠ %-22s ╠ %-15d ╠ Rp %-9d ╠\n",
				no,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			i++
			no++
		}

		fmt.Println("╚═════════╩════════════════════════╩═════════════════╩══════════════╝")
	} else {
		fmt.Println("[System] Tidak ada lowongan dengan judul tersebut")
	}
}

func sortJobTitle() {
	var i, j int
	var temp JobListing
	for i = 0; i < jobCount-1; i++ {
		for j = i + 1; j < jobCount; j++ {
			if jobListings[i].Title > jobListings[j].Title {
				temp = jobListings[i]
				jobListings[i] = jobListings[j]
				jobListings[j] = temp
			}
		}
	}
}

// func sortJobTitle() {
// 	var i int
// 	for i < jobCount-1 {
// 		var j int = i + 1
// 		for j < jobCount {
// 			if jobListings[i].Title > jobListings[j].Title {
// 				var temp JobListing
// 				temp = jobListings[i]
// 				jobListings[i] = jobListings[j]
// 				jobListings[j] = temp
// 			}
// 			j++
// 		}
// 		i++
// 	}
// }
