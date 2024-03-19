package entity

type Skill struct {
	ID   uint
	Name string
	User []*User `json:"user" gorm:"many2many:user_skill;"`
}

/*
INSERT INTO `minats` (`name`) VALUES
('Desain Interior'),
('Edukasi'),
('Hiburan'),
('Hewan Peliharaan'),
('Industri Hijau'),
('Kecantikan'),
('Kesehatan'),
('Keuangan'),
('Layanan Kreatif'),
('Lingkungan'),
('Logistik'),
('Manajemen'),
('Manufaktur'),
('Makanan dan Minuman'),
('Media'),
('Mode/Fashion'),
('Olahraga'),
('Otomotif'),
('Pariwisata'),
('Pemasaran'),
('Penjualan Daring/e-commerce'),
('Perangkat Lunak'),
('Perjalanan'),
('Pertanian dan Perkebunan'),
('Peternakan'),
('Roti dan Kue Basah'),
('Riasan'),
('Riset'),
('Seni dan Kerajinan'),
('Tanaman'),
('Teknologi'),
('Teknologi Informasi'),
('Tekstil')
*/
