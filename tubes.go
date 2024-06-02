package main

import (
	"fmt"
	"time"
)

const (
	NMAX      int    = 10000
	STARTDATE string = "2024-06-02"
	ENDDATE   string = "2024-06-08"
)

type data struct {
	Calon   InfoCalon
	Pemilih InfoPemilih
}

type InfoCalon struct {
	nama, partai  string
	noUrut, suara int
}
type InfoPemilih struct {
	nama, NIK, DoneVoting, status, CalonDipilih string
	vote                                        int
}

type infor [NMAX]data

func main() {
	var usn, pass string
	var choice1, nDataCalon, nDataPemilih int
	var Dta infor
	var waktuSekarang time.Time
	fmt.Println("Selamat Datang")
	//contoh data pemilih
	Dta[0].Pemilih.nama = "Budi"
	Dta[0].Pemilih.status = "petugas"
	Dta[0].Pemilih.NIK = "12345678"
	Dta[1].Pemilih.nama = "Gojo"
	Dta[1].Pemilih.status = "pemilih"
	Dta[1].Pemilih.NIK = "87654321"
	Dta[2].Pemilih.nama = "Naruto"
	Dta[2].Pemilih.status = "petugas"
	Dta[2].Pemilih.NIK = "01234567"
	Dta[3].Pemilih.nama = "Siti"
	Dta[3].Pemilih.status = "pemilih"
	Dta[3].Pemilih.NIK = "23456789"
	nDataPemilih = 4
	//contoh data calon
	Dta[0].Calon.nama = "Sai"
	Dta[0].Calon.noUrut = 1
	Dta[0].Calon.partai = "PDIP"
	Dta[0].Calon.suara = 1
	Dta[1].Calon.nama = "Didi"
	Dta[1].Calon.noUrut = 1
	Dta[1].Calon.partai = "PKS"
	Dta[2].Calon.nama = "Dudung"
	Dta[2].Calon.noUrut = 2
	Dta[2].Calon.partai = "PKS"
	nDataCalon = 3
	waktuSekarang = time.Now() // Dapatkan waktu saat ini
	// Format waktu sebagai "YYYY-MM-DD"
	if waktuSekarang.Format("2006-01-02") <= ENDDATE {
		//Tampilan awal
		tampilanAwal(&Dta, choice1, &nDataPemilih, &nDataCalon, &usn, &pass)
	} else {
		fmt.Println("Periode pemilihan telah lewat")
		TampilanSetelahPeriode(&Dta, nDataPemilih, nDataCalon, usn, pass)
	}
}
func tampilanAwal(A *infor, choice int, nDP, nDC *int, usn, pass *string) {
	//menampilkan pilihan pada pengguna untuk login
	fmt.Println("Masuk sebagai:")
	fmt.Println("1. Pemilih")
	fmt.Println("2. Petugas")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	verificationChoice(A, choice, nDP, nDC, *usn, *pass)
}
func verificationChoice(A *infor, choice int, nP, nC *int, usn, pass string) {
	//mengecek apakah pengguna merupakan pemilih atau petugas
	var cek bool
	if choice == 1 {
		//Tampilan berikut akan muncul apabila pengguna adalah pemilih
		pemilih(A, nP, choice, *nC, &usn, &pass)
	} else if choice == 2 {
		//Tampilan berikut akan muncul apabila pengguna adalah petugas
		petugas(A, nC, nP, usn, pass, &cek)
	} else if choice == 3 {
		fmt.Println("Terima kasih!")
	} else {
		fmt.Print("Maaf pilihan anda tidak tertera pada menu")
		tampilanAwal(A, choice, nP, nC, &usn, &pass)
	}
}

func petugas(A *infor, nC, nV *int, usn, pass string, cek *bool) {
	//memverifikasi apakah pengguna merupakan petugas
	var i int
	fmt.Print("Nama: ")
	fmt.Scan(&usn)
	fmt.Print("NIK: ")
	fmt.Scan(&pass)
	*cek = false
	for i < *nV && !*cek {
		*cek = A[i].Pemilih.status == "petugas"
		i += 1
	}
	//Apabila pengguna ter verifikasi sebagai petugas, pengguna dapat mengakses pilihan berikut
	if *cek == true {
		pilihanPetugas(A, nC, nV, usn, pass)
	}
}

