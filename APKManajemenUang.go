package main

import "fmt"

const NMAX int = 100

// ================= STRUCT =================
type Income struct {
	nama    string
	jumlah  int
	tanggal int
	bulan   int
}

type Expense struct {
	nama    string
	jumlah  int
	tanggal int
	bulan   int
}

type Subscription struct {
	nama    string
	harga   int
	tanggal int
	bulan   int
	aktif   bool
}

// Struct bantu untuk hasil pencarian/filter yang sudah digabung & siap diurutkan
type HasilTransaksi struct {
	kategori string // "Income", "Expense", "Subscription"
	nama     string
	jumlah   int
	tanggal  int
	bulan    int
	aktif    bool // hanya relevan untuk Subscription
}

// ================= GLOBAL DATA =================
var incomeData [NMAX]Income
var expenseData [NMAX]Expense
var subData [NMAX]Subscription

var nIncome, nExpense, nSub int
var saldo int
var saldoAwal int

// ================= UPDATE SALDO =================
func updateSaldo() {
	totalIncome := 0
	totalExpense := 0
	totalSub := 0

	for i := 0; i < nIncome; i++ {
		totalIncome += incomeData[i].jumlah
	}
	for i := 0; i < nExpense; i++ {
		totalExpense += expenseData[i].jumlah
	}
	for i := 0; i < nSub; i++ {
		if subData[i].aktif {
			totalSub += subData[i].harga
		}
	}
	saldo = saldo + totalIncome - (totalExpense + totalSub)
}

// ================= INCOME =================
func tambahIncome() {
	if nIncome >= NMAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Nama income: ")
	fmt.Scan(&incomeData[nIncome].nama)

	fmt.Print("Jumlah: ")
	fmt.Scan(&incomeData[nIncome].jumlah)

	fmt.Print("Tanggal: ")
	fmt.Scan(&incomeData[nIncome].tanggal)

	fmt.Print("Bulan: ")
	fmt.Scan(&incomeData[nIncome].bulan)

	saldo += incomeData[nIncome].jumlah
	nIncome++
	fmt.Println("Income berhasil ditambahkan!")
}

func editIncome() {
	if nIncome == 0 {
		fmt.Println("Belum ada data income!")
		return
	}
	fmt.Println("=== DAFTAR INCOME ===")
	for i := 0; i < nIncome; i++ {
		fmt.Printf("%d. %s | Rp%d | %d/%d\n", i+1, incomeData[i].nama, incomeData[i].jumlah, incomeData[i].tanggal, incomeData[i].bulan)
	}
	var idx int
	fmt.Print("Pilih nomor income yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nIncome {
		fmt.Println("Nomor tidak valid!")
		return
	}
	saldo -= incomeData[idx].jumlah
	fmt.Print("Nama baru: ")
	fmt.Scan(&incomeData[idx].nama)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&incomeData[idx].jumlah)
	fmt.Print("Tanggal baru: ")
	fmt.Scan(&incomeData[idx].tanggal)
	fmt.Print("Bulan baru: ")
	fmt.Scan(&incomeData[idx].bulan)
	saldo += incomeData[idx].jumlah
	fmt.Println("Income berhasil diupdate!")
}

func hapusIncome() {
	if nIncome == 0 {
		fmt.Println("Belum ada data income!")
		return
	}
	fmt.Println("=== DAFTAR INCOME ===")
	for i := 0; i < nIncome; i++ {
		fmt.Printf("%d. %s | Rp%d | %d/%d\n", i+1, incomeData[i].nama, incomeData[i].jumlah, incomeData[i].tanggal, incomeData[i].bulan)
	}
	var idx int
	fmt.Print("Pilih nomor income yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nIncome {
		fmt.Println("Nomor tidak valid!")
		return
	}
	saldo -= incomeData[idx].jumlah
	for i := idx; i < nIncome-1; i++ {
		incomeData[i] = incomeData[i+1]
	}
	nIncome--
	fmt.Println("Income berhasil dihapus!")
}

// ================= EXPENSE =================
func tambahExpense() {
	if nExpense >= NMAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Nama expense: ")
	fmt.Scan(&expenseData[nExpense].nama)

	fmt.Print("Jumlah: ")
	fmt.Scan(&expenseData[nExpense].jumlah)

	fmt.Print("Tanggal: ")
	fmt.Scan(&expenseData[nExpense].tanggal)

	fmt.Print("Bulan: ")
	fmt.Scan(&expenseData[nExpense].bulan)

	saldo -= expenseData[nExpense].jumlah
	nExpense++
	fmt.Println("Expense berhasil ditambahkan!")
}

