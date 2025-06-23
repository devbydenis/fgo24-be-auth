package utils

import (
	m "auth/models"
	"fmt"
)

// Fungsi cek email sudah ada apa belom
func EmailExists(email string) bool {
	for _, user := range m.Users {
		if user.Email == email {
			return true
		}
	}
	return false
}

// Fungsi cari user berdasarkan email
func FindUserByEmail(email string) *m.User {
	fmt.Println(m.Users)

	for i, user := range m.Users {
		if user.Email == email {
			return &m.Users[i]
		}
	}
	return nil
}