package response

import model "ims-be/entity"

type BranchResponse struct {
	Message string `json:"message" bson:"message"`
	Success bool   `json:"success" bson:"success"`
}

type BranchDetailResponse struct {
	Data    *model.Branch `json:"data" bson:"data"`
	Message string         `json:"message" bson:"message"`
	Success bool           `json:"success" bson:"success"`
}

type ListBranchResponse struct {
	Data    []*model.Branch `json:"data" bson:"data"`
	Meta    *model.ReponseJson      `json:"meta" bson:"meta"`
	Message string           `json:"message" bson:"message"`
	Success bool             `json:"success" bson:"success"`
}

type UpdateBranch struct {
	Name string `json:"name" bson:"name"`
}