func editExpense() {
	if nExpense == 0 {
		fmt.Println("Belum ada data expense!")
		return
	}
	fmt.Println("=== DAFTAR EXPENSE ===")
	for i := 0; i < nExpense; i++ {
		fmt.Printf("%d. %s | Rp%d | %d/%d\n", i+1, expenseData[i].nama, expenseData[i].jumlah, expenseData[i].tanggal, expenseData[i].bulan)
	}
	var idx int
	fmt.Print("Pilih nomor expense yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nExpense {
		fmt.Println("Nomor tidak valid!")
		return
	}
	saldo += expenseData[idx].jumlah
	fmt.Print("Nama baru: ")
	fmt.Scan(&expenseData[idx].nama)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&expenseData[idx].jumlah)
	fmt.Print("Tanggal baru: ")
	fmt.Scan(&expenseData[idx].tanggal)
	fmt.Print("Bulan baru: ")
	fmt.Scan(&expenseData[idx].bulan)
	saldo -= expenseData[idx].jumlah
	fmt.Println("Expense berhasil diupdate!")
}

func hapusExpense() {
	if nExpense == 0 {
		fmt.Println("Belum ada data expense!")
		return
	}
	fmt.Println("=== DAFTAR EXPENSE ===")
	for i := 0; i < nExpense; i++ {
		fmt.Printf("%d. %s | Rp%d | %d/%d\n", i+1, expenseData[i].nama, expenseData[i].jumlah, expenseData[i].tanggal, expenseData[i].bulan)
	}
	var idx int
	fmt.Print("Pilih nomor expense yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nExpense {
		fmt.Println("Nomor tidak valid!")
		return
	}
	saldo += expenseData[idx].jumlah
	for i := idx; i < nExpense-1; i++ {
		expenseData[i] = expenseData[i+1]
	}
	nExpense--
	fmt.Println("Expense berhasil dihapus!")
}

// ================= SUBSCRIPTION =================
func tambahSubscription() {
	if nSub >= NMAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Nama subscription: ")
	fmt.Scan(&subData[nSub].nama)

	fmt.Print("Harga: ")
	fmt.Scan(&subData[nSub].harga)

	fmt.Print("Tanggal: ")
	fmt.Scan(&subData[nSub].tanggal)

	fmt.Print("Bulan: ")
	fmt.Scan(&subData[nSub].bulan)

	subData[nSub].aktif = true
	saldo -= subData[nSub].harga
	nSub++
	fmt.Println("Subscription berhasil ditambahkan!")
}

func editSubscription() {
	if nSub == 0 {
		fmt.Println("Belum ada data subscription!")
		return
	}
	fmt.Println("=== DAFTAR SUBSCRIPTION ===")
	for i := 0; i < nSub; i++ {
		status := "Aktif"
		if !subData[i].aktif {
			status = "Nonaktif"
		}
		fmt.Printf("%d. %s | Rp%d | %d/%d | %s\n", i+1, subData[i].nama, subData[i].harga, subData[i].tanggal, subData[i].bulan, status)
	}
	var idx int
	fmt.Print("Pilih nomor subscription yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nSub {
		fmt.Println("Nomor tidak valid!")
		return
	}
	if subData[idx].aktif {
		saldo += subData[idx].harga
	}
	fmt.Print("Nama baru: ")
	fmt.Scan(&subData[idx].nama)
	fmt.Print("Harga baru: ")
	fmt.Scan(&subData[idx].harga)
	fmt.Print("Tanggal baru: ")
	fmt.Scan(&subData[idx].tanggal)
	fmt.Print("Bulan baru: ")
	fmt.Scan(&subData[idx].bulan)
	var aktifInput int
	fmt.Print("Status aktif? (1=Ya, 0=Tidak): ")
	fmt.Scan(&aktifInput)
	subData[idx].aktif = aktifInput == 1
	if subData[idx].aktif {
		saldo -= subData[idx].harga
	}
	fmt.Println("Subscription berhasil diupdate!")
}

