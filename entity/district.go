package entity

type District struct {
	ID    uint `json:"id" gorm:"primary_key;autoIncrement"`
	Name  string
	Users []User
}

// 3507 kab malang
/*

INSERT INTO `districts` (`name`) VALUES
('DONOMULYO'),
('KALIPARE'),
('PAGAK'),
('BANTUR'),
('GEDANGAN'),
('SUMBERMANJING'),
('DAMPIT'),
('TIRTO YUDO'),
('AMPELGADING'),
('PONCOKUSUMO'),
('WAJAK'),
('TUREN'),
('BULULAWANG'),
('GONDANGLEGI'),
('PAGELARAN'),
('KEPANJEN'),
('SUMBER PUCUNG'),
('KROMENGAN'),
('NGAJUM'),
('WONOSARI'),
('WAGIR'),
('PAKISAJI'),
('TAJINAN'),
('TUMPANG'),
('PAKIS'),
('JABUNG'),
('LAWANG'),
('SINGOSARI'),
('KARANGPLOSO'),
('DAU'),
('PUJON'),
('NGANTANG'),
('KASEMBON'),
('KEDUNGKANDANG'),
('SUKUN'),
('KLOJEN'),
('BLIMBING'),
('LOWOKWARU'),

*/

// MINTA CREDENTIAL DATABASE / MINTA KAK VINCENT JALANIN QUERY //
// DATABASE DEPLOYMENT/PRODUCTION
