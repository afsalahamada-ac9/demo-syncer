package entity

type SF struct {
	Operation string     `json:"operation"`
	Value     EventValue `json:"value"`
}

// data sent from sf
type EventValue struct {
	CoTeacher2     string `json:"Co_Teacher_2__c"`
	CoTeacher1     string `json:"Co_Teacher_1__c"`
	PrimaryTeacher string `json:"Primary_Teacher__c"`
	Email          string `json:"Email__c"`
	Status         string `json:"Status__c"`
	Notes          string `json:"Notes__c"`
	Name           string `json:"Name"`
	ID             string `json:"Id"`
}
