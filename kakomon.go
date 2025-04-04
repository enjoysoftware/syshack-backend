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
		grade := ctx.Query("grade")
		subject := ctx.Query("subject")
		teacher := ctx.Query("teacher")

		if grade != "" {
			if subject != "" {
				if teacher != "" {
					// grade, subject, teacherが指定されている場合
					var kakomons []Kakomon
					result := db.Where("grade = ? AND subject = ? AND teacher = ?", grade, subject, teacher).Find(&kakomons)
					if result.Error != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}

					type TitleID struct {
						ID    uuid.UUID `json:"id"`
						Title string    `json:"title"`
					}

					var titleIDs []TitleID
					for _, kakomon := range kakomons {
						titleIDs = append(titleIDs, TitleID{ID: kakomon.ID, Title: kakomon.Title})
					}
					ctx.JSON(http.StatusOK, titleIDs)
					return

				} else {
					// grade, subjectが指定されている場合
					var teachers []string
					result := db.Model(&Kakomon{}).Where("grade = ? AND subject = ?", grade, subject).Distinct("teacher").Pluck("teacher", &teachers)
					if result.Error != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					ctx.JSON(http.StatusOK, teachers)
					return
				}
			} else {
				// gradeのみ指定されている場合
				var subjects []string
				result := db.Model(&Kakomon{}).Where("grade = ?", grade).Distinct("subject").Pluck("subject", &subjects)
				if result.Error != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
					return
				}
				ctx.JSON(http.StatusOK, subjects)
				return
			}
		} else {
			// パラメータが指定されていない場合、gradeのリストを出力
			var grades []string
			result := db.Model(&Kakomon{}).Distinct("grade").Pluck("grade", &grades)
			if result.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			ctx.JSON(http.StatusOK, grades)
			return
		}
	}
}

func GetKakomon(db *gorm.DB) gin.HandlerFunc { //過去問を一つ読み出す
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var kakomon Kakomon
		result := db.First(&kakomon, "id = ?", id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		mimeType := "application/octet-stream" // デフォルト値
		ext := filepath.Ext(kakomon.Path)
		switch ext {
		case ".pdf":
			mimeType = "application/pdf"
		case ".jpeg", ".jpg":
			mimeType = "image/jpeg"
		case ".png":
			mimeType = "image/png"
		}
		ctx.Header("Content-Type", mimeType)
		ctx.Header("Content-Disposition", "attachment; filename="+kakomon.Title+ext) //ファイル名を正しく
		ctx.File(kakomon.Path)
	}
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
		userID := kakomonInfo.UploadUserID
		var user User
		if err := db.Where("user_id = ?", userID).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		user.CountPost++
		// user.FeedingButterflyID
		db.Save(&user)

		// Create or Update Butterfly
		var butterfly Butterfly
		if err := db.Where("feed_user_id = ?", user.UserID).First(&butterfly).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new butterfly
				butterfly.FeedUserID = user.UserID
				db.Create(&butterfly)
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			// Update existing butterfly
			butterfly.GrowthStage++
			db.Save(&butterfly)
		}

		ctx.JSON(http.StatusOK, kakomonInfo)

	}
}

func DeleteKakomon(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID format"})
		}

		var kakomon Kakomon
		if err := db.Where("id = ?", id).First(&kakomon).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found!" + kakomon.Path})
			return
		}

		fileToDelete := kakomon.Path
		result := db.Unscoped().Where("id = ?", id).Delete(&Kakomon{}) //論理削除を防ぎ強制的に削除
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
			return
		}

		if os.Remove(fileToDelete) != nil {
			// ファイル削除に失敗した場合でも、データベースからは削除された状態なので、
			// エラーをログに出力するだけに留めることも検討する。
			// 例: log.Printf("Failed to delete file: %s, error: %v", fileToDelete, err)
			// 今回はエラーを返す
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to delete file: %s, error: %v", fileToDelete, err)})
			return
		}
		userID := kakomon.UploadUserID
		var user User
		if err := db.Where("user_id = ?", userID).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		if user.CountPost > 0 {
			user.CountPost--
			db.Save(&user)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
