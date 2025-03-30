package main

//kakomonエンドポイントに関する実装を行ったファイル
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetKakomons(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetKakomon(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func CreateKakomon(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//過去問アップロード
		kakomonInfoStr := ctx.Request.FormValue("formData")
		if kakomonInfoStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "formData is required"})
			return
		}
		var kakomonInfo Kakomon
		if err := json.Unmarshal([]byte(kakomonInfoStr), &kakomonInfo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//年_専攻_中身名の形に書き換える
		kakomonInfo.Title = fmt.Sprintf("%d_%s_%s", kakomonInfo.Year, kakomonInfo.Major, kakomonInfo.Title)
		//UUIDを算出(ファイル名決定のためにGORMに任せない)
		kakomonInfo.ID = uuid.New()

		//ファイルアップロードに対応
		data, header, _ := ctx.Request.FormFile("file")
		ext := filepath.Ext(header.Filename) //拡張子取り出し
		kakomonInfo.Path = fmt.Sprintf("/kakomons/%s%s", kakomonInfo.ID, ext)

		dstFile, err := os.Create(kakomonInfo.Path)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer dstFile.Close()

		_, err = io.Copy(dstFile, data)

		// ファイル書き込みエラー時は整合性を失うかもしれないのでここでエラーハンドリングでストップ
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&kakomonInfo)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, kakomonInfo)

	}
}

func DeleteKakomon(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
