package account

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities/account"
	domain "github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
)

func DTOAccount(et entity.Account) domain.Account {
	return domain.Account{
		ID:         et.ID,
		Username:   et.Username,
		Password:   et.Password,
		RoleID:     et.RoleId,
		IsVerified: et.IsVerified,
		IsActive:   et.IsActive,
		CreatedAt:  et.CreatedAt,
		UpdatedAt:  et.UpdatedAt,
	}
}

func DTOAccountList(et []entity.Account) []domain.Account {
	var result []domain.Account = make([]domain.Account, 0)
	for _, v := range et {
		result = append(result, DTOAccount(v))
	}

	return result
}

func DTOApproval(et entity.Approval) domain.Approval {
	return domain.Approval{
		ID:           et.ID,
		AdminId:      et.AdminId,
		SuperAdminID: et.SuperAdminId,
		Status:       et.Status,
	}
}

func DTOListApprovalAdd(et []entity.Approval) []domain.Approval {
	result := make([]domain.Approval, 0)
	for _, v := range et {
		result = append(result, DTOApproval(v))
	}

	return result
}
