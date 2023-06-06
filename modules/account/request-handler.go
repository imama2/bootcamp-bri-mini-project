package account

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type RequestHandlerAccount struct {
	ctr ControllerAccount
}

func NewAccountRequestHandler(dbCrud *gorm.DB) RequestHandlerAccount {
	return RequestHandlerAccount{
		ctr: controllerAccount{
			accountUseCase: useCaseAccount{
				accountRepo: repositories.NewAccount(dbCrud)},
		},
	}
}

func (h RequestHandlerAccount) GetAccountByID(c *gin.Context) {
	request := AccountParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	custId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.GetAccountByID(uint(custId))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) CreateAccount(c *gin.Context) {
	request := AccountParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateAccount(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RequestHandlerAccount) GetAccountByUsernameAndPassword(c *gin.Context) {
	request := AccountParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := h.ctr.GetAccountByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}
