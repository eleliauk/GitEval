package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/GitEval/GitEval-Backend/api/request"
	"github.com/GitEval/GitEval-Backend/api/response"
	"github.com/GitEval/GitEval-Backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceProxy interface {
	GetUserById(ctx context.Context, id int64) (model.User, error)
	GetLeaderboard(ctx context.Context, userId int64) ([]model.Leaderboard, error)
	GetDomains(ctx context.Context, userId int64) []string
	GetEvaluation(ctx context.Context, userId int64) (string, error)
	GetNationByUserId(ctx context.Context, userId int64) (string, error)
	GetDomainByUserId(ctx context.Context, userId int64) ([]string, error)
	SearchUser(ctx context.Context, nation *string, domain string, page int, pageSize int) ([]model.User, error)
}
type UserController struct {
	userService UserServiceProxy
}

func NewUserController(userService UserServiceProxy) *UserController {
	return &UserController{userService: userService}
}

// GetUser 获取用户
// @Summary 从userid获取用户
// @Tags User
// @Produce json
// @Success 200 {object} response.Success{data=response.User} "登录成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Router /api/v1/user/getInfo [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	UserID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("auth: %w", err),
		})
		return
	}

	user, err := c.userService.GetUserById(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{
			Err: fmt.Errorf("GetUserById: %w", err),
		})
		return
	}

	domain := c.userService.GetDomains(ctx, UserID)
	ctx.JSON(http.StatusOK, response.Success{
		Data: response.User{
			U:      user,
			Domain: domain,
		},
		Msg: "success",
	})
	return
}

// GetRanking 获取用户排行(和自己的好友之间的
// @Summary 根据userid获取用户的score的排行榜
// @Tags User
// @Produce json
// @Success 200 {object} response.Success{data=response.Ranking} "登录成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Router /api/v1/user/getRank [get]
func (c *UserController) GetRanking(ctx *gin.Context) {
	UserID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("auth: %w", err),
		})
		return
	}

	rankings, err := c.userService.GetLeaderboard(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{
			Err: fmt.Errorf("GetLeaderboard: %w", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{Data: response.Ranking{Leaderboard: rankings}, Msg: "success"})
	return
}

// GetEvaluation 获取用户评价
// @Summary 根据userid获取用户评价
// @Tags User
// @Produce json
// @Success 200 {object} response.Success{data=response.EvaluationResp} "登录成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Router /api/v1/user/getEvaluation [get]
func (c *UserController) GetEvaluation(ctx *gin.Context) {
	UserID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("auth: %w", err),
		})
		return
	}

	evaluation, err := c.userService.GetEvaluation(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{Err: fmt.Errorf("GetEvaluation: %w", err)})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{Data: response.EvaluationResp{Evaluation: evaluation}, Msg: "success"})
	return

}

// GetNation 获取用户所在国家
// @Summary 根据用户 ID 获取用户所在国家
// @Tags User
// @Produce json
// @Success 200 {object} response.Success{data=response.NationResp} "国家获取成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Failure 404 {object} response.Err "用户未找到"
// @Router /api/v1/user/getNation [get]
func (c *UserController) GetNation(ctx *gin.Context) {
	UserID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("auth: %w", err),
		})
		return
	}

	nation, err := c.userService.GetNationByUserId(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{Err: fmt.Errorf("GetNation: %w", err)})
		return
	}
	ctx.JSON(http.StatusOK, response.Success{Data: response.NationResp{Nation: nation}, Msg: "success"})

	return

}

// GetDomain 获取用户的领域
// @Summary 根据用户 ID 获取用户的领域
// @Tags User
// @Produce json
// @Success 200 {object} response.Success{data=response.DomainResp} "领域获取成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Failure 404 {object} response.Err "用户未找到"
// @Router /api/v1/user/getDomain [get]
func (c *UserController) GetDomain(ctx *gin.Context) {
	UserID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("auth: %w", err),
		})
		return
	}

	domain, err := c.userService.GetDomainByUserId(ctx, UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{Err: fmt.Errorf("GetDomain: %w", err)})
		return
	}
	ctx.JSON(http.StatusOK, response.Success{Data: response.DomainResp{Domain: domain}, Msg: "success"})
	return
}

// SearchUser 根据国家和领域搜索用户
// @Summary 根据国家和领域搜索用户
// @Tags User
// @Param nation query string false "国家，选择性参数"
// @Param domain query string true "领域，选择性参数"
// @Param page query int true "分页参数表示这是第几页"
// @Param page_size query int true "每页返回的用户数量，建议一次返回10个"
// @Produce json
// @Success 200 {object} response.Success{data=response.SearchResp} "用户搜索成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Router /api/v1/user/search [get]
func (c *UserController) SearchUser(ctx *gin.Context) {
	var req request.SearchUser
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: err,
		})
		return
	}

	users, err := c.userService.SearchUser(ctx, req.Nation, req.Domain, req.Page, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{
			Err: fmt.Errorf("FailSearch: %w", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{Data: response.SearchResp{Users: users}, Msg: "success"})
	return
}

// GetUserInfo 根据userid获取用户详细信息
// @Summary 根据userid获取用户详细信息
// @Tags User
// @Param user_id query string true "用户的user_id"
// @Produce json
// @Success 200 {object} response.Success{data=response.User} "用户信息获取成功"
// @Failure 400 {object} response.Err "请求参数错误"
// @Router /api/v1/user/getUserInfo [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	var req request.GetUserInfo

	// 解析请求体，绑定到结构体
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{
			Err: fmt.Errorf("invalid request: %w", err),
		})
		return
	}

	user, err := c.userService.GetUserById(ctx, req.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err{
			Err: fmt.Errorf("GetUserById: %w", err),
		})
		return
	}

	domain := c.userService.GetDomains(ctx, req.UserId)
	ctx.JSON(http.StatusOK, response.Success{
		Data: response.User{
			U:      user,
			Domain: domain,
		},
		Msg: "success",
	})
}

func getUserID(ctx *gin.Context) (int64, error) {
	userID, exist := ctx.Get("user_id")
	if !exist {
		return 0, errors.New("get user_id from ctx err")
	}
	UserID, ok := userID.(int64)
	if !ok {
		return 0, errors.New("transform interface{} to int64 err")
	}
	return UserID, nil
}