func hapusSubscription() {
	if nSub == 0 {
		fmt.Println("Belum ada data subscription!")
		return
	}
	fmt.Println("=== DAFTAR SUBSCRIPTION ===")
	for i := 0; i < nSub; i++ {
		status := "Aktif"
		if !subData[i].aktif {
			status = "Nonaktif"
		}
		fmt.Printf("%d. %s | Rp%d | %d/%d | %s\n", i+1, subData[i].nama, subData[i].harga, subData[i].tanggal, subData[i].bulan, status)
	}
	var idx int
	fmt.Print("Pilih nomor subscription yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nSub {
		fmt.Println("Nomor tidak valid!")
		return
	}
	if subData[idx].aktif {
		saldo += subData[idx].harga
	}
	for i := idx; i < nSub-1; i++ {
		subData[i] = subData[i+1]
	}
	nSub--
	fmt.Println("Subscription berhasil dihapus!")
}

func toggleSubscription() {
	if nSub == 0 {
		fmt.Println("Belum ada data subscription!")
		return
	}
	fmt.Println("=== DAFTAR SUBSCRIPTION ===")
	for i := 0; i < nSub; i++ {
		status := "Aktif"
		if !subData[i].aktif {
			status = "Nonaktif"
		}
		fmt.Printf("%d. %s | Rp%d | %s\n", i+1, subData[i].nama, subData[i].harga, status)
	}
	var idx int
	fmt.Print("Pilih nomor subscription untuk toggle aktif/nonaktif: ")
	fmt.Scan(&idx)
	idx--
	if idx < 0 || idx >= nSub {
		fmt.Println("Nomor tidak valid!")
		return
	}
	if subData[idx].aktif {
		subData[idx].aktif = false
		saldo += subData[idx].harga
		fmt.Println("Subscription dinonaktifkan.")
	} else {
		subData[idx].aktif = true
		saldo -= subData[idx].harga
		fmt.Println("Subscription diaktifkan.")
	}
}

// ================= DASHBOARD =================
func hitungDashboard() {
	totalIncome := 0
	totalExpense := 0
	totalSub := 0

	for i := 0; i < nIncome; i++ {
		totalIncome += incomeData[i].jumlah
	}
	for i := 0; i < nExpense; i++ {
		totalExpense += expenseData[i].jumlah
	}
	for i := 0; i < nSub; i++ {
		if subData[i].aktif {
			totalSub += subData[i].harga
		}
	}

	totalKeluar := totalExpense + totalSub
	// Rekonsiliasi: hitung ulang saldo dari saldoAwal agar selalu akurat
	saldo = saldoAwal + totalIncome - totalKeluar

	fmt.Println("===== DASHBOARD =====")
	fmt.Printf("Saldo Awal     : Rp%d\n", saldoAwal)
	fmt.Println("Total Income   :", totalIncome)
	fmt.Println("Total Expense  :", totalExpense)
	fmt.Println("Subscription   :", totalSub)
	fmt.Println("----------------------")
	fmt.Printf("Saldo Saat Ini : Rp%d\n", saldo)

	// Basis persentase: saldo awal + semua income
	totalPemasukan := saldoAwal + totalIncome
	if totalPemasukan == 0 {
		fmt.Println("Belum ada pemasukan sama sekali.")
		return
	}

	if totalKeluar > totalPemasukan {
		fmt.Println("WARNING: Kamu defisit (boros)")
	} else if totalKeluar > totalPemasukan*70/100 {
		fmt.Println("Hati-hati, pengeluaran tinggi")
	} else {
		fmt.Println("Keuangan aman")
	}
}

