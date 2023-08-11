package model

type Pembelis struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Nama_pembeli string `json:"nama_pembeli"`
	No_telp      string `json:"no_telp"`
	Alamat       string `json:"alamat"`
}
