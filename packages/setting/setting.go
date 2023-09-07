package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ApiMsg struct{}

type App struct {
	JwtSecret   string
	AppId       string
	AppKey      string
	AppName     string
	UseBlueposh bool
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	HTTP200     string
	HTTP201     string
	HTTP203     string
	HTTP206     string
	HTTP302     string
	HTTP400     string
	HTTP401     string
	HTTP403     string
	HTTP404     string
	HTTP405     string
	HTTP408     string
	HTTP415     string
	HTTP429     string
	HTTP500     string
	HTTP502     string
	HTTP503     string
	HTTP504     string
	BasePath    string
	ServiceName string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type ApiGroup struct {
	AddressGroup          string
	LogGroup              string
	ImageServiceGroup     string
	BluePoshGroup         string
	V1Group               string
	
	SmsGroup              string
	LineBotUserGroup      string
	AppointmentSetupGroup string
	MedicalGroup          string
	EmployeeGroup         string
	UserGroup             string
	TlphGroup             string
}

var ApiGroupSetting = &ApiGroup{}

type ApiEndpoint struct {
	AddressApiPath                   string
	LogApiPath                       string
	ImageUpload                      string
	ImageUrl                         string
	BluePoshPersonalEndPoint         string
	BluePoshPersonalLoginEndPoint    string
	ImedAutenEndPoint                string
	ImedProfilePersonalEndPoint      string
	ImedProfilePetPersonalIdEndPoint string
	ImedProfilePetIdEndPoint         string
	ImedHistoryVaccineEndPoint       string
	ImedHistoryDrugEndPoint          string
	ImedHistoryDiseaseEndPoint       string
	ImedHistoryDrugIgnoreEndPoint    string
	RequestOtpEndPoint               string
	VerifyOtpEndPoint                string
	ImedPettype                      string
	ImedPetCurrentVisitEndPoint      string
	ImedHistoryDrugNoteEndPoint      string
	ImedSatisfactionEndPoint         string
	ImedCalendarPersonalEndpoint     string
	MedicalPersonalEndPoint          string
	MedicalHistoryPersonalEndPoint   string
	EmployeePersonalEndPoint         string
	UserPersonalEndPoint             string
	ImageDeleteEndPoint              string
	ImedMembertierEndPoint           string
	UidEndPoint						 string
	NodeRestUpdaterichEndPoint       string
}

var ApiEndpointSetting = &ApiEndpoint{}

type IpAddress struct {
	AddressContextPath          string
	LogContextPath              string
	ImageContextPath            string
	BluePoshContextPath         string
	BluePoshMiddleContextPath   string
	TestnoderestContextPath     string
	ProfilePersonalContextPath  string
	SmsContextPath              string
	LineBotContextPath          string
	AppointmentSetupContextPath string
	TlphContextPath             string
	MedicalContextPath          string
	EmployeeContextPath         string
	UserContextPath             string
}

var IpAddressSetting = &IpAddress{}

type HttpMethod struct {
	PostMethod   string
	GetMethod    string
	PutMethod    string
	DeleteMethod string
}

var HttpMethodSetting = &HttpMethod{}

type Collection struct {
	User               string
	PersonalDev        string
	AddressDev         string
	EmergencyContact   string
	Cocare             string
	CocareDaily        string
	HuaweiAuten        string
	PersonalRelation   string
	ImedPetImage       string
	ImedPreRegister    string
	ImedPreRegisterPet string
}

var CollectionSetting = &Collection{}

type LogLevel struct {
	Debug   string
	Warning string
	Info    string
	Error   string
	Fatal   string
}

var LogLevelSetting = &LogLevel{}

var cfg *ini.File

// Setup initialize the configuration instance
func (a ApiMsg) Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)

	mapTo("apiGroup", ApiGroupSetting)
	mapTo("apiEndpoint", ApiEndpointSetting)
	mapTo("ipAddress", IpAddressSetting)
	mapTo("httpMethod", HttpMethodSetting)
	mapTo("collection", CollectionSetting)
	mapTo("logLevel", LogLevelSetting)

	readTimeStr := fmt.Sprint(viper.Get("ReadTimeout"))
	writeTimeStr := fmt.Sprint(viper.Get("WriteTimeout"))

	// Using ParseDuration() function
	readTime, _ := time.ParseDuration(readTimeStr)
	writeTime, _ := time.ParseDuration(writeTimeStr)

	ServerSetting.ReadTimeout = readTime * time.Second
	ServerSetting.WriteTimeout = writeTime * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func (a ApiMsg) EnvSetup() {
	godotenv.Load()
}
