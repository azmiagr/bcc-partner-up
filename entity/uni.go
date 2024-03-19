package entity

type Uni struct {
	ID    uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Name  string `json:"name" gorm:"type:varchar(255);not null;"`
	Users []User `json:"users" gorm:"foreignKey:UniID"`
}

/*
INSERT INTO `unis` (`name`) VALUES
('IKIP Budi Utomo Malang '),
('Institut Pertanian Malang'),
('Institut Teknologi dan Bisnis Asia Malang'),
('Institut Teknologi Nasional Malang'),
('Politeknik Kota Malang'),
('Politeknik Negeri Malang'),
('STIE Malangkucecwara'),
('STIKI Malang'),
('Universitas Bina Nusantara Malang'),
('Universitas Brawijaya'),
('Universitas Gajayana'),
('Universitas Islam Malang'),
('Universitas Islam Negeri Maulana Malik Ibrahim'),
('Universitas Katolik Widya Karya Malang'),
('Universitas Merdeka'),
('Universitas Muhammadiyah Malang'),
('Universitas Negeri Malang'),
('Universitas PGRI Kanjuruhan Malang'),
('Universitas Terbuka Malang'),
('Universitas Tribhuwana Tunggadewi'),
('Universitas Widyagama Malang'),
('Universitas Wisnuwardhana Malang')

*/