func pemilih(A *infor, nV *int, choice, nC int, usn, pass *string) {
	//menampilkan menu untuk pemilih
	fmt.Println("MENU:")
	fmt.Println("1. Daftarkan diri sebagai pemilih")
	fmt.Println("2. Pilih calon")
	fmt.Println("3. Cari calon")
	fmt.Println("4. Lihat data calon")
	fmt.Println("5. Kembali")
	fmt.Println("6. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	fiturPemilih(choice, A, nV, nC, usn, pass)
}

func pilihanPetugas(A *infor, nC, nV *int, usn, pass string) {
	//menampilkan menu untuk petugas
	var choice int
	fmt.Println("MENU:")
	fmt.Println("1. Tambahkan calon")
	fmt.Println("2. Edit calon")
	fmt.Println("3. Hapus calon")
	fmt.Println("4. Lihat data calon")
	fmt.Println("5. Cari calon")
	fmt.Println("6. Hapus pemilih")
	fmt.Println("7. Edit pemilih")
	fmt.Println("8. Kembali")
	fmt.Println("9. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	fiturPetugas(choice, A, nC, nV, &usn, &pass)
}

func isVotingPeriod() bool {
	//mengecek apakah sedang di dalam periode pemilihan
	var currentDate string
	currentDate = time.Now().Format("2006-01-02")
	return currentDate >= STARTDATE && currentDate <= ENDDATE
}
func TampilanSetelahPeriode(A *infor, nP, nC int, usn, pass string) {
	var choice int
	//tampilan berikut muncul jika masa pemilihan telah lewat
	fmt.Println("MENU:")
	fmt.Println("1. Lihat daftar calon yang terpilih")
	fmt.Println("2. Lihat hasil pemilihan suara")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	fiturDiluarPeriode(A, nP, nC, choice, usn, pass)
}
func fiturDiluarPeriode(A *infor, nP, nC, choice int, usn, pass string) {
	//menerima hasil pilihan pengguna dari tampilansetelahperiode
	if choice == 1 {
		calonTerpilih(A, nC, nP, usn, pass)
	} else if choice == 2 {
		DataCalon(A, nP, nC, usn, pass, false)
	} else if choice == 3 {
		fmt.Println("Terima kasih telah menggunakan aplikasi pemilihan umum")
	} else {
		fmt.Println("Maaf pilihan anda tidak tertera pada menu")
		TampilanSetelahPeriode(A, nP, nC, usn, pass)
	}
}
func fiturPetugas(CP int, A *infor, nC, nP *int, usn, pass *string) {
	//IS. terdefinisi var CP yang berisi nomor urut fitur yang di pilih pengguna, var A yang berisi data array, serta var nC dan nP yang berisi total calon dan pemilih yang terdaftar
	//serta usn dan pass dengan var bertipe string
	//FS. akan dijalankan fungsi berdasarkan pilihan pengguna
	var petugas bool
	petugas = true
	if CP == 1 {
		InputCalon(A, nC, nP, *usn, *pass)
	} else if CP == 2 {
		EditCalon(A, *nP, *nC, *usn, *pass)
	} else if CP == 3 {
		HapusCalon(A, nC, *nP, *usn, *pass)
	} else if CP == 4 {
		DataCalon(A, *nP, *nC, *usn, *pass, petugas)
	} else if CP == 5 {
		CariCalon(A, *nC, *nP, *usn, *pass, petugas)
	} else if CP == 6 {
		HapusPemilih(A, nC, nP, *usn, *pass)
	} else if CP == 7 {
		EditPemilih(A, nP, *nC, usn, pass)
	} else if CP == 8 {
		tampilanAwal(A, CP, nP, nC, usn, pass)
	} else if CP == 9 {
		fmt.Println("Terima kasih telah menggunakan aplikasi pemilihan umum")
	} else {
		fmt.Println("Maaf pilihan anda tidak tertera pada menu")
		pilihanPetugas(A, nC, nP, *usn, *pass)
	}
}

func fiturPemilih(CP int, A *infor, nP *int, nC int, usn, pass *string) {
	//IS. terdefinisi var CP yang berisi nomor urut fitur yang di pilih pengguna, var A yang berisi data array, serta var nC dan nP yang berisi total calon dan pemilih yang terdaftar
	//serta usn dan pass dengan var bertipe string
	//FS. akan dijalankan fungsi berdasarkan pilihan pengguna
	var petugas bool
	var choice int
	petugas = false
	if CP == 1 {
		InputPemilih(A, &nC)
	} else if CP == 2 {
		lakukanVoting(A, *nP, nC, *usn, *pass)
	} else if CP == 3 {
		CariCalon(A, nC, *nP, *usn, *pass, petugas)
	} else if CP == 4 {
		DataCalon(A, *nP, nC, *usn, *pass, petugas)
	} else if CP == 5 {
		tampilanAwal(A, CP, nP, &nC, usn, pass)
	} else if CP == 6 {
		fmt.Println("Terima kasih telah menggunakan aplikasi pemilihan umum")
	} else {
		fmt.Println("Maaf pilihan anda tidak tertera pada menu")
		pemilih(A, nP, choice, nC, usn, pass)
	}
}

func DataCalon(A *infor, nP, nC int, usn, pass string, petugas bool) {
	//menampilkan menu kepada pengguna dan menampilkan data calon berdasarkan pilihan pengguna
	var choice, choice1, choice2, i int
	var waktu bool
	waktu = isVotingPeriod()
	fmt.Println("Tampilkan berdasarkan: ")
	fmt.Println("1. Hasil perolehan suara terbesar")
	fmt.Println("2. Hasil perolehan suara terkecil")
	fmt.Println("3. Partai")
	fmt.Println("4. Nama calon dan partai")
	fmt.Println("5. Kembali")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	if choice == 1 {
		selectionSortCalonBySuara(A, nC, false)
	} else if choice == 2 {
		selectionSortCalonBySuara(A, nC, true)
	} else if choice == 3 {
		SortCalonByPartai(A, nC)
	} else if choice == 4 {
		insertionSortCalonByNamaDanPartai(A, nC)
	} else if choice == 5 && petugas {
		pilihanPetugas(A, &nC, &nP, usn, pass)
	} else if choice == 5 && !petugas && waktu {
		pemilih(A, &nP, choice1, nC, &usn, &pass)
	} else if choice == 5 && !petugas && !waktu {
		TampilanSetelahPeriode(A, nP, nC, usn, pass)
	} else {
		fmt.Println("Pilihan tidak valid")
		DataCalon(A, nP, nC, usn, pass, petugas)
	}
	for i < nC && choice != 3 {
		fmt.Println("Nama:", A[i].Calon.nama)
		fmt.Println("Nomor Urut:", A[i].Calon.noUrut)
		fmt.Println("Partai:", A[i].Calon.partai)
		fmt.Println("Perolehan Suara:", A[i].Calon.suara)
		fmt.Println()
		i++
	}
	pemilih(A, &nP, choice2, nC, &usn, &pass)
}

func calonTerpilih(A *infor, nC, nP int, usn, pass string) {
	// fungsi ini akan menampilkan calon yang terpilih
	// calon terpilih apabila berhasil memenuhi ambang batas yaitu 4% dari total suara pemilih
	var i int
	var ambangBatas float64
	ambangBatas = 4 * float64(nP) / 100
	for i < nC {
		if float64(A[i].Calon.suara) >= ambangBatas {
			fmt.Println("Nama:", A[i].Calon.nama)
			fmt.Println("Nomor Urut:", A[i].Calon.noUrut)
			fmt.Println("Partai:", A[i].Calon.partai)
			fmt.Println("Perolehan suara:", A[i].Calon.suara)
			fmt.Println()
		}
		i++
	}
	TampilanSetelahPeriode(A, nP, nC, usn, pass)
}
func CariCalon(A *infor, nC, nP int, usn, pass string, petugas bool) {
	//mencari calon berdasarkan pilihan pengguna
	var i, j, pilihan, choice2 int
	var choice, name string
	fmt.Println("Cari Calon Berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Pemilih dari calon")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		var nama string
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&nama)
		i = sequentialSearchCalonByNama(*A, nC, nama)
	} else if pilihan == 2 {
		var pemilih string
		fmt.Print("Masukkan nama pemilih dari calon: ")
		fmt.Scan(&pemilih)
		j = sequentialSearchCalonByPemilih(*A, nP, pemilih)
		name = A[j].Pemilih.CalonDipilih
		i = sequentialSearchCalonByNama(*A, nC, name)
	} else {
		fmt.Println("Pilihan tidak valid")
		CariCalon(A, nC, nP, usn, pass, petugas)
	}
	fmt.Println("Berikut ialah data Calon yang anda cari:")
	fmt.Println("Nama:", A[i].Calon.nama)
	fmt.Println("Nomor Urut:", A[i].Calon.noUrut)
	fmt.Println("Partai:", A[i].Calon.partai)
	fmt.Println("Perolehan Suara:", A[i].Calon.suara)

	fmt.Println("Apakah anda ingin mencari data calon lain? Y/N?")
	fmt.Scan(&choice)
	if choice == "Y" {
		CariCalon(A, nC, nP, usn, pass, petugas)
	} else if choice == "N" && petugas {
		pilihanPetugas(A, &nC, &nP, usn, pass)
	} else if choice == "N" && !petugas {
		pemilih(A, &nP, choice2, nC, &usn, &pass)
	} else {
		fmt.Println("pilihan tidak valid")
	}
}
func InputCalon(A *infor, nC, nP *int, usn, pass string) {
	//menginput data calon baru
	var i, choice int
	var ans string
	i = *nC
	*nC += 1
	fmt.Println("Mohon gunakan tanda spasi bawah(_) jika nama lebih dari satu kata")

	fmt.Print("Nama calon: ")
	fmt.Scan(&A[i].Calon.nama)

	fmt.Println("Pilihan Partai")
	fmt.Println("1. PDIP")
	fmt.Println("2. PKS")
	fmt.Println("3. Partai Buruh")
	fmt.Print("Partai 1/2/3?")
	fmt.Scan(&choice)
	if choice == 1 {
		A[i].Calon.partai = "PDIP"
	} else if choice == 2 {
		A[i].Calon.partai = "PKS"
	} else if choice == 3 {
		A[i].Calon.partai = "Partai Buruh"
	} else {
		fmt.Println("Pilihan anda tidak valid")
	}

	fmt.Print("Nomor Urut: ")
	fmt.Scan(&A[i].Calon.noUrut)

	fmt.Println("Calon berhasil di daftarkan")
	fmt.Println("Apakah anda masih ingin menambahkan calon? Y/N? ")
	fmt.Scan(&ans)
	if ans == "Y" {
		InputCalon(A, nC, nP, usn, pass)
	} else if ans == "N" {
		pilihanPetugas(A, nC, nP, usn, pass)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}
func EditCalon(A *infor, nP, nC int, usn, pass string) {
	//mengedit data calon
	var ans, ans3, i, choice int
	var ans2, ans4 string
	fmt.Print("Partai Calon yang di edit: ")
	fmt.Scan(&ans2)
	fmt.Print("No urut Calon yang di edit: ")
	fmt.Scan(&ans3)
	fmt.Println("Edit data: ")
	fmt.Println("1. Nomor urut")
	fmt.Println("2. Nama")
	fmt.Println("3. Partai")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&ans)

	i = sequentialSearchCalonByNoUrutDanPartai(*A, nC, ans3, ans2)

	if ans == 1 {
		fmt.Print("Nomor urut baru: ")
		fmt.Scan(&A[i].Calon.noUrut)
	} else if ans == 2 {
		fmt.Print("Nama baru: ")
		fmt.Scan(&A[i].Calon.nama)
	} else if ans == 3 {
		fmt.Println("Pilihan Partai")
		fmt.Println("1. PDIP")
		fmt.Println("2. PKS")
		fmt.Println("3. Partai Buruh")
		fmt.Print("Partai 1/2/3?")
		fmt.Scan(&choice)
		if choice == 1 {
			A[i].Calon.partai = "PDIP"
		} else if choice == 2 {
			A[i].Calon.partai = "PKS"
		} else if choice == 3 {
			A[i].Calon.partai = "Partai Buruh"
		} else {
			fmt.Println("Pilihan anda tidak valid")
		}
	} else {
		fmt.Println("Maaf pilihan anda tidak tertera pada menu")
		EditCalon(A, nP, nC, usn, pass)
	}
	fmt.Println("Data telah berhasil di edit.")
	fmt.Print("Apakah anda masih ingin mengedit data calon lainnya? Y/N?")
	fmt.Scan(&ans4)
	if ans4 == "Y" {
		EditCalon(A, nP, nC, usn, pass)
	} else if ans4 == "N" {
		pilihanPetugas(A, &nC, &nP, usn, pass)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}
func HapusCalon(A *infor, nC *int, nP int, usn, pass string) {
	//menghapus data calon dari array
	var ans2, ans3 string
	var i, j, nJ, ans int
	fmt.Print("Nomor urut calon: ")
	fmt.Scan(&ans)
	fmt.Print("Partai calon: ")
	fmt.Scan(&ans3)
	nJ = *nC - 1
	for i < *nC {
		if A[i].Calon.noUrut == ans && A[i].Calon.partai == ans3 {
			j = i
			for j < nJ {
				A[j] = A[j+1]
				j += 1
			}
		}
		i += 1
	}
	*nC = nJ
	fmt.Println("Data calon telah di hapus")
	fmt.Print("Apakah anda masih ingin menghapus data calon lainnya? Y/N? ")
	fmt.Scan(&ans2)
	if ans2 == "Y" {
		HapusCalon(A, nC, nP, usn, pass)
	} else if ans2 == "N" {
		pilihanPetugas(A, nC, &nP, usn, pass)
	}
}
func InputPemilih(A *infor, n *int) {
	//IS. terdefenisi array A dan n bilangan bulat
	//FS. menambahkan data baru pemilih ke dalam array
	var i int
	i = *n
	*n += 1
	for i < *n {
		fmt.Print("Nama: ")
		fmt.Scan(&A[i].Pemilih.nama)
		fmt.Print("NIK: ")
		fmt.Scan(&A[i].Pemilih.NIK)
		A[i].Pemilih.DoneVoting = "Belum"
		A[i].Pemilih.status = "pemilih"
		i++
	}
	fmt.Println("Data anda berhasil di daftarkan")
}
func EditPemilih(A *infor, nP *int, nC int, usn, pass *string) {
	//mengedit data pemilih
	var ans, i, choice int
	var nik string
	var ketemu bool
	ketemu = false
	fmt.Print("Masukkan NIK pemilih: ")
	fmt.Scan(&nik)
	i = 0
	for i < *nP && !ketemu {
		if A[i].Pemilih.NIK == nik {
			fmt.Println("Edit Data Pemilih")
			fmt.Println("1. Nama")
			fmt.Println("2. NIK")
			fmt.Println("3. Nama dan NIK")
			fmt.Println("Masukkan pilihan: ")
			fmt.Scan(&ans)

			if ans == 1 {
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&A[i].Pemilih.nama)
				ketemu = true
				fmt.Println("Data pemilih berhasil diubah.")
			} else if ans == 2 {
				fmt.Print("Masukkan NIK baru: ")
				fmt.Scan(&A[i].Pemilih.NIK)
				ketemu = true
				fmt.Println("Data pemilih berhasil diubah.")
			} else if ans == 3 {
				fmt.Print("Masukkan nama baru: ")
				fmt.Scan(&A[i].Pemilih.nama)
				fmt.Print("Masukkan NIK baru: ")
				fmt.Scan(&A[i].Pemilih.NIK)
				ketemu = true
				fmt.Println("Data pemilih berhasil diubah.")
			}
		}
		i++
	}
	if i == *nP {
		fmt.Println("NIK tidak ditemukan.")
	}
	pemilih(A, nP, choice, nC, usn, pass)
}
func HapusPemilih(A *infor, nC, nV *int, usn, pass string) {
	//IS terdefinisi var A yang menampung nilai array, nC, nV yang menampung jumlah calon dan pemilih yang terdaftar
	//serta variabel usn dan pass bertipe string
	//FS menghapus data pemilih sesuai NIK dari pemilih
	var ans, ans2 string
	var i, j, nJ int
	fmt.Print("NIK: ")
	fmt.Scan(&ans)
	nJ = *nV - 1
	for i < *nV {
		if A[i].Pemilih.NIK == ans {
			j = i
			for j < nJ {
				A[j] = A[j+1]
				j += 1
			}
		}
		i += 1
	}
	*nV = nJ
	fmt.Println("Data pemilih telah di hapus")
	fmt.Print("Apakah anda masih ingin menghapus data pemilih lainnya? Y/N? ")
	fmt.Scan(&ans2)
	if ans2 == "Y" {
		HapusPemilih(A, nC, nV, usn, pass)
	} else if ans2 == "N" {
		pilihanPetugas(A, nC, nV, usn, pass)
	} else {
		fmt.Println("Pilihan tidak valid")
		HapusPemilih(A, nC, nV, usn, pass)
	}
}
func lakukanVoting(A *infor, nP, nC int, usn, pass string) {
	//IS terdefenisi var A yang berisi array dan var n yang berisi jumlah pemilih yang terdaftar
	//FS memasukkan data calon yang dipilih oleh pemilih
	var ans3, i, j, choice int
	var ans, ans2 string
	fmt.Print("Harap masukkan NIK anda: ")
	fmt.Scan(&ans)
	for i < nP {
		if A[i].Pemilih.NIK == ans {
			if A[i].Pemilih.DoneVoting == "Sudah" {
				fmt.Println("Maaf anda sudah pernah melakukan pemilihan")
			} else {
				fmt.Print("Masukkan nama partai calon yang akan dipilih: ")
				fmt.Scan(&ans2)
				fmt.Print("Masukkan nomor urut calon yang akan dipilih: ")
				fmt.Scan(&ans3)
				A[i].Pemilih.vote = ans3
				A[i].Pemilih.DoneVoting = "Sudah"
				j = sequentialSearchCalonByNoUrutDanPartai(*A, nC, ans3, ans2)
				A[j].Calon.suara += 1
				A[i].Pemilih.CalonDipilih = A[j].Calon.nama
			}
		}
		i++
	}
	pemilih(A, &nP, choice, nC, &usn, &pass)
}

