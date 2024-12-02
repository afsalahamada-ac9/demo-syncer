package rds_to_sf_entity

type Timing_rds struct {
	End_time   string `json:"End_Time__c"`
	Start_time string `json:"Start_Time__c"`
	End_date   string `json:"End_Date__c"`
	Start_date string `json:"Start_Date__c"`
	Course     string `json:"Event__c"`
}

type Timing_Data struct {
	Operation string     `json:"operation"`
	Value     Timing_rds `json:"value"`
}
