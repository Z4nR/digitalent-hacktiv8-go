package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Friend struct {
	Absent  int
	Name    string
	Address string
	Work    string
	Reason  string
}

func getFriendByAbsent(absent int) Friend {

	friendList := []Friend{
		{1, "Zulham Ari", "Semarang", "Backend Developer", "Ingin mempelajari bahasa Go"},
		{2, "Ali", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa Go"},
		{3, "Anggit", "Bandung", "Data Scientist", "Tertarik dengan kemampuan konkurensi di Go"},
		{4, "Alfit", "Surabaya", "Backend Developer", "Menginginkan performa tinggi dari aplikasinya"},
	}

	for _, friend := range friendList {
		if friend.Absent == absent {
			return friend
		}
	}

	return Friend{}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		os.Exit(1)
	}

	absent := os.Args[1]

	var absentInt int
	_, err := fmt.Sscanf(absent, "%d", &absentInt)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	friend := getFriendByAbsent(absentInt)

	// Periksa apakah teman ditemukan
	if friend.Name == "" {
		fmt.Println("Teman dengan absen tersebut tidak ditemukan.")
		os.Exit(1)
	}

	// Tampilkan data teman
	fmt.Println("Nama:", friend.Name)
	fmt.Println("Alamat:", friend.Address)
	fmt.Println("Pekerjaan:", friend.Work)
	fmt.Println("Alasan memilih kelas Golang:", friend.Reason)
}
