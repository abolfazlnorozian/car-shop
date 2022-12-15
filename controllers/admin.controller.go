package controllers

import (
	"fmt"
	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
	"io"
	"log"
	"os"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

type AdminControllers struct {
	Adminservice services.Admins
}

func NewAdminService(adminService services.Admins) AdminControllers {
	return AdminControllers{
		Adminservice: adminService,
	}

}

func (uc *AdminControllers) Uploadpath(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename

	out, err := os.Create("content/" + filename)
	if err != nil {
		log.Fatal(err)

	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)

	}

	filepath := "http://localhost:3000/v1/admin/download/" + filename
	ctx.JSON(http.StatusOK, gin.H{"filepath": filepath})

}
func (uc *AdminControllers) RegistrAdmin(ctx *gin.Context) {

	var admin models.Admin

	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	if err := ctx.ShouldBindJSON(&admin); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.Adminservice.CreateAdmin(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *AdminControllers) LoginAdmin(ctx *gin.Context) {

	var admin models.Admin

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this email not exist"})
		return

	}
	err := uc.Adminservice.LoginAdmin(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "this email not exist"})
		return

	}

	ctx.JSON(200, gin.H{"message": "success"})

}

func (uc *AdminControllers) AdminRoutes(rg *gin.RouterGroup) {
	adminroute := rg.Group("/admin")

	adminroute.POST("/register", uc.RegistrAdmin)
	adminroute.POST("/login", uc.LoginAdmin)
	// adminroute.GET("/serve", uc.DownloadImages)
	// adminroute.StaticFile("/download", "./content/pride.jpg")

	adminroute.Static("/download", "./content")
	// adminroute.POST("/upload", uc.Uploadpath)
	adminroute.POST("/upload", uc.Uploadpath)

}
