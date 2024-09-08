package model

type Tender struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	ServiceType    string `json:"serviceType"`
	CreatedAt      string `json:"createdAt"`
	OrganizationID string `json:"organizationId"`
	UsernameID     string `json:"creatorUserName"`
}
