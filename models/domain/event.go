package domain

import (
	"time"

	"gorm.io/gorm"
)

// ! Model
// ● Model atau Entity adalah struct representasi dari tabel di database
// ● Saat kita membuat tabel di database, direkomendasikan dibuatkan struct representasinya
// ● Hal ini agar kita tidak perlu melakukan pembuatan perintah SQL secara manual lagi

// ! Convention
// ● Saat membuat struct, secara default GORM akan melakukan mapping secara otomatis, dimana nama tabel akan dipilih dari nama Struct menggunakan lower_case jamak, sedangkan nama kolom akan dipilih menggunakan lower_case.
// ● Selain itu, secara otomatis GORM akan memilih field ID sebagai primary key
// ● Namun, sebenarnya disarankan dibanding dilakukan secara otomatis menggunakan GORM, lebih
// baik kita deklarasikan secara manual menggunakan tag
// ● https://gorm.io/docs/models.html#Fields-Tags

// ! GORM Soft Delete
// ● GORM Mendukung fitur Soft Delete secara otomatis, caranya kita cukup membuat field DeletedAt dengan time gorm.DeletedAt (alias untuk time.Time)
// ● Jika GORM mendeteksi terdapat field dengan nama DeletedAt, secara otomatis GORM akan melakukan Soft Delete
// ● Selain itu, ketika melakukan Query, secara otomatis juga GORM akan menambah filter Soft Delete, yang artinya hasil query hanya data yang belum di delete

type Event struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserID      int64          `gorm:"column:user_id;not null"`
	Name        string         `gorm:"column:name;not null"`
	Description string         `gorm:"column:description;not null"`
	Image       string         `gorm:"column:image;not null"`
	ImageId     string         `gorm:"column:image_id;not null"`
	Location    string         `gorm:"column:location;not null"`
	DateTime    time.Time      `gorm:"column:date_time;not null"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;not null"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;null"`
	// ! Belongs To di One to Many(milik ...)
	User User `gorm:"foreignKey:UserID;references:ID"`
}

// ! Table Name
// ● Secara default, nama tabel akan menggunakan lower_case dan jamak.
// ● Misal struct User akan menggunakan tabel users
// ● Misal struct OrderDetail akan menggunakan tabel order_details
// ● Namun jika kita ingin menggunakan manual nama tabel nya, kita bisa menggunakan interface
// Tabler, yang mewajibkan membuat method dengan nama TableName()
func (Event *Event) TableName() string {
	return "events" // ? ingat ini static func, tidak boleh dinamis, karena hanya akan dipanggil sekali saat aplikasi dijalankan, jadi tidak boleh menggunakan parameter atau variabel yang nilainya bisa berubah-ubah, karena akan menyebabkan ketidakpastian dalam penentuan nama tabel yang digunakan oleh GORM.
}
