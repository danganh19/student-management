package dao

import "ims-be/entity"

type IBranchDao interface {
	CreateBrancheDaos(branches *entity.Branch) error
	GetBranchesByIdDaos(branchId int) (branches *entity.Branch, err error)
	GetBranchesDaos(page int, pageSize int) (branches []*entity.Branch, total int, err error)
	UpdateBranchesDaos(branchId int, name string) (err error)
}