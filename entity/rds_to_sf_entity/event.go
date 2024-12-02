package rds_to_sf_entity

type Event_rds struct {
	Ext_Id             string `json:"Ext_Id"`
	Number_Of_Students int    `json:"Number_Of_Students__c"`
	CType_Id           string `json:"CType_Id__c"`
	Location           string `json:"Location__c"`
	Timezone           string `json:"Timezone__c"`
	Max_Attendees      int    `json:"Max_Attendees__c"`
	Country            string `json:"Country__c"`
	Zip_Postal_Code    string `json:"Zip_Postal_Code__c"`
	State              string `json:"State__c"`
	City               string `json:"City__c"`
	Street_Address_2   string `json:"Street_Address_2__c"`
	Street_Address_1   string `json:"Street_Address_1__c"`
	Status             string `json:"Status__c"`
	Notes              string `json:"Notes__c"`
	Workshop_Type      string `json:"Workshop_Type__c"`
	Event_Start_Date   string `json:"Event_Start_Date__c"`
	Event_End_Date     string `json:"Event_End_Date__c"`
}

type Event_Data struct {
	Operation string    `json:"operation"`
	Value     Event_rds `json:"value"`
}
