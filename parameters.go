package sila

import "sila/domain"

type GetBusinessTypes interface {
	Do() (domain.GetBusinessTypesResponse, error)
}

type GetBusinessRoles interface {
	Do() (domain.GetBusinessRolesResponse, error)
}

type GetNaicsCategories interface {
	Do() (domain.GetNaicsCategoriesResponse, error)
}
