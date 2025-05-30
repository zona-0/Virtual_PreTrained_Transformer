package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

func CreateResume() {
	var i int
	var user string
	SuggestionAI()

	fmt.Scan(&user)
	// Clear()
	// TODO: AI resume sec
	if user == "resume" || user == "RESUME" || user == "Resume" {
		fmt.Println(">> Tentu! untuk membuat resume yang efektif, aku butuh informasi dasar terlebih dahulu. Silahkan isi data berikut: ")
		fmt.Println("[A] Informasi Pribadi: ")
		fmt.Printf("	[X] Nama: \n	[X] Alamat: \n	[X] No. HP: \n	[X] Email: \n")
		fmt.Println("[B] Jelaskan tentang dirimu: ")
		fmt.Println("	[X] About Me: ")
		// TODO: TAKE EDUCATION DATA
		fmt.Println("[C] Pendidikan: ")
		if data.EducationCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data pendidikan. Silakan isi melalui menu utama!")
		} else {
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("	[System] Sekolah/Universitas: %s\n	[System] Jenjang: %s\n	[System] Tahun Lulus: %d\n",
					data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
		}

		// TODO: TAKE SKILLS DATA
		fmt.Println("[D] Keterampilan: ")
		if data.SkillCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data keterampilan. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.SkillCount {
				fmt.Printf("	[System] %d. %s\n", i+1, data.Skills[i].Name)
				i += 1
			}
		}

		// TODO: PENGALAMAN KERJA
		fmt.Println("[E] Pengalaman Kerja: ")
		if data.ExperienceCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data pengalaman kerja. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.ExperienceCount {
				fmt.Printf("	[System] %d. %s %s\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
				i += 1
			}
		}

		// TODO: CERTIFICATION SEC
		fmt.Println("[F] Sertifikat: ")

		// TODO: END SEC
		fmt.Println()
		fmt.Printf(">> Silahkan isi informasi di atas atau cukup jawab yang kamu punya sekarang.\n   Setelah itu, aku akan buatkan resume versi teks yang rapi dan siap pakai\n")
		fmt.Println()
		endSec()

		// TODO: USER INPUT SEC
		var nama, alamat, hp, email string

		fmt.Print(">> Nama: ")
		fmt.Scan(&nama)
		fmt.Print(">> Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print(">> No. Hp: ")
		fmt.Scan(&hp)
		fmt.Print(">> Email: ")
		fmt.Scan(&email)
		fmt.Println()

		var kata, aboutme string
		var selesai bool = false

		fmt.Println(">> Tentang Saya: ")
		fmt.Println("   [System] Akhiri dengan tanda '.' untuk mengakhiri")
		for !selesai {
			fmt.Scan(&kata)

			if kata == "." {
				selesai = true
			} else {
				if aboutme == "" {
					aboutme = kata
				} else {
					aboutme = aboutme + " " + kata
				}
			}
		}

		var sertifikat string
		fmt.Print(">> Sertifikat: ")
		fmt.Scan(&sertifikat)

		// TODO: RESUME PRINT SECTION
		fmt.Printf("\n\n")
		fmt.Println(">> Berikut ini adalah resume yang berhasil saya susun berdasarkan data yang Anda berikan.")
		fmt.Println(">> Silakan periksa di bawah ini. Jika ada yang ingin direvisi, Anda bisa kembali kapan saja.")
		fmt.Printf("\n\n")
		fmt.Println("==================================================================")
		fmt.Printf("Nama: %s\nAlamat: %s\n\nTelp: %s | Email: %s\n", nama, alamat, hp, email)

		fmt.Println("------------------------------------------------------------")
		fmt.Println("Tentang Saya: ")
		fmt.Println(aboutme)

		fmt.Println("------------------------------------------------------------")
		fmt.Println("KETERAMPILAN:")
		if data.SkillCount == 0 {
			fmt.Println("  [System] Belum ada keterampilan")
		} else {
			for i = 0; i < data.SkillCount; i++ {
				fmt.Printf("  - %s\n", data.Skills[i].Name)
			}
		}

		fmt.Println("\n------------------------------------------------------------------")

		if data.EducationCount == 0 && data.ExperienceCount == 0 {
			fmt.Println("Pendidikan:")
			fmt.Println("  [System] Belum ada pendidikan")
			fmt.Println("Pengalaman Kerja:")
			fmt.Println("  [System] Belum ada pengalaman kerja")
		} else if data.EducationCount == 0 {
			fmt.Println("PENDIDIKAN:")
			fmt.Println("  [System] Belum ada pendidikan")
			fmt.Println("Pengalaman Kerja:")
			for i = 0; i < data.ExperienceCount; i++ {
				fmt.Printf("  - %s di %s\n", data.Experiences[i].Title, data.Experiences[i].Company)
			}
		} else if data.ExperienceCount == 0 {
			fmt.Println("Pendidikan:")
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("  - %s, %s (%d)\n", data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
			fmt.Println("Pengalaman Kerja:")
			fmt.Println("  [System] Belum ada pengalaman kerja")
		} else {
			fmt.Printf("%-38s| %-38s\n", ">> Pendidikan: ", ">> Pengalaman Kerja: ")

			var maxRows int
			if data.EducationCount > data.ExperienceCount {
				maxRows = data.EducationCount
			} else {
				maxRows = data.ExperienceCount
			}

			for i = 0; i < maxRows; i++ {
				var eduStr string
				var expStr string

				if i < data.EducationCount {
					eduStr = fmt.Sprintf("- %s, %s %d", data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
				}
				if i < data.ExperienceCount {
					expStr = fmt.Sprintf("- %s di %s", data.Experiences[i].Title, data.Experiences[i].Company)
				}

				fmt.Printf("%-38s| %-38s\n", eduStr, expStr)
			}
		}

		fmt.Println("\n------------------------------------------------------------------")
		fmt.Println("Sertifikat:")
		if sertifikat != "" {
			fmt.Printf("  - %s\n", sertifikat)
		} else {
			fmt.Println("  [System] Tidak ada sertifikat")
		}

		fmt.Printf("\n\n")
		fmt.Println(">> Resume telah selesai dibuat. Semoga harimu menyenangkan ", nama)
		fmt.Println()

		//endSec()

		//fmt.Println("TEST: ", aboutme)
		// fmt.Println("[A] Informasi Pribadi: ")
		// fmt.Printf("	[X] Nama: %s\n	[X] Alamat: %s\n	[X] No. HP: %s\n	[X] Email: %s\n", nama, alamat, hp, email)
		//ManageEducation()
	} else {
		fmt.Println()
		fmt.Println("[System] ketik 'resume' untuk bantuan membuat resume atau ketik 'suratkerja' untuk bantuan membuat surat lamaran kerja")
	}
}
