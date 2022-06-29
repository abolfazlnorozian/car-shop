package controllers

import (
	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminControllers struct {
	Adminservice services.Admins
}

func NewAdminService(adminservice services.Admins) AdminControllers {
	return AdminControllers{
		Adminservice: adminservice,
	}

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
	//var foundUser models.User

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this email not exist"})
		return

	}
	err := uc.Adminservice.LoginAdmin(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "this email not exist"})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *AdminControllers) AdminRoutes(rg *gin.RouterGroup) {
	adminroute := rg.Group("/admin")

	adminroute.POST("/register", uc.RegistrAdmin)
	adminroute.POST("/login", uc.LoginAdmin)
}
