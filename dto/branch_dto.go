package dto

import (
	"ims-be/dao"
	model "ims-be/entity"
)

type BranchData struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// Interface
type IBranchDto interface {
	CreateBranchDto(data *BranchData) (err error)
	GetBranchByIdDto(IdProduct int) (product *model.Branch, err error)
	GetBranchesDtos(Page int, PageSize int) (products []*model.Branch, meta *model.ReponseJson, err error)
	UpdateBranchDto(IdProduct int, Name string) (err error)
}

// Extend method from dao
type branchDao struct {
	dao.IBranchDao
}

// Create branch from data to object
func (bR *branchDao) CreateBranchDto(data *BranchData) (err error) {
	branch := model.Branch{}
	branch.Name = data.Name
	err = bR.CreateBrancheDaos(&branch)
	return err
}

func (bR *branchDao) GetBranchByIdDto(IdProduct int) (product *model.Branch, err error) {
	return bR.GetBranchesByIdDaos(IdProduct)
}

func (bR *branchDao) GetBranchesDtos(Page int, PageSize int) (products []*model.Branch, meta *model.Meta, err error) {
	var total int
	products, total, err = bR.GetBranchesDaos(Page, PageSize)
	meta = &model.ReponseJson{
		Page:     Page,
		PageSize: PageSize,
		Total:    total,
	}
	return products, meta, err

}

func (bR *branchDao) UpdateBranchDto(IdProduct int, Name string) (err error) {
	return bR.UpdateBranchesDaos(IdProduct, Name)
}

func BranchDto() (IBranchDto, error) {
	return &branchDao{
		 ,
	}, nil
}
