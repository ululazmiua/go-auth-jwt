package helper

// ! package Helper
// package helper itu apa?
// ● Tempat fungsi bantuan (utility) aplikasi
// ● Dipakai ulang di banyak bagian program
// ● Bukan logika bisnis utama
// ● Biasanya berisi:
// ● helper database (commit/rollback tx)
// ● helper response JSON
// ● helper error handling
// ● helper konversi data

// * helper error handling
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
