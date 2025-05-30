package modules

import (
	"fmt"
)

const maxJobs = 100

type JobListing struct {
	Title            string
	Salary, Industry int
}

var jobListings [maxJobs]JobListing
var jobCount int

func JobMenu() {
	var choice int
	var running bool = true

	Clear()
	for running {
		// Clear()
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                   [MENU] Lowongan Pekerjaan                  ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Add]    Tambah Lowongan Pekerjaan                        ║")
		fmt.Println("║ 2. [Edit]   Ubah data lowongan pekerjaan                     ║")
		fmt.Println("║ 3. [Delete] Hapus data                                       ║")
		fmt.Println("║ 4. [List]   Tampilkan Data                                   ║")
		fmt.Println("║ 5. [Done]   Selesai dan Kembali                              ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		if choice == 1 {
			addJobListing()
		} else if choice == 2 {
			editJobListing()
		} else if choice == 3 {
			deleteJobListing()
		} else if choice == 4 {
			listJobListings()
		} else if choice == 5 {
			running = false
			// sortJobMenu()
		} else {
			fmt.Println("[System] Pilihan tidak valid")
		}
	}
}

func addJobListing() {
	if jobCount >= maxJobs {
		fmt.Println("[System] Kapasitas lowongan penuh, tidak bisa tambah")
	} else {
		var title string
		var salary, industry int
		fmt.Print(">> Nama Lowongan: ")
		//hJob_NamaLowongan()
		fmt.Scan(&title)
		Header()
		fmt.Println(">> Kode Industri: ")
		fmt.Println("[Suggestion] Gunakan 3 angka (Contoh: 241, 294, 129)")
		fmt.Println("[System] Angka tidak boleh mengandung unsur '0'")
		fmt.Println()
		endSec()
		fmt.Print(">> Input: ")
		fmt.Scan(&industry)
		fmt.Print(">> Gaji: ")
		fmt.Scan(&salary)

		jobListings[jobCount].Title = title
		jobListings[jobCount].Industry = industry
		jobListings[jobCount].Salary = salary
		jobCount++
		fmt.Println("[System] Lowongan berhasil ditambahkan")
		Clear()
	}
}

func deleteJobListing() {
	if jobCount == 0 {
		fmt.Println("[System] Tidak ada lowongan untuk dihapus")
	} else {
		listJobListings()
		var idx, i int
		fmt.Print("Masukkan nomor lowongan yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx < 1 || idx > jobCount {
			fmt.Println("[System] Nomor tidak valid.")
		} else {
			i = idx - 1
			for i < jobCount-1 {
				jobListings[i] = jobListings[i+1]
				i++
			}
			jobListings[jobCount-1].Title = ""
			jobListings[jobCount-1].Industry = 0
			jobListings[jobCount-1].Salary = 0

			jobCount--
			fmt.Println("[System] Lowongan berhasil dihapus")
		}
	}
}

func editJobListing() {
	if jobCount == 0 {
		fmt.Println("[System] Tidak ada lowongan untuk diedit.")
	} else {
		listJobListings()
		var idx int
		fmt.Print(">> Masukkan nomor lowongan yang ingin diedit: ")
		fmt.Scan(&idx)

		if idx < 1 || idx > jobCount {
			fmt.Println("[System] Nomor tidak valid")
		} else {
			idx = idx - 1
			var title, choice string
			var salary, industry int

			fmt.Printf(">> Judul Lowongan (sekarang: %s): ", jobListings[idx].Title)
			fmt.Scan(&title)
			if title != "" {
				jobListings[idx].Title = title
			}

			fmt.Printf(">> Kode Industri (sekarang: %d): ", jobListings[idx].Industry)
			fmt.Scan(&industry)
			if industry != 0 {
				jobListings[idx].Industry = industry
			}

			fmt.Printf(">> Gaji (sekarang: RP %d). Ubah gaji? (Y/N): ", jobListings[idx].Salary)
			fmt.Scan(&choice)
			if choice == "y" || choice == "Y" {
				fmt.Print(">> Masukkan Gaji baru: ")
				fmt.Scan(&salary)
				jobListings[idx].Salary = salary
			} else {
				fmt.Println("[System] Gaji tidak diubah")
			}

			fmt.Println("[System] Lowongan berhasil diperbarui")
		}
	}
}

func listJobListings() {
	Clear()
	var i int

	if jobCount == 0 {
		fmt.Println("[System] Belum ada data lowongan")
	} else {
		fmt.Println("╔═════════╦════════════════════════╦═════════════════╦══════════════╗")
		fmt.Println("║   No    ║        Title           ║    Industri     ║     Gaji     ║")
		fmt.Println("╠═════════╬════════════════════════╬═════════════════╬══════════════╣")
		i = 0
		for i < jobCount {
			fmt.Printf("║  %-6d ║ %-22s ║ %-15d ║ Rp %-9d ║\n",
				i+1,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			i++
		}
		fmt.Println("╚═════════╩════════════════════════╩═════════════════╩══════════════╝")
	}
}

// func sortJobMenu() {
// 	var choose int
// 	if jobCount == 0 {
// 		fmt.Println("[System] Belum ada data lowongan untuk diurutkan")
// 		BackToMenu()
// 	}

// 	fmt.Scan(&choose)
// 	if choose == 55 {
// 		SortBySalary()
// 	} else if choose == 56 {
// 		// SortByReleveance()
// 	}

// 	listJobListings()
// }
