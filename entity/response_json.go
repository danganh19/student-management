package entity

type ReponseJson struct {
	Page     int `json:"page" bson:"page"`
	PageSize int `json:"pageSize" bson:"page_size"`
	Total    int `json:"total" bson:"total"`
}