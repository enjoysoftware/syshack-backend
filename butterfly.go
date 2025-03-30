package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetButterflies(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetButterfly(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func CreateButterfly(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateButterfly(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
