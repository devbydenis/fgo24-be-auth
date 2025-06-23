package controller

import (
	m "auth/models"
	u "auth/utils"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	var req m.RegisterRequest
	err := ctx.ShouldBindJSON(&req)  // hanya untuk content-type json

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// validasi kalo input username, email, password kosong
	if req.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username tidak boleh kosong",
		})
		return
	}
	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email tidak boleh kosong",
		})
		return
	}
	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password tidak boleh kosong",
		})
		return
	}

	// Validasi panjang karakter username
	if len(req.Username) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username minimal 3 karakter",
		})
		return
	}

	// Validasi panjang karaker password
	if len(req.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password minimal 6 karakter",
		})
		return
	}

	// Cek email sudah ada atau belum
	if u.EmailExists(req.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email sudah terdaftar",
		})
		return
	}

	// Buat user baru
	newUser := m.User{
		ID:       m.CurrentID,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Dalam aplikasi nyata, password harus di-hash
	}

	// Simpan user ke slice users
	m.Users = append(m.Users, newUser)
	m.CurrentID++

	// Response sukses
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"user": gin.H{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}

func LoginHandler(ctx *gin.Context) {
	var req m.LoginRequest
	// ctx.ShouldBind(&req) // should bind bisa menerima content type aja secara dinamis
	err := ctx.ShouldBindJSON(&req); 
	
	fmt.Println("Email", req.Email)
	fmt.Println("Password", req.Password)
	
	// validasi kalo format yang kita masukin bukan json
	if  err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// validasi input kosong
	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email tidak boleh kosong",
		})
		return
	}
	
	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password tidak boleh kosong",
		})
		return
	}

	// cari user
	user := u.FindUserByEmail(req.Email)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email tidak ditemukan",
		})
		return
	}

	// cek password
	if user.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Password salah",
		})
		return
	}

	// response login berhasil
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func ForgotPasswordHandler(ctx *gin.Context) {
	var req m.ForgotPasswordRequest
	ctx.ShouldBindJSON(&req)
	
	if !u.EmailExists(req.Email) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Email tidak ditemukan, Pastikan email yang anda masukkan benar",
		})
		return
	}
	
	otp := u.GenerateOTP()
	m.UserOTP = otp
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP berhasil dikirim",
		"otp": otp,
	})
}

func ResetPasswordHandler(ctx *gin.Context) {
	var req m.ResetPasswordRequest
	ctx.ShouldBindJSON(&req)
	
	if !u.EmailExists(req.Email) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Email tidak ditemukan, Pastikan email yang anda masukkan benar",
		})
		return
	}
	
	if req.Otp == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "OTP tidak boleh kosong",
		})
		return
	}
	
	if req.Otp != strconv.Itoa(m.UserOTP) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "OTP tidak valid",
		})
		return
	}
	
	if req.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password baru tidak boleh kosong",
		})
		return
	}
	
	if req.NewPassword != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password baru dan konfirmasi password tidak cocok",
		})
		return
	}
	
	u.UpdateUserPassword(req.Email, req.NewPassword)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
	
	//testing
	fmt.Println("hasil new password: ", u.FindUserByEmail(req.Email))
}