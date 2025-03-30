package main

//userに関する処理を記述するファイル
import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var users []User

		db.Find(&users)

		ctx.JSON(http.StatusOK, gin.H{
			"user": users,
		})
	}
}

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		id := ctx.Param("user_id")
		res := db.Where("user_id = ?", id).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			fmt.Println("レコードが見つかりませんでした")
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc { //ユーザ登録
	return func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		fmt.Printf("Received JSON: %+v\n", user)
		db.Create(&user)

		ctx.JSON(http.StatusOK, user)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateUserAdministrator(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func Hello(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}
