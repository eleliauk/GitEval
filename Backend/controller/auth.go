package controller

import (
	"context"
	"github.com/GitEval/GitEval-Backend/api/request"
	"github.com/GitEval/GitEval-Backend/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthServiceProxy interface {
	Login(ctx context.Context) (url string, err error)
	CallBack(ctx context.Context, code string) (userId int64, err error)
}

type GenerateJWTer interface {
	GenerateToken(userId int64) (string, error)
	BlackJWT(ctx *gin.Context)
}

type AuthController struct {
	authService AuthServiceProxy
	jwt         GenerateJWTer
}

func NewAuthController(authService AuthServiceProxy, jwt GenerateJWTer) *AuthController {
	return &AuthController{
		authService: authService,
		jwt:         jwt,
	}
}

// Login 用户登录
// @Summary github用户登录授权接口
// @Description github用户登录授权接口,会自动重定向到github的授权接口上
// @Tags Auth
// @Produce json
// @Success 200 {object} response.Success "登录成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Failure 500 {object} response.Err "内部错误"
// @Router /api/v1/auth/login [get]
func (c *AuthController) Login(ctx *gin.Context) {
	url, err := c.authService.Login(ctx)
	if err != nil {
		// 处理错误，比如返回一个错误页面或重定向到错误页面
		ctx.JSON(http.StatusInternalServerError, response.Err{Err: err})
		return // 或根据需要返回其他值
	}

	// 重定向到 URL
	ctx.Redirect(http.StatusFound, url) // HTTP 302
	return
}

// CallBack 使用code进行最终登录
// @Summary 使用code进行最终登录
// @Description 使用code进行最终登录同时异步用来初始化这个用户,会返回一个token
// @Tags Auth
// @Param code query string true "github重定向的code"
// @Produce json
// @Success 200 {object} response.Success{data=response.CallBack} "初始化成功!"
// @Failure 400 {object} response.Err "请求参数错误"
// @Failure 500 {object} response.Err "内部错误"
// @Router /api/v1/auth/callBack [get]
func (c *AuthController) CallBack(ctx *gin.Context) {

	// 绑定查询参数
	var req request.CallBackReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{Err: err})
		return
	}

	userid, err := c.authService.CallBack(ctx, req.Code)
	if err != nil {
		return
	}

	token, err := c.jwt.GenerateToken(userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{Err: err})
		return
	}
	ctx.JSON(http.StatusOK, response.Success{
		Data: response.CallBack{
			Token: token,
		},
		Msg: "success",
	})
	return
}

// Logout 登出接口
// @Summary 登出
// @Description 登出之后会把jwt加到黑名单里面去
// @Tags Auth
// @Produce json
// @Success 200 {object} response.Success "登出成功!"
// @Failure 400 {object} response.Err "请求参数错误"
// @Failure 500 {object} response.Err "内部错误"
// @Router /api/v1/auth/logout [get]
func (c *AuthController) Logout(ctx *gin.Context) {
	c.jwt.BlackJWT(ctx)
	return
}
