package models

type Item struct {
	ID                    uint   `gorm:"primary key;autoIncrement" json:"ID"`
	Name                  string `json:"name"`
	Kind                  string `json:"kind"`
	Place                 string `json:"place"`
	ExpectedMonetaryValue string `json:"expectedMonetaryValue"`
	Image                 string `json:"image"`
}

type APIResponse struct {
	id      string `json:"id"`
	message string `json:"message"`
}
