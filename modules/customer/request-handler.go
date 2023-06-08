package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/exception"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer/do"
	"github.com/imama2/bootcamp-bri-mini-project/utils/middleware"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	CustomerUseCase UseCaseCustomerInterface
}

func NewCostumerRequestHandler(uc UseCaseCustomerInterface) *RequestHandlerCustomer {
	return &RequestHandlerCustomer{
		CustomerUseCase: uc,
	}
}

func (h *RequestHandlerCustomer) RouteCustomer(app *gin.Engine) {
	// account
	a := NewCostumerRequestHandler(h.CustomerUseCase)
	a.RouteHandler(app)
}
func (h *RequestHandlerCustomer) RouteHandler(app *gin.Engine) {
	g := app.Group("/customer", middleware.Auth())

	g.GET("", h.GetAllCustomer)
	g.GET("/:id", h.GetCustomerByID)
	g.POST("", h.CreateCustomer)
	g.PUT("/:id", h.UpdateCustomerByID)
	g.DELETE("/:id", h.DeleteCustomerByID)
}

func (h *RequestHandlerCustomer) GetAllCustomer(c *gin.Context) {
	var req ReqGetAllCustomer

	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := do.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	dmPaging := do.Pagination{
		Page: pageInt,
	}

	result, err := h.CustomerUseCase.GetAllCustomer(dm, dmPaging)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	combineResult := ToResGetAllCustomerWithPaging(result)

	res := WebResponse{
		Message: "Success",
		Data:    combineResult,
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerCustomer) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}

	dm := do.Customer{
		ID: int64(idInt),
	}
	result, err := h.CustomerUseCase.GetCustomerByID(dm)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: "Success",
		Data:    ToResponseCustomer(result),
	}

	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	var req ReqAddCustomer

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := do.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}
	result, err := h.CustomerUseCase.CreateCustomer(dm)
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

func (h *RequestHandlerCustomer) UpdateCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}

	var req ReqAddCustomer

	err = c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := do.Customer{
		ID:        int64(idInt),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}
	result, err := h.CustomerUseCase.UpdateCustomerByID(dm)
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

func (h *RequestHandlerCustomer) DeleteCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}

	dm := do.Customer{
		ID: int64(idInt),
	}
	result, err := h.CustomerUseCase.DeleteCustomerByID(dm)
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
