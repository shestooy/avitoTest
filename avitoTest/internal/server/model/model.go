package model

type Tender struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Version        uint   `json:"verstion"`
	ServiceType    string `json:"serviceType"`
	CreatedAt      string `json:"createdAt"`
	OrganizationID string `json:"organizationId"`
	UsernameID     string `json:"creatorUserName"`
}

type Bids struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	TenderId        string `json:"tenderId"`
	OrganizationId  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
	AuthorID        string `json:"authorID"`
	AuthorType      string `json:"authorType"`
	CreatedAt       string `json:"createdAt"`
	Version         uint   `json:"version"`
}
