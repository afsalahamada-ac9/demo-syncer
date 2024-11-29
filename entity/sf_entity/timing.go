package entity

type Timing struct {
	Course_id   int    `json:"Name"` // todo: map with id from course
	Ext_id      string `json:"Id"`
	Course_date string `json:"Start_Date__c"`
	Start_time  string `json:"Start_Time__c"`
	End_time    string `json:"End_Time__c"`
	Updated_at  string `json:"LastModifiedDate"`
	Created_at  string `json:"CreatedDate"`
}

func (*Timing) TableName() string {
	return "course_timing"
}
