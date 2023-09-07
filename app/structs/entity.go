package structs

type Jsonresponse struct {
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Results    interface{} `json:"results"`
	SaveStatus bool        `json:"saveStatus"`
}

//Jsonbody to service
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
}

type NameInfo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	IsDefault  *bool `json:"isDefault,omitempty"`
}
