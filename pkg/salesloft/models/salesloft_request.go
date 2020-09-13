package models

type SalesloftRequest struct{

	Metadata interface{}  	`json:"metadata"`
	Data	 []People		`json:"data"`

}
