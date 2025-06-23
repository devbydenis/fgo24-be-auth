package utils

import (
	m "auth/models"
	"math/rand"
	"strconv"
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
	// fmt.Println(m.Users)

	for i, user := range m.Users {
		if user.Email == email {
			return &m.Users[i]
		}
	}
	return nil
}

func GenerateOTP() int {
	result := 0
	for {
		randomNumber := rand.Intn(9999)
		if len(strconv.Itoa(randomNumber)) == 4 {
			result = randomNumber
			break
		}
	}
	return result
}

func UpdateUserPassword(requestEmail, newPassword string) {
	user := FindUserByEmail(requestEmail)
	if user != nil {
		user.Password = newPassword
	}
}
