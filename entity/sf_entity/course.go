package entity

type Course struct {
	Url           string   `json:"url" gorm:"url"`
	Max_attendees int      `json:"Max_attendees__c" gorm:"max_attendees"`
	Address       Location `json:"Address" gorm:"address"`
	Tenant_id     int      `json:"Tenant_id" gorm:"tenant_id"`
	Ext_id        string   `json:"Id" gorm:"ext_id"`
	Name          string   `json:"Name" gorm:"name"`
	Timezone      string   `json:"Timezone__c" gorm:"timezone"`
	Mode          string   `json:"Mode" gorm:"mode"`
	Center_id     int      `json:"Location__c" gorm:"center_id"`
	Status        string   `json:"Status__c" gorm:"status"`
	Created_at    string   `json:"CreatedDate" gorm:"created_at"`
	Num_attendees int      `json:"Number_Of_Students__c" gorm:"num_attendees"`
	Product_id    int      `json:"Workshop_Type__c" gorm:"product_id"`
	Updated_at    string   `json:"LastModifiedDate" gorm:"updated_at"`
	Notes         string   `json:"Notes__c" gorm:"notes"`
	Short_url     string   `json:"Short_url" gorm:"short_url"`
}
