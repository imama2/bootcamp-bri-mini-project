package customer

type controllerCustomer struct {
	customerUseCase UseCaseCustomer
}

type ControllerCustomer interface {
	CreateCustomer(req CustomerParam) (any, error)
	//GetCustomerByID(id uint) (FindCustomer, error)
}
