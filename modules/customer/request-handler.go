package customer

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type RequestHandlerCustomer struct {
	ctr ControllerCustomer
}

func NewCostumerRequestHandler(dbCrud *gorm.DB) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctr: controllerCustomer{
			customerUseCase: useCaseCustomer{
				customerRepo: repositories.NewCustomer(dbCrud),
			},
		},
	}
}

func (h RequestHandlerCustomer) GetCustomerByID(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	custId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.GetCustomerByID(uint(custId))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
