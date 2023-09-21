package conf

import (
	"log"

	"github.com/spf13/viper"
)

// var Config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func InitConfig(env string) {
	var err error
	viper.New()
	viper.SetConfigName(env) // ชื่อไฟล์ Config

	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf") // path จะให้ค้นหากี่ที่ก็ได้ แค่เรียกคำสั่งนี้ไปอีกก็พอ
	// viper.AddConfigPath("ost.utils/conf/") // path ที่ให้ค้นหาไฟล์ Config
	// การค้นหาตาม Path จะเรียงลำดับตาม Path ที่ถูกเพิ่มก่อน ถ้าเจอก็จะหยุดค้นหาใน Path ต่อไปเลย

	// เริ่มการค้นหาไฟล์ Config และอ่านไฟล์
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

}
