package entity

type Product struct {
	Updated_at          string `json:"LastModifiedDate" gorm:"column:updated_at"`
	Created_at          string `json:"CreatedDate" gorm:"column:created_at"`
	Is_deleted          bool   `json:"IsDeleted" gorm:"column:is_deleted"`
	Format              string `json:"Online_Or_In_Person__c" gorm:"column:format"`
	Max_Attendees       int32  `json:"Max_Attendees__c,omitempty" gorm:"column:max_attendees"`
	Listing_Visibity    string `json:"Listing_Visibity__c,omitempty" gorm:"column:visibility"` // todo: rename visibity to visibility -> spelling mistake in sf
	Event_Duration      int32  `json:"Event_Duration__c,omitempty" gorm:"column:duration_days"`
	Product             string `json:"Product__c,omitempty" gorm:"column:base_product_id"`
	CType               string `json:"CType_Id__c" gorm:"column:ctype"`
	Title               string `json:"Title__c" gorm:"column:title"`
	Name                string `json:"name" gorm:"column:name"`
	TenantID            int32  `json:"Tenant_id" gorm:"column:tenant_id"`
	ExtID               string `json:"Id" gorm:"column:ext_id"`
	ExtName             string `json:"Name" gorm:"column:Name"`
	Base_product_ext_id string `json:"base_product_ext_id" gorm:"base_product_ext_id"`
	Is_auto_approve     bool   `json:"Auto_Approve_Event__c" gorm:"column:is_auto_approve"`
}

func (Product) TableName() string {
	return "product"
}
