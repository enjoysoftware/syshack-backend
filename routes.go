package main

//apiエンドポイントを記載するファイル
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", Hello(db))
	r.GET("/users", GetUsers(db))
	r.GET("/user/:google_id", GetUser(db))
	r.POST("/user", CreateUser(db))
	r.DELETE("/user/:user_id", DeleteUser(db))
	r.PUT("/user/:user_id/administrator", UpdateUserAdministrator(db))

	r.GET("/kakomons", GetKakomons(db))
	r.GET("/kakomon/:id", GetKakomon(db))
	r.POST("/kakomon", CreateKakomon(db))
	r.DELETE("/kakomon/:id", DeleteKakomon(db))

	r.GET("/butterflies/:google_id", GetButterflies(db))
	r.GET("/butterfly/:id", GetButterfly(db))
	r.POST("/butterfly/:google_id", CreateButterfly(db)) //後でgoogle_idにする必要あるかも？
	// r.PUT("/butterfly/:id", UpdateButterfly(db))エンドポイントとして提供しない
}
