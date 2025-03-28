package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
			"uuid":    uuid.New(),
		})
	})
	r.GET("/users", func(ctx *gin.Context) {})            //ユーザ全件取得
	r.GET("/user/:user_id", func(ctx *gin.Context) {})    //ユーザ１件取得
	r.POST("/user", func(ctx *gin.Context) {})            //ユーザ登録(google_idをPOSTリクエストボディのJSONに含めて送信します)
	r.DELETE("/user/:user_id", func(ctx *gin.Context) {}) //user_idに指定されたユーザ削除

	r.PUT("/user/:user_id/administrator", func(ctx *gin.Context) {})

	r.GET("/kakomons", func(ctx *gin.Context) {})       //過去問一覧取得(取得する過去問の条件をGETパラメータで送信してください.複数の指定はできません)
	r.GET("/kakomon/:id", func(ctx *gin.Context) {})    //過去問指定取得(指定したidの過去問を取得します)
	r.POST("/kakomon", func(ctx *gin.Context) {})       //過去問登録(過去問情報はjsonで送信、ファイル本体はmultipart-formdataで送信します)
	r.DELETE("/kakomon/:id", func(ctx *gin.Context) {}) //指定した過去問を削除します

	r.GET("/butterflies/:feed_user_id", func(ctx *gin.Context) {}) //蝶取得一覧
	r.GET("/butterfly/:id", func(ctx *gin.Context) {})             //蝶指定取得
	r.POST("/butterfly/:id", func(ctx *gin.Context) {})            //蝶登録
	r.PUT("/butterfly/:id", func(ctx *gin.Context) {})             //蝶更新
}
