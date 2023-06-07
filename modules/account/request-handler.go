package account

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/exception"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"github.com/imama2/bootcamp-bri-mini-project/package/middleware"
	"net/http"
	"strconv"
)

type RequestHandlerAccount struct {
	AccountUseCase UseCaseAccountInterface
}

func NewAccountRequestHandler(uc UseCaseAccountInterface) *RequestHandlerAccount {
	return &RequestHandlerAccount{
		AccountUseCase: uc,
	}
}
func (h *RequestHandlerAccount) RouteAccount(app *gin.Engine) {
	// account
	a := NewAccountRequestHandler(h.AccountUseCase)
	a.RouteHandler(app)
}

func (h *RequestHandlerAccount) RouteHandler(app *gin.Engine) {
	app.POST("/login", h.SignIn) // Generate token JWT

	g := app.Group("/account", middleware.Auth()) // using middleware
	g.POST("", h.RegisterAccount)
	// !mplement goroutine, sometimes errors bad connection, and a busy buffer,
	// !Error tx with go routine, solution using simple db queries
	g.GET("", h.GetAllAdmin)

	// only super_admin
	g.GET("/admin-menu", middleware.AuthSuperAdmin(), h.GetAllAppovalAdmin)
	g.PUT("/admin-menu", middleware.AuthSuperAdmin(), h.UpdateAdminStatus)
	g.DELETE("/admin-menu", middleware.AuthSuperAdmin(), h.DeleteAdminByID)
}

func (h *RequestHandlerAccount) SignIn(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	do := do.Account{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.AccountUseCase.AccountAuthentication(do)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerAccount) RegisterAccount(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	do := do.Account{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.AccountUseCase.AccountRegistration(do)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := RowsAffected{
		Message:      "Success",
		RowsAffected: result,
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerAccount) GetAllAppovalAdmin(c *gin.Context) {

	result, err := h.AccountUseCase.GetAllApprovalAdmin()
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: "Success",
		Data:    ResponseListAdminReg(result),
	}

	c.JSON(http.StatusOK, res)
}
func (h *RequestHandlerAccount) GetAllAdmin(c *gin.Context) {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}
	username := c.Query("username")

	dm := do.Account{
		Username: username,
	}
	dmPaging := do.Pagination{
		Page: pageInt,
	}

	result, err := h.AccountUseCase.GetAllAdmin(dm, dmPaging)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	// combine result
	combineResult := ToResGetAllAdminWithPaging(result)

	res := WebResponse{
		Message: "Success",
		Data:    combineResult,
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerAccount) DeleteAdminByID(c *gin.Context) {
	var req ReqIDActor
	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := do.Account{
		ID: req.ID,
	}

	result, err := h.AccountUseCase.DeleteAdminByID(dm)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := RowsAffected{
		Message:      "Success",
		RowsAffected: result,
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerAccount) UpdateAdminStatus(c *gin.Context) {
	var req ReqUpdateAdminStatus

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dmActor := do.Account{
		ID:         req.AdminID,
		IsVerified: req.IsVerified,
		IsActive:   req.IsActive,
	}
	dmAdminReg := do.Approval{
		AdminId: req.AdminID,
		Status:  req.Status,
	}

	result, err := h.AccountUseCase.UpdateAdminStatusByID(dmAdminReg, dmActor)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := RowsAffected{
		Message:      "Success",
		RowsAffected: result,
	}

	c.JSON(http.StatusOK, res)
}
