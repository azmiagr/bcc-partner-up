package entity

type Minat struct {
	ID   uint
	Name string
	User []*User `json:"user" gorm:"many2many:user_minat;"`
}

/*
INSERT INTO `skills` (`name`) VALUES
('Analisis Data'),
('Analisis Keuangan'),
('Bahasa Asing'),
('Desain Grafis'),
('Customer Service'),
('Desain Interior'),
('Desain Mode'),
('Desain Produk'),
('Elektronik Otomotif'),
('Keamanan Jaringan'),
('Keamanan Server'),
('Kepemimpinan'),
('Kolaborasi'),
('Komunikasi'),
('Kreativitas'),
('Layanan Pelanggan'),
('Manajemen Investaris'),
('Manajemen Proyek'),
('Manajemen Risiko'),
('Manajemen Waktu'),
('Melukis'),
('Membuat dan mendekorasi kue'),
('Menjahit'),
('Negoisasi'),
('Pemahaman UI/UX'),
('Pemasaran'),
('Pemecahan Masalah'),
('Pemograman'),
('Pengembangan Perangkat Lunak'),
('Pengembangan Produk'),
('Pengembangan Tim'),
('Pengetahuan Ilmiah dan Klinis'),
('Penulisan Naskah Konten'),
('Perencanaan Strategis'),
('Penyuntingan Foto'),
('Penyuntingan Video'),
('Riset')
*/
