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

type User struct {
	ID        int64          `gorm:"primaryKey;column:id;autoIncrement;<-:create"`                    // ? ID string `gorm:"primaryKey;column:id;<-:create"` digunakan untuk mendeklarasikan field ID sebagai primary key dalam model User, dengan tipe data string. Tag gorm:"primaryKey" menunjukkan bahwa field ini adalah primary key, sedangkan tag gorm:"column:id" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "id". Tag gorm:"<-:create" menunjukkan bahwa field ini hanya dapat diisi saat membuat record baru. Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan mengenali field ID sebagai primary key dan akan menggunakan nama kolom "id" saat melakukan operasi database.
	Name      string         `gorm:"column:name;not null"`                                            // ? Name string `gorm:"column:name;not null"` digunakan untuk mendeklarasikan field Name dalam model User, dengan tipe data string. Tag gorm:"column:name" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "name". Tag gorm:"not null" menunjukkan bahwa field ini tidak boleh bernilai null, sehingga setiap record yang dibuat harus memiliki nilai untuk field Name. Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan menggunakan nama kolom "name" saat melakukan operasi database yang melibatkan field Name.
	Email     string         `gorm:"column:email;unique;not null"`                                    // ? Email string `gorm:"column:email;unique;not null"` digunakan untuk mendeklarasikan field Email dalam model User, dengan tipe data string. Tag gorm:"column:email" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "email". Tag gorm:"unique" menunjukkan bahwa nilai field ini harus unik di seluruh tabel, sehingga tidak boleh ada dua record yang memiliki nilai email yang sama. Tag gorm:"not null" menunjukkan bahwa field ini tidak boleh bernilai null, sehingga setiap record yang dibuat harus memiliki nilai untuk field Email. Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan menggunakan nama kolom "email" saat melakukan operasi database yang melibatkan field Email.
	Password  string         `gorm:"column:password"`                                                 // ? Password string `gorm:"column:password"` digunakan untuk mendeklarasikan field Password dalam model User, dengan tipe data string. Tag gorm:"column:password" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "password". Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan menggunakan nama kolom "password" saat melakukan operasi database yang melibatkan field Password.
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`                      // ? CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"` digunakan untuk mendeklarasikan field CreatedAt dalam model User, dengan tipe data time.Time. Tag gorm:"column:created_at" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "created_at". Tag gorm:"autoCreateTime" menunjukkan bahwa GORM akan secara otomatis mengisi nilai CreatedAt dengan waktu saat record dibuat di database. Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan menangani pengisian nilai CreatedAt secara otomatis saat melakukan operasi database yang melibatkan pembuatan record baru.
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime;<-:update"` // ? UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;<-:update"` digunakan untuk mendeklarasikan field UpdatedAt dalam model User, dengan tipe data time.Time. Tag gorm:"column:updated_at" menunjukkan bahwa nama kolom di database yang sesuai dengan field ini adalah "updated_at". Tag gorm:"autoUpdateTime" menunjukkan bahwa GORM akan secara otomatis mengisi nilai UpdatedAt dengan waktu saat record diperbarui di database. Dengan menggunakan tag ini, kita dapat memastikan bahwa GORM akan menangani pengisian nilai UpdatedAt secara otomatis saat melakukan operasi database yang melibatkan pembaruan record.
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;null"`
	// ! one to many
	Events []Event `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"` // ? Event []Event `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"` digunakan untuk memberitahu gorm bahwa ada relasi one to many, dimana column user_id di table event foreign key dan reference id di table user. Artinya user memiliki banyak event, dan event memiliki satu user, ondelete cascade artinya ketika user dihapus maka eventnya juga akan dihapus, onupdate cascade artinya ketika user diupdate maka eventnya juga akan diupdate
}

// ! Table Name
// ● Secara default, nama tabel akan menggunakan lower_case dan jamak.
// ● Misal struct User akan menggunakan tabel users
// ● Misal struct OrderDetail akan menggunakan tabel order_details
// ● Namun jika kita ingin menggunakan manual nama tabel nya, kita bisa menggunakan interface
// Tabler, yang mewajibkan membuat method dengan nama TableName()
func (User *User) TableName() string {
	return "users" // ? ingat ini static func, tidak boleh dinamis, karena hanya akan dipanggil sekali saat aplikasi dijalankan, jadi tidak boleh menggunakan parameter atau variabel yang nilainya bisa berubah-ubah, karena akan menyebabkan ketidakpastian dalam penentuan nama tabel yang digunakan oleh GORM.
}
