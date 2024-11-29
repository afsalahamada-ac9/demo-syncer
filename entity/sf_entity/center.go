// this represents the location entity in salesforce
package entity

type Center struct {
	Ext_id             string      `json:"Id" gorm:"column:ext_id"`
	Tenant_id          int         `json:"Tenant_id" gorm:"column:tenant_id"`
	Ext_name           string      `json:"Name" gorm:"column:ext_name"`
	Address            Location    `json:"address" gorm:"column:address"`
	Geo_Location       GeoLocation `json:"geolocation" gorm:"column:geo_location"`
	Capacity           int         `json:"Max_Capacity__c" gorm:"column:capacity"`
	Mode               string      `json:"Center_Mode__c" gorm:"column:mode"`
	Webpage            string      `json:"Center_URL__c" gorm:"column:webpage"`
	Is_national_center bool        `json:"Is_National_Center__c" gorm:"column:is_national_center"`
	Is_enabled         bool        `json:"Is_enable__c" gorm:"column:is_enabled"`
	Created_at         string      `json:"CreatedDate" gorm:"column:created_at"`
	Updated_at         string      `json:"UpdatedDate" gorm:"column:updated_at"`
}

func (*Center) TableName() string {
	return "center"
}