// ================= HELPER: SUBSTRING SEARCH =================
// Mengembalikan true jika keyword ditemukan di dalam teks (case-insensitive)
func mengandungKeyword(teks string, keyword string) bool {
	kw := []byte(keyword)
	nm := []byte(teks)
	if len(kw) == 0 {
		return true
	}
	if len(kw) > len(nm) {
		return false
	}
	for j := 0; j <= len(nm)-len(kw); j++ {
		match := true
		for k := 0; k < len(kw); k++ {
			a := kw[k]
			b := nm[j+k]
			if a >= 65 && a <= 90 {
				a += 32
			}
			if b >= 65 && b <= 90 {
				b += 32
			}
			if a != b {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// ================= HELPER: INSERTION SORT HASIL TRANSAKSI =================
// Mengurutkan slice HasilTransaksi berdasarkan bulan (ascending),
// lalu tanggal (ascending) jika bulan sama.
func insertionSortHasil(data []HasilTransaksi) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		// Geser elemen yang lebih besar ke kanan
		for j >= 0 && (data[j].bulan > key.bulan || (data[j].bulan == key.bulan && data[j].tanggal > key.tanggal)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// ================= HELPER: CETAK HASIL TRANSAKSI =================
func cetakHasil(data []HasilTransaksi) {
	if len(data) == 0 {
		fmt.Println("Tidak ada transaksi yang ditemukan.")
		return
	}
	bulanSekarang := -1
	for _, h := range data {
		// Cetak header bulan setiap ganti bulan
		if h.bulan != bulanSekarang {
			bulanSekarang = h.bulan
			fmt.Printf("\n--- Bulan %d ---\n", bulanSekarang)
		}
		if h.kategori == "Subscription" {
			status := "Aktif"
			if !h.aktif {
				status = "Nonaktif"
			}
			fmt.Printf("  [%s] %s | Rp%d | Tgl %d | %s\n", h.kategori, h.nama, h.jumlah, h.tanggal, status)
		} else {
			fmt.Printf("  [%s] %s | Rp%d | Tgl %d\n", h.kategori, h.nama, h.jumlah, h.tanggal)
		}
	}
}

// ================= CARI & FILTER TRANSAKSI (Menu 8) =================
func cariDanFilterTransaksi() {
	fmt.Println("\n=== CARI & FILTER TRANSAKSI ===")
	fmt.Println("1. Cari by Keyword")
	fmt.Println("2. Filter by Bulan")
	fmt.Println("3. Filter by Tanggal & Bulan")
	fmt.Print("Pilih: ")

	var sub int
	fmt.Scan(&sub)

	var hasil []HasilTransaksi

	if sub == 1 {
		// ---- CARI BY KEYWORD ----
		var keyword string
		fmt.Print("Masukkan kata kunci: ")
		fmt.Scan(&keyword)

		for i := 0; i < nIncome; i++ {
			if mengandungKeyword(incomeData[i].nama, keyword) {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Income",
					nama:     incomeData[i].nama,
					jumlah:   incomeData[i].jumlah,
					tanggal:  incomeData[i].tanggal,
					bulan:    incomeData[i].bulan,
				})
			}
		}
		for i := 0; i < nExpense; i++ {
			if mengandungKeyword(expenseData[i].nama, keyword) {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Expense",
					nama:     expenseData[i].nama,
					jumlah:   expenseData[i].jumlah,
					tanggal:  expenseData[i].tanggal,
					bulan:    expenseData[i].bulan,
				})
			}
		}
		for i := 0; i < nSub; i++ {
			if mengandungKeyword(subData[i].nama, keyword) {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Subscription",
					nama:     subData[i].nama,
					jumlah:   subData[i].harga,
					tanggal:  subData[i].tanggal,
					bulan:    subData[i].bulan,
					aktif:    subData[i].aktif,
				})
			}
		}

		fmt.Printf("\n=== HASIL PENCARIAN: \"%s\" ===", keyword)

	} else if sub == 2 {
		// ---- FILTER BY BULAN ----
		var bulan int
		fmt.Print("Masukkan bulan (1-12): ")
		fmt.Scan(&bulan)

		totalIncome := 0
		totalExpense := 0
		totalSub := 0

		for i := 0; i < nIncome; i++ {
			if incomeData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Income",
					nama:     incomeData[i].nama,
					jumlah:   incomeData[i].jumlah,
					tanggal:  incomeData[i].tanggal,
					bulan:    incomeData[i].bulan,
				})
				totalIncome += incomeData[i].jumlah
			}
		}
		for i := 0; i < nExpense; i++ {
			if expenseData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Expense",
					nama:     expenseData[i].nama,
					jumlah:   expenseData[i].jumlah,
					tanggal:  expenseData[i].tanggal,
					bulan:    expenseData[i].bulan,
				})
				totalExpense += expenseData[i].jumlah
			}
		}
		for i := 0; i < nSub; i++ {
			if subData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Subscription",
					nama:     subData[i].nama,
					jumlah:   subData[i].harga,
					tanggal:  subData[i].tanggal,
					bulan:    subData[i].bulan,
					aktif:    subData[i].aktif,
				})
				if subData[i].aktif {
					totalSub += subData[i].harga
				}
			}
		}

		fmt.Printf("\n=== FILTER BULAN %d ===", bulan)
		insertionSortHasil(hasil)
		cetakHasil(hasil)

		fmt.Println("\n--- Ringkasan Bulan", bulan, "---")
		fmt.Println("Total Income   :", totalIncome)
		fmt.Println("Total Expense  :", totalExpense)
		fmt.Println("Subscription   :", totalSub)
		fmt.Println("Net            :", totalIncome-(totalExpense+totalSub))
		return

	} else if sub == 3 {
		// ---- FILTER BY TANGGAL & BULAN ----
		var tanggal, bulan int
		fmt.Print("Masukkan tanggal (1-31): ")
		fmt.Scan(&tanggal)
		fmt.Print("Masukkan bulan (1-12): ")
		fmt.Scan(&bulan)

		totalIncome := 0
		totalExpense := 0
		totalSub := 0

		for i := 0; i < nIncome; i++ {
			if incomeData[i].tanggal == tanggal && incomeData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Income",
					nama:     incomeData[i].nama,
					jumlah:   incomeData[i].jumlah,
					tanggal:  incomeData[i].tanggal,
					bulan:    incomeData[i].bulan,
				})
				totalIncome += incomeData[i].jumlah
			}
		}
		for i := 0; i < nExpense; i++ {
			if expenseData[i].tanggal == tanggal && expenseData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Expense",
					nama:     expenseData[i].nama,
					jumlah:   expenseData[i].jumlah,
					tanggal:  expenseData[i].tanggal,
					bulan:    expenseData[i].bulan,
				})
				totalExpense += expenseData[i].jumlah
			}
		}
		for i := 0; i < nSub; i++ {
			if subData[i].tanggal == tanggal && subData[i].bulan == bulan {
				hasil = append(hasil, HasilTransaksi{
					kategori: "Subscription",
					nama:     subData[i].nama,
					jumlah:   subData[i].harga,
					tanggal:  subData[i].tanggal,
					bulan:    subData[i].bulan,
					aktif:    subData[i].aktif,
				})
				if subData[i].aktif {
					totalSub += subData[i].harga
				}
			}
		}

		fmt.Printf("\n=== FILTER TANGGAL %d BULAN %d ===", tanggal, bulan)
		insertionSortHasil(hasil)
		cetakHasil(hasil)

		fmt.Println("\n--- Ringkasan ---")
		fmt.Println("Total Income   :", totalIncome)
		fmt.Println("Total Expense  :", totalExpense)
		fmt.Println("Subscription   :", totalSub)
		fmt.Println("Net            :", totalIncome-(totalExpense+totalSub))
		return

	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	// Untuk sub==1 (keyword): urutkan lalu cetak
	insertionSortHasil(hasil)
	cetakHasil(hasil)
}

