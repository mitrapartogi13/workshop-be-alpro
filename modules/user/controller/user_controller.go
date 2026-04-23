package controller

import (
	"net/http"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"

	// import baru
	"strconv"

	"github.com/Mobilizes/materi-be-alpro/modules/user/dto"
)

type UserController struct {
    service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
    return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    req, err := validation.ValidateCreateUser(c)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := ctrl.service.CreateUser(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}


func (ctrl *UserController) GetUserByID(c *gin.Context) {
    // ambil ID dari URL parameter
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
        return
    }

    // panggil service untuk cari user
    user, err := ctrl.service.GetUserByID(uint(id))
    if err != nil {
        // 3. Kembalikan 404 jika tidak ditemukan
        utils.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan")
        return
    }

    // format response menggunakan DTO (agar password tidak ikut terkirim)
    res := dto.UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        Role:  user.Role,
    }

    utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil data user", res)

    
}

// fungsi untuk menerima request HTTP, memanggil service, melakukan format data ke DTO, dan mengirim response JSON
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
    users, err := ctrl.service.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data user")
        return
    }

    // format response ke array DTO agar password tidak bocor
    var res []dto.UserResponse
    for _, user := range users {
        res = append(res, dto.UserResponse{
            ID:    user.ID,
            Name:  user.Name,
            Email: user.Email,
            Role:  user.Role,
        })
    }

    utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil semua data user", res)
}