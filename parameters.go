package sila

import "github.com/kappapay/silaclient/domain"

type GetBusinessTypes interface {
	Do() (domain.GetBusinessTypesResponse, error)
}

type GetBusinessRoles interface {
	Do() (domain.GetBusinessRolesResponse, error)
}

type GetNaicsCategories interface {
	Do() (domain.GetNaicsCategoriesResponse, error)
}
