package structs

import "time"

type Jsonresponse struct {
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Results    interface{} `json:"results"`
	SaveStatus bool        `json:"saveStatus"`
}

// Jsonbody to service
type JsonService struct {
	Collection   string      `json:"collection"`
	Reference    string      `json:"reference"`
	Condition    interface{} `json:"condition"`
	Data         interface{} `json:"data"`
	Projection   interface{} `json:"projection"`
	ArrayFilter  interface{} `json:"arrayFilter"`
	Sort         interface{} `json:"sort"`
	Limit        int         `json:"limit"`
	Offset       int         `json:"offset"`
	Timezone     string      `json:"timezone"`
	Atomicity    bool        `json:"atomicity"`
	Duplicate    bool        `json:"duplicate"`
	Multi        bool        `json:"multi"`
	Replacement  bool        `json:"replacement"`
	UpdateFilter interface{} `json:"updateFilter"`
}
type Loginresponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"userId"`
	PersonalId   string `json:"personalId"`
}

type LoginresponseToken struct {
	Token               string `json:"token"`
	RefreshToken        string `json:"refreshToken"`
	UserId              string `json:"userId"`
	AuthenticationToken string `json:"authenticationToken"`
}

type Validateresponse struct {
	UserId string `json:"userId"`
	AppId  string `json:"appId"`
}

type Image struct {
	Id       string `json:"id,omitempty"`
	Url      string `json:"url,omitempty"`
	FileType string `json:"fileType,omitempty"`
	Name     string `json:"name,omitempty"`
}

type NameInfo struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsDefault *bool  `json:"isDefault,omitempty"`
}

type DataInfo struct {
	Id        string `json:"id,omitempty"`
	Value     string `json:"value,omitempty"`
	IsDefault *bool  `json:"IsDefault,omitempty"`
}

type UserActivity struct {
	RefId      string `json:"refId"`
	LastUpdate string `json:"LastUpdate"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Req_method string `json:"req_method"`
}

//LogService
type LogService struct {
	App_id          string      `json:"app_id"`
	App_name        string      `json:"app_name"`
	Level           string      `json:"level"`
	Req_method      string      `json:"req_method"`
	Req_uri         string      `json:"req_uri"`
	Ref_id          string      `json:"ref_id"`
	Collection_name string      `json:"collection_name"`
	Req_message     interface{} `json:"req_message"`  /// body
	Resp_message    interface{} `json:"resp_message"` /// response
	User            interface{} `json:"user"`         /// userinfo
	Created_at      time.Time   `json:"created_at"`
	Updated_at      time.Time   `json:"updated_at"`
}

type Recorder struct {
	RefId       string `json:"refId"`
	PersonalId  string `json:"personalId"`
	CompanyId   string `json:"companyId"`
	CompanyName string `json:"companyName"`
	BranchId    string `json:"branchId"`
	BranchName  string `json:"branchName"`
	LastUpdate  string `json:"lastUpdate"`
}

type LogConfiguration struct {
	Server      string `json:"id,omitempty"`
	AppId       string `json:"appId,omitempty"`
	AppName     string `json:"appName,omitempty"`
	Level       string `json:"logLevel,omitempty"`
	OnServerLog bool   `json:"onServerLog,omitempty"`
}