// ================= MAX, MIN, EXTREME =================
func getMaxIncome() (nama string, jumlah int) {
	if nIncome == 0 {
		return "-", 0
	}
	max := 0
	idx := 0
	for i := 0; i < nIncome; i++ {
		if incomeData[i].jumlah > max {
			max = incomeData[i].jumlah
			idx = i
		}
	}
	return incomeData[idx].nama, incomeData[idx].jumlah
}

func getMinIncome() (nama string, jumlah int) {
	if nIncome == 0 {
		return "-", 0
	}
	min := incomeData[0].jumlah
	idx := 0
	for i := 1; i < nIncome; i++ {
		if incomeData[i].jumlah < min {
			min = incomeData[i].jumlah
			idx = i
		}
	}
	return incomeData[idx].nama, incomeData[idx].jumlah
}

func getMaxExpense() (nama string, jumlah int) {
	if nExpense == 0 {
		return "-", 0
	}
	max := 0
	idx := 0
	for i := 0; i < nExpense; i++ {
		if expenseData[i].jumlah > max {
			max = expenseData[i].jumlah
			idx = i
		}
	}
	return expenseData[idx].nama, expenseData[idx].jumlah
}

func getMinExpense() (nama string, jumlah int) {
	if nExpense == 0 {
		return "-", 0
	}
	min := expenseData[0].jumlah
	idx := 0
	for i := 1; i < nExpense; i++ {
		if expenseData[i].jumlah < min {
			min = expenseData[i].jumlah
			idx = i
		}
	}
	return expenseData[idx].nama, expenseData[idx].jumlah
}

