package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

func RecommendJob() {
	var i, j,
		backend, frontend, analyst, writer, teacher, marketing, designer, devops, mobile, security, management, finance, customer_service,
		game, video, law, health, admin, logistics, hr, psychology, chef, events, cashier, operator, fnb, creator, total,
		fullstack, it, companyScore,
		educationScore int
	var name, title, school, degree, major, company string

	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║               AI CAREER SUGGESTION                 ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║ [SYSTEM] Analyzing profile: Skills, Experience,    ║")
	fmt.Println("║         and Education to generate recommendations  ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	for i = 0; i < data.SkillCount; i++ {
		name = data.Skills[i].Name

		if name == "go" || name == "golang" || name == "backend" || name == "java" || name == "python" || name == "csharp" || name == "nodejs" || name == "restapi" || name == "spring" || name == "laravel" || name == "pengembang_backend" {
			backend += 1
		}

		if name == "html" || name == "css" || name == "javascript" || name == "react" || name == "vuejs" || name == "angular" || name == "typescript" || name == "nextjs" || name == "tailwind" || name == "pengembang_frontend" {
			frontend += 1
		}

		if name == "sql" || name == "excel" || name == "analysis" || name == "machine_learning" || name == "data_mining" || name == "statistics" || name == "r" || name == "python_data" || name == "powerbi" || name == "tableau" || name == "analisis_data" {
			analyst += 1
		}

		if name == "writing" || name == "creativewriting" || name == "content_creation" || name == "copywriting" || name == "editing" || name == "menulis" || name == "penulis" || name == "konten" {
			writer += 1
		}

		if name == "teaching" || name == "publicspeaking" || name == "curriculum_design" || name == "education" || name == "training" || name == "mentoring" || name == "mengajar" || name == "guru" {
			teacher += 1
		}

		if name == "marketing" || name == "sales" || name == "digital_marketing" || name == "seo" || name == "social_media" || name == "brand_management" || name == "ads" || name == "market_research" || name == "pemasaran" || name == "penjualan" {
			marketing += 1
		}

		if name == "design" || name == "photoshop" || name == "illustrator" || name == "graphic_design" || name == "ux" || name == "ui" || name == "figma" || name == "adobe_xd" || name == "desain_grafis" {
			designer += 1
		}

		if name == "docker" || name == "kubernetes" || name == "linux" || name == "ci_cd" || name == "jenkins" || name == "devops" || name == "infrastructure" || name == "terraform" {
			devops += 1
		}

		if name == "android" || name == "ios" || name == "flutter" || name == "kotlin" || name == "swift" || name == "react_native" || name == "mobile_dev" || name == "pengembang_mobile" {
			mobile += 1
		}

		if name == "cybersecurity" || name == "network_security" || name == "penetration_testing" || name == "ethical_hacking" || name == "firewall" || name == "encryption" || name == "keamanan_siber" {
			security += 1
		}

		if name == "project_management" || name == "scrum" || name == "agile" || name == "kanban" || name == "pmp" || name == "jira" || name == "planning" || name == "manajemen_proyek" {
			management += 1
		}

		if name == "finance" || name == "accounting" || name == "budgeting" || name == "financial_analysis" || name == "tax" || name == "bookkeeping" || name == "akuntansi" || name == "keuangan" {
			finance += 1
		}

		if name == "customer_service" || name == "support" || name == "call_center" || name == "crm" || name == "client_relations" || name == "layanan_pelanggan" {
			customer_service += 1
		}

		if name == "game_dev" || name == "unity" || name == "unreal" || name == "gamemaker" || name == "godot" || name == "game_design" || name == "pengembang_game" {
			game += 1
		}

		if name == "video_editing" || name == "premiere" || name == "after_effects" || name == "davinci_resolve" || name == "final_cut" || name == "motion_graphics" || name == "editor_video" {
			video += 1
		}

		if name == "law" || name == "legal" || name == "paralegal" || name == "contracts" || name == "litigation" || name == "compliance" || name == "legal_research" || name == "hukum" || name == "pengacara" {
			law += 1
		}

		if name == "nursing" || name == "healthcare" || name == "patient_care" || name == "clinical" || name == "public_health" || name == "medical_assistant" || name == "perawat" || name == "kesehatan" {
			health += 1
		}

		if name == "admin" || name == "administration" || name == "data_entry" || name == "secretary" || name == "office_management" || name == "filing" || name == "admin_staf" || name == "sekretaris" {
			admin += 1
		}

		if name == "logistics" || name == "supply_chain" || name == "warehouse" || name == "inventory" || name == "transportation" || name == "procurement" || name == "logistik" || name == "gudang" || name == "pengiriman" {
			logistics += 1
		}

		if name == "human_resources" || name == "recruitment" || name == "talent_acquisition" || name == "payroll" || name == "hr_policy" || name == "training_development" || name == "hrd" || name == "sdm" {
			hr += 1
		}

		if name == "psychology" || name == "counseling" || name == "mental_health" || name == "therapy" || name == "social_work" || name == "psikologi" || name == "konselor" {
			psychology += 1
		}

		if name == "chef" || name == "cooking" || name == "kitchen" || name == "food" || name == "baking" || name == "catering" || name == "juru_masak" || name == "koki" {
			chef += 1
		}

		if name == "event_planning" || name == "event_management" || name == "wedding_planning" || name == "coordination" || name == "acara" || name == "event_organizer" {
			events += 1
		}

		if name == "kasir" || name == "cashier" || name == "penjaga_toko" {
			cashier += 1
		}

		if name == "operator_produksi" || name == "produksi" || name == "pekerja_pabrik" {
			operator += 1
		}

		if name == "barista" || name == "kopi" || name == "waiter" || name == "waitress" || name == "pelayan" {
			fnb += 1
		}

		if name == "content_creator" || name == "influencer" || name == "youtuber" || name == "tiktok" || name == "pembuat_konten" {
			creator += 1
		}

	}

	for i = 0; i < data.ExperienceCount; i++ {
		title = data.Experiences[i].Title
		company = data.Experiences[i].Company
		if company == "telkom_indonesia" || company == "bukalapak" || company == "tokopedia" || company == "shopee" || company == "gojek" ||
			company == "pertamina" || company == "bank_mandiri" || company == "bca" || company == "bni" || company == "indosat" || company == "xl" ||
			company == "astra_international" || company == "unilever_indonesia" || company == "indomie" || company == "indofood" || company == "mayora" ||
			company == "djarum" || company == "kalbe_farmina" || company == "mandiri_sekuritas" || company == "bpjs_kesehatan" || company == "bpjs_ketenagakerjaan" ||
			company == "bank_bri" || company == "pegadaian" || company == "wings_group" || company == "kimia_farma" || company == "blue_bird" ||
			company == "traveloka" || company == "tiket_com" || company == "pln" || company == "pelni" || company == "garuda_indonesia" ||
			company == "citilink" || company == "angkasapura" || company == "kai" || company == "pertani" || company == "bumn" ||
			company == "antarakita" || company == "sampoerna" || company == "gudang_garam" || company == "lippo_group" || company == "sinarmas" ||
			company == "pakuwon_group" || company == "agung_podomoro" || company == "summarecon" || company == "maspion" || company == "adaro" ||
			company == "freeport_indonesia" || company == "mnc_group" || company == "transcorp" || company == "net_tv" || company == "tvri" ||
			company == "kompas_gramedia" || company == "detik_com" || company == "tribunnews" || company == "kumparan" || company == "liputan6" {
			companyScore += 5
		} else if company != "" {
			companyScore += 1
		} else {
			companyScore += 0
		}

		if title == "backend_developer" || title == "backend_engineer" || title == "server_engineer" || title == "api_developer" {
			backend += 1
		}

		if title == "frontend_developer" || title == "frontend_engineer" || title == "web_developer" || title == "react_developer" || title == "vue_developer" {
			frontend += 1
		}

		if title == "fullstack_developer" || title == "fullstack_engineer" {
			fullstack += 1
		}

		if title == "mobile_developer" || title == "android_developer" || title == "ios_developer" || title == "flutter_developer" {
			mobile += 1
		}

		if title == "data_analyst" || title == "business_analyst" || title == "data_scientist" || title == "data_engineer" {
			analyst += 1
		}

		if title == "writer" || title == "content_writer" || title == "editor" || title == "copywriter" || title == "creative_writer" {
			writer += 1
		}

		if title == "teacher" || title == "lecturer" || title == "instructor" || title == "guru" || title == "pengajar" || title == "tutor" {
			teacher += 1
		}

		if title == "marketing_specialist" || title == "sales_executive" || title == "digital_marketer" || title == "seo_specialist" || title == "social_media_specialist" {
			marketing += 1
		}

		if title == "graphic_designer" || title == "ui_designer" || title == "ux_designer" || title == "visual_designer" || title == "motion_graphic_designer" {
			designer += 1
		}

		if title == "devops_engineer" || title == "site_reliability_engineer" || title == "infrastructure_engineer" {
			devops += 1
		}

		if title == "system_administrator" || title == "network_engineer" || title == "it_support" {
			it += 1
		}

		if title == "product_manager" || title == "project_manager" || title == "scrum_master" {
			management += 1
		}

		if title == "finance_staff" || title == "accountant" || title == "financial_analyst" || title == "auditor" {
			finance += 1
		}

		if title == "customer_service" || title == "call_center_agent" || title == "client_support" {
			customer_service += 1
		}

		if title == "hr_specialist" || title == "recruiter" || title == "hr_manager" {
			hr += 1
		}

		if title == "lawyer" || title == "legal_staff" || title == "legal_advisor" {
			law += 1
		}

		if title == "nurse" || title == "doctor" || title == "medical_assistant" || title == "healthcare_worker" {
			health += 1
		}

		if title == "admin_staff" || title == "administrative_assistant" || title == "secretary" {
			admin += 1
		}

		if title == "warehouse_staff" || title == "logistics_officer" || title == "supply_chain_specialist" || title == "delivery_staff" {
			logistics += 1
		}

		if title == "event_organizer" || title == "event_planner" || title == "event_coordinator" {
			events += 1
		}

		if title == "cashier" || title == "storekeeper" || title == "shop_assistant" {
			cashier += 1
		}

		if title == "production_operator" || title == "factory_worker" || title == "machine_operator" {
			operator += 1
		}

		if title == "barista" || title == "waiter" || title == "kitchen_staff" || title == "cook" || title == "chef" {
			fnb += 1
		}

		if title == "content_creator" || title == "influencer" || title == "youtuber" || title == "streamer" {
			creator += 1
		}

		if title == "psychologist" || title == "counselor" || title == "therapist" {
			psychology += 1
		}

		if title == "video_editor" || title == "videographer" || title == "motion_editor" {
			video += 1
		}

		if title == "game_developer" || title == "game_designer" || title == "game_programmer" {
			game += 1
		}
	}

	for j = 0; j < data.EducationCount; j++ {
		school = data.Educations[j].School
		degree = data.Educations[j].Degree
		major = data.Educations[j].Major

		// Data Source: https://institutbanten.ac.id/26-universitas-terbaik-di-indonesia-tahun-2025-versi-qs-world-university-rankings/
		if school == "universitas_indonesia" || school == "ui" ||
			school == "gajah_mada_university" || school == "universitas_gajah_mada" || school == "ugm" ||
			school == "bandung_institute_of_technology" || school == "institut_teknologi_bandung" || school == "itb" ||
			school == "airlangga_university" || school == "universitas_airlangga" || school == "unair" ||
			school == "ipb_university" || school == "bogor_agricultural_university" || school == "ipb" {
			educationScore += 5
		} else if school == "institut_teknologi_sepuluh_nopember" || school == "its" ||
			school == "universitas_padjadjaran" || school == "unpad" ||
			school == "diponegoro_university" || school == "universitas_diponegoro" || school == "undip" ||
			school == "universitas_brawijaya" || school == "ub" ||
			school == "bina_nusantara_university" || school == "binus" {
			educationScore += 3
		} else if school == "telkom_university" || school == "telkom" ||
			school == "universitas_hasanuddin" || school == "unhas" ||
			school == "universitas_sebelas_maret" || school == "uns" ||
			school == "atma_jaya_catholic_university_jakarta" || school == "atma_jaya" ||
			school == "universitas_islam_indonesia" || school == "uii" ||
			school == "universitas_muhammadiyah_yogyakarta" || school == "umy" ||
			school == "universitas_pendidikan_indonesia" || school == "upi" ||
			school == "universitas_sumatera_utara" || school == "usu" ||
			school == "yogyakarta_state_university" || school == "universitas_negeri_yogyakarta" || school == "uny" {
			educationScore += 2
		} else if school == "state_university_of_malang" || school == "universitas_negeri_malang" || school == "um" ||
			school == "udayana_university" || school == "universitas_udayana" ||
			school == "universitas_andalas" || school == "unand" ||
			school == "universitas_kristen_petra" || school == "ukp" ||
			school == "universitas_muhammadiyah_surakarta" || school == "ums" ||
			school == "university_of_lampung" || school == "universitas_lampung" || school == "unila" ||
			school == "university_of_mataram" || school == "universitas_mataram" || school == "unram" {
			educationScore += 1
		} else {
			educationScore += 0
		}

		if degree == "d1" || degree == "d2" || degree == "d3" {
			educationScore += 1
		} else if degree == "d4" || degree == "s1" {
			educationScore += 3
		} else if degree == "s2" {
			educationScore += 4
		} else if degree == "s3" {
			educationScore += 5
		} else {
			educationScore += 0
		}

		if major == "pendidikan" {
			teacher += 1
		} else if major == "hukum" {
			law += 1
		} else if major == "ekonomi" {
			finance += 1
		} else if major == "akuntansi" {
			finance += 1
		} else if major == "manajemen" {
			management += 1
		} else if major == "komputer" {
			backend += 1
			frontend += 1
			fullstack += 1
			analyst += 1
		} else if major == "teknik" {
			operator += 1
			designer += 1
		} else if major == "sains" {
			analyst += 1
		} else if major == "sosial" {
			analyst += 1
		} else if major == "seni" {
			designer += 1
			writer += 1
		} else if major == "sastra" {
			writer += 1
		}
	}

	fmt.Println("═════════════════════════════════════════════════════════════════════════════")
	fmt.Println(">> Saran AI:")
	fmt.Println("═════════════════════════════════════════════════════════════════════════════")

	total = backend + frontend + analyst + writer + teacher + marketing + designer + devops + mobile +
		security + management + finance + customer_service + game + video + law + health + admin +
		logistics + hr + psychology + chef + events + cashier + operator + fnb + creator + educationScore

	if total == 0 {
		fmt.Println()
		fmt.Println("[!!] Tidak ada rekomendasi. Silakan lengkapi data profil Anda")
	} else {
		fmt.Println()
		fmt.Println("[!!] AI telah menganalisis profil Anda. Berdasarkan skill, pengalaman, dan pendidikan,")
		fmt.Println("[!!] berikut jalur karier dan perusahaan terbaik yang direkomendasikan untuk Anda:")
		fmt.Println()

		// TODO: Backend Sec
		if backend >= 5 {
			fmt.Println("[Perusahaan Besar] Backend Developer di Tokopedia, Gojek, Bukalapak")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", backend)
			fmt.Println("[••] Saran: Tonjolkan pengalaman Anda dengan sistem skala besar dan microservices")
			fmt.Println()
		} else if backend >= 3 {
			fmt.Println("[Perusahaan Menengah] Backend Developer di Ruangguru, Qasir, Mamikos")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", backend)
			fmt.Println("[••] Saran: Tonjolkan pengalaman Anda dengan sistem skala besar dan microservices")
			fmt.Println()
		} else if backend > 0 {
			fmt.Println("[Startup] Backend Developer di LocalTech, NextDev")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", backend)
			fmt.Println("[••] Saran: Tonjolkan pengalaman Anda dengan sistem skala besar dan microservices")
			fmt.Println()
		}

		// TODO: Frontend Sec
		if frontend >= 5 {
			fmt.Println("[Perusahaan Besar] Frontend Developer di Traveloka, Blibli, Gojek")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", frontend)
			fmt.Println("[••] Saran: Tingkatkan skill UI/UX dan gunakan React/Vue di portofolio Anda")
			fmt.Println()
		} else if frontend >= 3 {
			fmt.Println("[Perusahaan Menengah] Frontend Developer di KlikDokter, Sayurbox")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", frontend)
			fmt.Println("[••] Saran: Tingkatkan skill UI/UX dan gunakan React/Vue di portofolio Anda")
			fmt.Println()
		} else if frontend > 0 {
			fmt.Println("[Startup] Frontend Developer di CodeCraft, DevStart")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", frontend)
			fmt.Println("[••] Saran: Tingkatkan skill UI/UX dan gunakan React/Vue di portofolio Anda")
			fmt.Println()
		}

		// TODO: Data Analyst
		if analyst >= 5 {
			fmt.Println("[Perusahaan Besar] Data Analyst di Telkom Indonesia, Shopee, Bank BCA")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", analyst)
			fmt.Println("[••] Saran: Latihan SQL dan dashboard dengan Power BI atau Tableau")
			fmt.Println()
		} else if analyst >= 3 {
			fmt.Println("[Perusahaan Menengah] Data Analyst di Xendit, Sociolla")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", analyst)
			fmt.Println("[••] Saran: Latihan SQL dan dashboard dengan Power BI atau Tableau.")
			fmt.Println()
		} else if analyst > 0 {
			fmt.Println("[Startup] Data Analyst di AnalytIQ, StatLab")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", analyst)
			fmt.Println("[••] Saran: Latihan SQL dan dashboard dengan Power BI atau Tableau.")
			fmt.Println()
		}

		// TODO: Content Writer / Editor
		if writer >= 5 {
			fmt.Println("[Perusahaan Besar] Content Writer / Editor di Kumparan, IDN Media, Zenius")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", writer)
			fmt.Println("[••] Saran: Bangun portofolio dan perkuat penulisan SEO")
			fmt.Println()
		} else if writer >= 3 {
			fmt.Println("[Perusahaan Menengah] Content Writer / Editor di Hipwee, FemaleDaily")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", writer)
			fmt.Println("[••] Saran: Bangun portofolio dan perkuat penulisan SEO")
			fmt.Println()
		} else if writer > 0 {
			fmt.Println("[Startup] Content Writer / Editor di BlogHive, LocalStory")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", writer)
			fmt.Println("[••] Saran: Bangun portofolio dan perkuat penulisan SEO")
			fmt.Println()
		}

		// TODO: Teacher / Lecturer
		if teacher >= 5 {
			fmt.Println("[Perusahaan Besar] Guru / Dosen di Binus, Universitas Indonesia, Ruangguru")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", teacher)
			fmt.Println("[••] Saran: Tampilkan materi ajar dan dapatkan sertifikasi")
			fmt.Println()
		} else if teacher >= 3 {
			fmt.Println("[Perusahaan Menengah] Guru / Dosen di HarukaEdu, Zenius")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", teacher)
			fmt.Println("[••] Saran: Tampilkan materi ajar dan dapatkan sertifikasi")
			fmt.Println()
		} else if teacher > 0 {
			fmt.Println("[Startup] Guru / Dosen di KelasKita, EduStart")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", teacher)
			fmt.Println("[••] Saran: Tampilkan materi ajar dan dapatkan sertifikasi")
			fmt.Println()
		}

		// TODO: Marketing Specialist
		if marketing >= 5 {
			fmt.Println("[Perusahaan Besar] Marketing Specialist di Grab, Shopee, TikTok Indonesia")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", marketing)
			fmt.Println("[••] Saran: Tunjukkan kampanye media sosial dan pengetahuan analitik")
			fmt.Println()
		} else if marketing >= 3 {
			fmt.Println("[Perusahaan Menengah] Marketing Specialist di Evermos, Sirclo")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", marketing)
			fmt.Println("[••] Saran: Tunjukkan kampanye media sosial dan pengetahuan analitik")
			fmt.Println()
		} else if marketing > 0 {
			fmt.Println("[Startup] Marketing Specialist di PromoPilot, StartSell")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", marketing)
			fmt.Println("[••] Saran: Tunjukkan kampanye media sosial dan pengetahuan analitik")
			fmt.Println()
		}

		// TODO: Graphic Designer
		if designer >= 5 {
			fmt.Println("[Perusahaan Besar] Graphic Designer di Tokopedia, Bukalapak, Netflix ID")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", designer)
			fmt.Println("[••] Saran: Perbaiki portofolio Behance/Dribbble Anda")
			fmt.Println()
		} else if designer >= 3 {
			fmt.Println("[Perusahaan Menengah] Graphic Designer di Jakmall, Warung Pintar")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", designer)
			fmt.Println("[••] Saran: Perbaiki portofolio Behance/Dribbble Anda")
			fmt.Println()
		} else if designer > 0 {
			fmt.Println("[Startup] Graphic Designer di PixelCraft, LokalDesign")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", designer)
			fmt.Println("[••] Saran: Perbaiki portofolio Behance/Dribbble Anda")
			fmt.Println()
		}

		// TODO: DevOps Engineer
		if devops >= 5 {
			fmt.Println("[Perusahaan Besar] DevOps Engineer di Blibli, Gojek, JNE")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", devops)
			fmt.Println("[••] Saran: Tunjukkan pipeline CI/CD dan infrastructure as code")
			fmt.Println()
		} else if devops >= 3 {
			fmt.Println("[Perusahaan Menengah] DevOps Engineer di Alterra, Nodeflux")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", devops)
			fmt.Println("[••] Saran: Tunjukkan pipeline CI/CD dan infrastructure as code")
			fmt.Println()
		} else if devops > 0 {
			fmt.Println("[Startup] DevOps Engineer di InfraEdge, LiteOps")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", devops)
			fmt.Println("[••] Saran: Tunjukkan pipeline CI/CD dan infrastructure as code")
			fmt.Println()
		}

		// TODO: Mobile App Developer
		if mobile >= 5 {
			fmt.Println("[Perusahaan Besar] Mobile App Developer di Traveloka, TikTok, Tokopedia")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", mobile)
			fmt.Println("[••] Saran: Bangun proyek Flutter/React Native")
			fmt.Println()
		} else if mobile >= 3 {
			fmt.Println("[Perusahaan Menengah] Mobile App Developer di Kulina, MySkill")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", mobile)
			fmt.Println("[••] Saran: Bangun proyek Flutter/React Native")
			fmt.Println()
		} else if mobile > 0 {
			fmt.Println("[Startup] Mobile App Developer di AppHero, NextMob")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", mobile)
			fmt.Println("[••] Saran: Bangun proyek Flutter/React Native")
			fmt.Println()
		}

		// TODO: Cybersecurity Specialist
		if security >= 5 {
			fmt.Println("[Perusahaan Besar] Cybersecurity Specialist di Bank Indonesia, Tokopedia, Kominfo")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", security)
			fmt.Println("[••] Saran: Ambil sertifikasi keamanan siber dan ikut CTF")
			fmt.Println()
		} else if security >= 3 {
			fmt.Println("[Perusahaan Menengah] Cybersecurity Specialist di Vaksincom, Privy ID")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", security)
			fmt.Println("[••] Saran: Ambil sertifikasi keamanan siber dan ikut CTF")
			fmt.Println()
		} else if security > 0 {
			fmt.Println("[Startup] Cybersecurity Specialist di SecuStart, ByteSecure")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", security)
			fmt.Println("[••] Saran: Ambil sertifikasi keamanan siber dan ikut CTF")
			fmt.Println()
		}

		// TODO: Project Manager
		if management >= 5 {
			fmt.Println("[Perusahaan Besar] Project Manager di Telkomsel, Pertamina, BRI")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", management)
			fmt.Println("[••] Saran: Tonjolkan perencanaan proyek, kepemimpinan, dan metode agile")
			fmt.Println()
		} else if management >= 3 {
			fmt.Println("[Perusahaan Menengah] Project Manager di HappyFresh, Pinhome")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", management)
			fmt.Println("[••] Saran: Tonjolkan perencanaan proyek, kepemimpinan, dan metode agile")
			fmt.Println()
		} else if management > 0 {
			fmt.Println("[Startup] Project Manager di ManageOne, AgileStudio")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", management)
			fmt.Println("[••] Saran: Tonjolkan perencanaan proyek, kepemimpinan, dan metode agile")
			fmt.Println()
		}

		// TODO: Financial Analyst
		if finance >= 5 {
			fmt.Println("[Perusahaan Besar] Financial Analyst di Bank Mandiri, OJK, PwC")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", finance)
			fmt.Println("[••] Saran: Asah skill Excel dan financial modeling")
			fmt.Println()
		} else if finance >= 3 {
			fmt.Println("[Perusahaan Menengah] Financial Analyst di Bibit, TernakUang")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", finance)
			fmt.Println("[••] Saran: Asah skill Excel dan financial modeling")
			fmt.Println()
		} else if finance > 0 {
			fmt.Println("[Startup] Financial Analyst di FinPilot, BudgetBuddy")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", finance)
			fmt.Println("[••] Saran: Asah skill Excel dan financial modeling")
			fmt.Println()
		}

		// TODO: Customer Service
		if customer_service >= 5 {
			fmt.Println("[Perusahaan Besar] Customer Service di Shopee, Tokopedia, Traveloka")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", customer_service)
			fmt.Println("[••] Saran: Tingkatkan kemampuan komunikasi dan empati")
			fmt.Println()
		} else if customer_service >= 3 {
			fmt.Println("[Perusahaan Menengah] Customer Service di Sirclo, Evermos")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", customer_service)
			fmt.Println("[••] Saran: Tingkatkan kemampuan komunikasi dan empati")
			fmt.Println()
		} else if customer_service > 0 {
			fmt.Println("[Startup] Customer Service di HelpLine, FastDesk")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", customer_service)
			fmt.Println("[••] Saran: Tingkatkan kemampuan komunikasi dan empati")
			fmt.Println()
		}

		// TODO: Game Developer
		if game >= 5 {
			fmt.Println("[Perusahaan Besar] Game Developer di ArenaNet ID, Agate Studio, Garena")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", game)
			fmt.Println("[••] Saran: Kuasai Unity atau Unreal dan buat game kecil")
			fmt.Println()
		} else if game >= 3 {
			fmt.Println("[Perusahaan Menengah] Game Developer di Toge Productions, Digital Happiness")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", game)
			fmt.Println("[••] Saran: Kuasai Unity atau Unreal dan buat game kecil")
			fmt.Println()
		} else if game > 0 {
			fmt.Println("[Startup] Game Developer di PixelPlay, IndieDev")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", game)
			fmt.Println("[••] Saran: Kuasai Unity atau Unreal dan buat game kecil")
			fmt.Println()
		}

		//TODO:  Video Editor
		if video >= 5 {
			fmt.Println("[Perusahaan Besar] Video Editor di NET TV, Narasi, IDN Media")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", video)
			fmt.Println("[••] Saran: Buat showreel dan pelajari motion graphics")
			fmt.Println()
		} else if video >= 3 {
			fmt.Println("[Perusahaan Menengah] Video Editor di Kreator.ID, LokalTV")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", video)
			fmt.Println("[••] Saran: Buat showreel dan pelajari motion graphics")
			fmt.Println()
		} else if video > 0 {
			fmt.Println("[Startup] Video Editor di VideoVibe, ClipStart")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", video)
			fmt.Println("[••] Saran: Buat showreel dan pelajari motion graphics")
			fmt.Println()
		}

		//TODO:  Legal Advisor
		if law >= 5 {
			fmt.Println("[Perusahaan Besar] Legal Advisor di Mahkamah Agung, Kantor Hukum Ternama")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", law)
			fmt.Println("[••] Saran: Perdalam pengetahuan regulasi dan penulisan hukum")
		} else if law >= 3 {
			fmt.Println("[Perusahaan Menengah] Legal Advisor di Startup Hukum Lokal")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", law)
			fmt.Println("[••] Saran: Perdalam pengetahuan regulasi dan penulisan hukum")
		} else if law > 0 {
			fmt.Println("[Startup] Legal Advisor di LegalEase, LawEntry")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", law)
			fmt.Println("[••] Saran: Perdalam pengetahuan regulasi dan penulisan hukum")
		}

		//TODO: Healthcare Worker
		if health >= 5 {
			fmt.Println("[Perusahaan Besar] Tenaga Kesehatan di RSUPN, Siloam, Halodoc")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", health)
			fmt.Println("[••] Saran: Tampilkan sertifikasi dan pengalaman merawat pasien")
		} else if health >= 3 {
			fmt.Println("[Perusahaan Menengah] Tenaga Kesehatan di Good Doctor, ProSehat")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", health)
			fmt.Println("[••] Saran: Tampilkan sertifikasi dan pengalaman merawat pasien")
		} else if health > 0 {
			fmt.Println("[Startup] Tenaga Kesehatan di HealthStart, KlinikKita")
			fmt.Printf("[!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", health)
			fmt.Println("[••] Saran: Tampilkan sertifikasi dan pengalaman merawat pasien")
		}

		//TODO: Administrative Staff
		if admin >= 5 {
			fmt.Println("[Perusahaan Besar] Staf Administrasi di Pertamina, PLN, Bank Mandiri")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", admin)
			fmt.Println("[••] Saran: Perkuat skill MS Office dan organisasi")
		} else if admin >= 3 {
			fmt.Println("[Perusahaan Menengah] Staf Administrasi di DANA, OVO")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", admin)
			fmt.Println("[••] Saran: Perkuat skill MS Office dan organisasi")
		} else if admin > 0 {
			fmt.Println("[Startup] Staf Administrasi di OfficeHub, AdminPro")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", admin)
			fmt.Println("[••] Saran: Perkuat skill MS Office dan organisasi")
		}

		//TODO: Logistics Coordinator
		if logistics >= 5 {
			fmt.Println("[Perusahaan Besar] Koordinator Logistik di JNE, TIKI, Gojek")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", logistics)
			fmt.Println("[••] Saran: Fokus pada skill manajemen rantai pasok")
		} else if logistics >= 3 {
			fmt.Println("[Perusahaan Menengah] Koordinator Logistik di Deliveree, Lalamove")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", logistics)
			fmt.Println("[••] Saran: Fokus pada skill manajemen rantai pasok")
		} else if logistics > 0 {
			fmt.Println("[Startup] Koordinator Logistik di FastLog, QuickShip")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", logistics)
			fmt.Println("[••] Saran: Fokus pada skill manajemen rantai pasok")
		}

		//TODO:  HR Specialist
		if hr >= 5 {
			fmt.Println("[Perusahaan Besar] HR Specialist di Unilever, Telkom Indonesia, Danone")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", hr)
			fmt.Println("[••] Saran: Kembangkan skill interview dan hubungan karyawan")
		} else if hr >= 3 {
			fmt.Println("[Perusahaan Menengah] HR Specialist di Modalku, Kredivo")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", hr)
			fmt.Println("[••] Saran: Kembangkan skill interview dan hubungan karyawan")
		} else if hr > 0 {
			fmt.Println("[Startup] HR Specialist di PeopleOps, HRStart")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", hr)
			fmt.Println("[••] Saran: Kembangkan skill interview dan hubungan karyawan")
		}

		//TODO:  Psychologist / Counselor
		if psychology >= 5 {
			fmt.Println("[Perusahaan Besar] Psikolog / Konselor di RSJ, Klinik Pratama, Sekolah")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", psychology)
			fmt.Println("[••] Saran: Tampilkan sertifikasi konseling dan pengalaman")
		} else if psychology >= 3 {
			fmt.Println("[Perusahaan Menengah] Psikolog / Konselor di Klinik Prima, Sekolah Alam")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", psychology)
			fmt.Println("[••] Saran: Tampilkan sertifikasi konseling dan pengalaman")
		} else if psychology > 0 {
			fmt.Println("[Startup] Psikolog / Konselor di MindCare, SafeTalk")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", psychology)
			fmt.Println("[••] Saran: Tampilkan sertifikasi konseling dan pengalaman")
		}

		//TODO: Chef / Cook
		if chef >= 5 {
			fmt.Println("[Perusahaan Besar] Koki / Juru Masak di Hotel Mulia, Ritz Carlton, Resto Lokal")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", chef)
			fmt.Println("[••] Saran: Tampilkan sertifikasi kuliner dan menu andalan")
		} else if chef >= 3 {
			fmt.Println("[Perusahaan Menengah] Koki / Juru Masak di Hotel Santika, Bistro Lokal")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", chef)
			fmt.Println("[••] Saran: Tampilkan sertifikasi kuliner dan menu andalan")
		} else if chef > 0 {
			fmt.Println("[Startup] Koki / Juru Masak di FoodieStart, HomeKitchen")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", chef)
			fmt.Println("[••] Saran: Tampilkan sertifikasi kuliner dan menu andalan")
		}

		//TODO:  Event Organizer
		if events >= 5 {
			fmt.Println("[Perusahaan Besar] Event Organizer di Dyandra, Rajawali Corpora")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", events)
			fmt.Println("[••] Saran: Tampilkan portofolio event dan skill negosiasi")
		} else if events >= 3 {
			fmt.Println("[Perusahaan Menengah] Event Organizer di EO Lokal, PartyUp")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", events)
			fmt.Println("[••] Saran: Tampilkan portofolio event dan skill negosiasi")
		} else if events > 0 {
			fmt.Println("[Startup] Event Organizer di SmallEvents, FestStart")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", events)
			fmt.Println("[••] Saran: Tampilkan portofolio event dan skill negosiasi")
		}

		//TODO: Cashier
		if cashier >= 5 {
			fmt.Println("[Perusahaan Besar] Kasir di Indomaret, Alfamart, Circle K")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", cashier)
			fmt.Println("[••] Saran: Fokus pada ketelitian dan skill layanan pelanggan")
		} else if cashier >= 3 {
			fmt.Println("[Perusahaan Menengah] Kasir di Minimart Lokal, Toko Sembako")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", cashier)
			fmt.Println("[••] Saran: Fokus pada ketelitian dan skill layanan pelanggan")
		} else if cashier > 0 {
			fmt.Println("[Startup] Kasir di LocalStore, QuickMart")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", cashier)
			fmt.Println("[••] Saran: Fokus pada ketelitian dan skill layanan pelanggan")
		}

		//TODO: Machine Operator
		if operator >= 5 {
			fmt.Println("[Perusahaan Besar] Operator Mesin di Indofood, Unilever, Astra")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", operator)
			fmt.Println("[••] Saran: Tonjolkan sertifikasi penanganan mesin")
		} else if operator >= 3 {
			fmt.Println("[Perusahaan Menengah] Operator Mesin di LocalFactory, PlantX")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", operator)
			fmt.Println("[••] Saran: Tonjolkan sertifikasi penanganan mesin")
		} else if operator > 0 {
			fmt.Println("[Startup] Operator Mesin di SmallPlant, MachPro")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", operator)
			fmt.Println("[••] Saran: Tonjolkan sertifikasi penanganan mesin")
		}

		//TODO: Food & Beverage Staff section
		if fnb >= 5 {
			fmt.Println("[Perusahaan Besar] Staf F&B di Hotel Indonesia Kempinski, Marriott")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", fnb)
			fmt.Println("[••] Saran: Tonjolkan skill pelayanan dan sertifikasi")
		} else if fnb >= 3 {
			fmt.Println("[Perusahaan Menengah] Staf F&B di Local Cafe, Bistro X")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", fnb)
			fmt.Println("[••] Saran: Tonjolkan skill pelayanan dan sertifikasi")
		} else if fnb > 0 {
			fmt.Println("[Startup] Staf F&B di QuickBite, FoodStart")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", fnb)
			fmt.Println("[••] Saran: Tonjolkan skill pelayanan dan sertifikasi")
		}

		//TODO: Content Creator / Influencer section
		if creator >= 5 {
			fmt.Println("[Perusahaan Besar] Content Creator / Influencer di YouTube, TikTok, Instagram")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 75%% - 100%%\n", creator)
			fmt.Println("[••] Saran: Bangun konten konsisten dan basis pengikut")
		} else if creator >= 3 {
			fmt.Println("[Perusahaan Menengah] Content Creator / Influencer di Media Lokal, Blog")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 50%% - 75%%\n", creator)
			fmt.Println("[••] Saran: Bangun konten konsisten dan basis pengikut")
		} else if creator > 0 {
			fmt.Println("[Startup] Content Creator / Influencer di IndieVlog, MiniBlog")
			fmt.Printf("     [!!!] Poin: %d/10 | Peluang diterima: 15%% - 50%%\n", creator)
			fmt.Println("[••] Saran: Bangun konten konsisten dan basis pengikut")
		}

		fmt.Println("═════════════════════════════════════════════════════════════════════════════")
		fmt.Println("[INFO] Ini adalah saran berbasis AI. Terus tingkatkan profil Anda,")
		fmt.Println("       dan Career-Edge akan memperbarui peluang Anda secara real-time!")
		fmt.Println("═════════════════════════════════════════════════════════════════════════════")
	}
}
