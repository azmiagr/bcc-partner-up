package bcrypt

import lib_bcrypt "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPasswordstring(password string) (string, error)
	CompareAndHashPassword(password string, hashPass string) error
}

type bcrypt struct {
	cost int
}

func Init() Interface {
	return &bcrypt{
		cost: 10,
	}
}

func (b *bcrypt) GenerateFromPasswordstring(password string) (string, error) {
	bytePass, err := lib_bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(bytePass), nil
}

func (b *bcrypt) CompareAndHashPassword(password string, hashPass string) error {
	err := lib_bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return nil
	}
	return nil
}
