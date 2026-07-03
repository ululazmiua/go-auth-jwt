package app

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ! Database Connection
// ● Hal pertama yang perlu kita lakukan sebelum menggunakan GORM, adalah membuat koneksi ke database
// ● Gunakan database yang sesuai dengan yang kita gunakan
// ● Untuk membuat koneksi ke database, kita bisa gunakan function gorm.Open()
func OpenConnection() *gorm.DB {
	// ! MySQL
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dialect := mysql.Open("root:gaktau4321@tcp(localhost:3306)/go_auth_jwt?charset=utf8mb4&parseTime=True&loc=Local") // ? mysql.Open(dsn string) digunakan untuk membuat koneksi ke database MySQL, dengan parameter berupa string yang berisi informasi koneksi seperti username, password, host, port, dan nama database.
	db, err := gorm.Open(dialect, &gorm.Config{                                                                               // ? gorm.Open(dialector gorm.Dialector, config *gorm.Config)(db *gorm.DB, err error) digunakan untuk membuka koneksi ke database dengan menggunakan dialector yang sesuai dengan jenis database yang digunakan, serta konfigurasi tambahan jika diperlukan.
		Logger: logger.Default.LogMode(logger.Info), // ? &gorm.Config{Logger: logger.Default.LogMode(logger.Info)} digunakan untuk mengatur konfigurasi GORM, khususnya untuk mengaktifkan logging dengan level Info. Dengan menggunakan Logger: logger.Default.LogMode(logger.Info), GORM akan mencatat semua query SQL yang dieksekusi beserta informasi tambahan seperti waktu eksekusi dan jumlah baris yang terpengaruh. Hal ini sangat berguna untuk debugging dan memantau performa aplikasi saat berinteraksi dengan database.
	})
	if err != nil {
		panic(err)
	}

	// 	! Connection Pool
	// ● Saat kita belajar di kelas Golang Database, kita belajar tentang Pool, dimana Golang mengatur maintain koneksi yang terbuka dan tertutup secara otomatis
	// ● Kita hanya cukup menggunakan saja, tanpa harus pusing mengaturnya
	// ● Bagaimana dengan GORM?
	// ● GORM sendiri sebenarnya didalamnya tetap menggunakan sql.DB
	// ● Jadi jika kita ingin mengubah pengaturan Pool nya, kita bisa menggunakan sql.DB

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)                 // ? sqlDB.SetMaxOpenConns(maxOpenConns int) digunakan untuk mengatur jumlah maksimum koneksi aktif ke database dalam satu waktu
	sqlDB.SetMaxIdleConns(10)                  // ? sqlDB.SetMaxIdleConns(maxIdleConns int) digunakan untuk mengatur jumlah maksimum koneksi idle (tidak digunakan) yang tetap disimpan
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // ? sqlDB.SetConnMaxLifetime(maxConnLifetime time.Duration) digunakan untuk mengatur batas waktu maksimal sebuah koneksi boleh digunakan sebelum ditutup dan dibuat ulang
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // ? sqlDB.SetConnMaxIdleTime(maxConnIdleTime time.Duration) digunakan untuk mengatur batas waktu koneksi boleh menganggur sebelum ditutup secara otomatis

	return db
}