func sequentialSearchCalonByNama(A infor, n int, nama string) int {
	//mengembalikan ideks calon berdasarkan nama
	var i int
	for i < n {
		if A[i].Calon.nama == nama {
			return i
		}
		i++
	}
	return -1
}

func sequentialSearchCalonByNoUrutDanPartai(A infor, n int, noUrut int, partai string) int {
	//mengembalikan indeks calon berdasarkan no urut dan partai
	var i, idx int
	var ketemu bool
	for i < n && !ketemu {
		ketemu = A[i].Calon.noUrut == noUrut && A[i].Calon.partai == partai
		idx = i
		i++
	}
	return idx
}

func sequentialSearchCalonByPemilih(A infor, n int, pemilih string) int {
	//mengembalikan indeks pemilih yang memilih calon
	var i int
	for i < n {
		if A[i].Pemilih.nama == pemilih {
			return i
		}
		i++
	}
	return -1
}

func sequentialSearchPemilih(A infor, n int, NIK string) int {
	//mengembalikan indeks pemilih bedasarkan nik
	for i := 0; i < n; i++ {
		if A[i].Pemilih.NIK == NIK {
			return i
		}
	}
	return -1
}

func selectionSortCalonBySuara(A *infor, n int, ascending bool) {
	//menyusun data calon dari total suara yang dipunya
	var i, j, idx int
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if (ascending && A[j].Calon.suara < A[idx].Calon.suara) || (!ascending && A[j].Calon.suara > A[idx].Calon.suara) {
				idx = j
			}
		}
		// Swap
		A[i], A[idx] = A[idx], A[i]
	}
}

func SortCalonByPartai(A *infor, n int) {
	//menyusun data calon berdasarkan partai
	var i int
	fmt.Println("PDIP")
	for i < n {
		if A[i].Calon.partai == "PDIP" {
			fmt.Println(A[i].Calon.noUrut, A[i].Calon.nama)
		}
		i++
	}
	i = 0
	fmt.Println("PKS")
	for i < n {
		if A[i].Calon.partai == "PKS" {
			fmt.Println(A[i].Calon.noUrut, A[i].Calon.nama)
		}
		i++
	}
	i = 0
	fmt.Println("Partai_Buruh")
	for i < n {
		if A[i].Calon.partai == "Partai_Buruh" {
			fmt.Println(A[i].Calon.noUrut, A[i].Calon.nama)
		}
		i++
	}
}

func insertionSortCalonByNamaDanPartai(A *infor, n int) {
	//menyusun data calon berdasarkan nama dan partai
	var i, j int
	var key data
	for i < n {
		key = A[i]
		j = i - 1
		for j >= 0 && A[j].Calon.partai > key.Calon.partai && A[j].Calon.nama > key.Calon.nama {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
		i++
	}
}