func cetakStatistik() {
	fmt.Println("===== STATISTIK TRANSAKSI =====")

	if nIncome == 0 {
		fmt.Println("Belum ada data income.")
	} else {
		namaMaxInc, maxInc := getMaxIncome()
		namaMinInc, minInc := getMinIncome()
		fmt.Printf("Income Tertinggi : %s | Rp%d\n", namaMaxInc, maxInc)
		fmt.Printf("Income Terendah  : %s | Rp%d\n", namaMinInc, minInc)
	}

	if nExpense == 0 {
		fmt.Println("Belum ada data expense.")
	} else {
		namaMaxExp, maxExp := getMaxExpense()
		namaMinExp, minExp := getMinExpense()
		fmt.Printf("Expense Tertinggi: %s | Rp%d\n", namaMaxExp, maxExp)
		fmt.Printf("Expense Terendah : %s | Rp%d\n", namaMinExp, minExp)
	}

	fmt.Println("-- Extreme (Selisih Income-Expense per Bulan) --")
	extremeFound := false
	for b := 1; b <= 12; b++ {
		totInc := 0
		totExp := 0
		for i := 0; i < nIncome; i++ {
			if incomeData[i].bulan == b {
				totInc += incomeData[i].jumlah
			}
		}
		for i := 0; i < nExpense; i++ {
			if expenseData[i].bulan == b {
				totExp += expenseData[i].jumlah
			}
		}
		if totInc > 0 || totExp > 0 {
			selisih := totInc - totExp
			fmt.Printf("  Bulan %2d: Income Rp%d | Expense Rp%d | Selisih Rp%d\n", b, totInc, totExp, selisih)
			extremeFound = true
		}
	}
	if !extremeFound {
		fmt.Println("  Belum ada data untuk dihitung.")
	}
}

// ================= MAIN MENU =================
func main() {
	var pilihan int

	fmt.Print("Masukkan saldo awal: ")
	fmt.Scan(&saldo)
	saldoAwal = saldo

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1.  Tambah Income")
		fmt.Println("2.  Tambah Expense")
		fmt.Println("3.  Tambah Subscription")
		fmt.Println("4.  Dashboard")
		fmt.Println("5.  Edit Income")
		fmt.Println("6.  Edit Expense")
		fmt.Println("7.  Edit Toggle Subscription")
		fmt.Println("8.  Riwayat Transaksi")
		fmt.Println("9.  Statistik ")
		fmt.Println("10. Keluar")
		fmt.Printf("Saldo Saat Ini: Rp%d\n", saldo)

		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahIncome()

		} else if pilihan == 2 {
			tambahExpense()

		} else if pilihan == 3 {
			tambahSubscription()

		} else if pilihan == 4 {
			hitungDashboard()

		} else if pilihan == 5 {
			var sub int
			fmt.Println("1. Edit Income")
			fmt.Println("2. Hapus Income")
			fmt.Print("Pilih: ")
			fmt.Scan(&sub)
			if sub == 1 {
				editIncome()
			} else if sub == 2 {
				hapusIncome()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}

		} else if pilihan == 6 {
			var sub int
			fmt.Println("1. Edit Expense")
			fmt.Println("2. Hapus Expense")
			fmt.Print("Pilih: ")
			fmt.Scan(&sub)
			if sub == 1 {
				editExpense()
			} else if sub == 2 {
				hapusExpense()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}

		} else if pilihan == 7 {
			var sub int
			fmt.Println("1. Edit Subscription")
			fmt.Println("2. Hapus Subscription")
			fmt.Println("3. Toggle Aktif/Nonaktif Subscription")
			fmt.Print("Pilih: ")
			fmt.Scan(&sub)
			if sub == 1 {
				editSubscription()
			} else if sub == 2 {
				hapusSubscription()
			} else if sub == 3 {
				toggleSubscription()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}

		} else if pilihan == 8 {
			cariDanFilterTransaksi()

		} else if pilihan == 9 {
			cetakStatistik()

		} else if pilihan == 10 {
			fmt.Println("Keluar...")
			return

		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}