package handler

import (
	"Butterfly/token"
	"Butterfly/utils"
	"net/http"
	"time"

	db "Butterfly/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username        string `json:"username" binding:"required"`         // 用户名
	Password        string `json:"password" binding:"required"`         // 用户密码
	ConfirmPassword string `json:"confirm_password" binding:"required"` // 确认密码
	Email           string `json:"email" binding:"required,email"`      // 用户邮箱
	Gender          int8   `json:"gender" binding:"required"`           // 用户性别
}

func (h *Handler) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// h.ParamsError(ctx)
		return
	}

	// if !utils.ValidatePassword(req.Password) {
	// 	h.ParamsError(ctx, "密码格式不正确")
	// 	return
	// }

	// if !utils.ValidateUsername(req.Username) {
	// 	h.ParamsError(ctx, "用户名格式不正确")
	// 	return
	// }

	// 判断密码是否一致
	if req.Password != req.ConfirmPassword {
		// h.ParamsError(ctx, "密码不一致")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "密码不一致"})
		return
	}

	// 判断用户名是否存在
	if exists, _ := h.store.ExistsUsername(ctx, req.Username); exists {
		// h.ParamsError(ctx, "用户名已存在")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "用户名已存在"})
		return
	}

	// // 判断邮箱是否存在
	if exists, _ := h.store.ExistsEmail(ctx, req.Email); exists {
		// h.ParamsError(ctx, "邮箱已存在")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "邮箱已存在"})
		return
	}

	// 密码加密
	hashPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		// h.ParamsError(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 创建用户
	args := &db.CreateUserParams{
		Username: req.Username,
		Nickname: req.Username,
		Password: hashPwd,
		Email:    req.Email,
		Gender:   req.Gender,
	}

	_, err = h.store.CreateUser(ctx, args)
	if err != nil {
		// logs.Errorf("Create User Fail,Err:[%v]", err)
		// h.ServerError(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Create Success"})
	// h.Success(ctx, "Create Success")
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(ctx *gin.Context) {

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// h.ParamsError(ctx)
		return
	}

	// 判断用户是否存在
	user, err := h.store.GetUser(ctx, req.Username)
	if user.ID == 0 {
		// h.Error(ctx, http.StatusUnauthorized, "用户名不存在")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "用户名不存在"})
		return
	}
	// 校验密码
	err = utils.ComparePassword(req.Password, user.Password)
	if err != nil {
		// h.Error(ctx, http.StatusUnauthorized, "密码错误")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "密码错误"})
		return
	}
	// 生成Token
	accessToken, err := h.token.CreateToken(token.Token{
		Username: user.Username,
		Duration: time.Hour,
	})
	if err != nil {
		// return loginResponse{}, err
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// 邮箱脱敏
	// user.Email = utils.MaskEmail(user.Email)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{"data": accessToken})
}
