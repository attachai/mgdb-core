package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/attachai/core/packages/setting"
	"github.com/attachai/core/utils"
	"github.com/sirupsen/logrus"
)

type LoggingServiceBackend struct{}

func SaveLog(data Jsonbodylog) {
	path := utils.GetEnvVariable("ServerLogging")
	service := "api/logging"
	url := path + "/" + service

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	var jsonStr = []byte(string(byteData))
	log.Println("send log json :", string(byteData))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", os.Getenv("token"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result JsonResponseEror
	json.Unmarshal([]byte(body), &result)
	fmt.Printf("%+v\n", result)
}

func Logger(logTax string, input interface{}) {
	// var bodyMap map[string]interface{}
	logLevel := utils.GetEnvVariable("logLevel")
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	if logLevel == logTax && logLevel == setting.LogLevelSetting.Debug {
		log.Println(input)
		// logBody := logger.WithFields(logrus.Fields{
		// 	"app_id":   setting.AppSetting.AppId,
		// 	"app_name": setting.AppSetting.AppName,
		// })
		// logBody.Level = getLogLevel(logTax)
		// bye, _ := logBody.Bytes()
		// err := json.Unmarshal(bye, &bodyMap)
		// if err != nil {
		// 	panic(err)
		// }
		// bodyMap["message"] = input
		// bodyMap["time"] = time.Now().Format("01-02-2006 15:04:05")
		var logObj Jsonbodylog
		logObj.App_id = utils.GetEnvVariable("AppID")
		logObj.App_name = utils.GetEnvVariable("AppName")
		logObj.Level = getLogLevel(logTax).String()
		logObj.Message = input
		log.Println("send data : ", logObj)
		SaveLog(logObj)
	} else {
		///เขียนอย่างเดียว
		state := getLogState(logLevel)
		if utils.ContainInSlice(state, logTax) {
			// logBody := logger.WithFields(logrus.Fields{
			// 	"app_id":   setting.AppSetting.AppId,
			// 	"app_name": setting.AppSetting.AppName,
			// })
			// logBody.Level = getLogLevel(logTax)
			// bye, _ := logBody.Bytes()
			// err := json.Unmarshal(bye, &bodyMap)
			// if err != nil {
			// 	panic(err)
			// }
			// bodyMap["msg"] = input
			// bodyMap["time"] = time.Now().Format("01-02-2006 15:04:05")
			var logObj Jsonbodylog
			logObj.App_id = utils.GetEnvVariable("AppID")
			logObj.App_name = utils.GetEnvVariable("AppName")
			logObj.Level = getLogLevel(logTax).String()
			logObj.Message = input
			log.Println("send data : ", logObj)
			SaveLog(logObj)
		}
	}
}

func getLogState(level string) []string {
	debugLvl := []string{"debug", "info", "warning", "error", "fatal"}
	infoLvl := []string{"info", "warning", "error", "fatal"}
	warningLvl := []string{"warning", "error", "fatal"}
	errorLvl := []string{"error", "fatal"}
	fatalLvl := []string{"fatal"}
	switch level {
	case "debug":
		return debugLvl
	case "info":
		return infoLvl
	case "warning":
		return warningLvl
	case "error":
		return errorLvl
	case "fatal":
		return fatalLvl

	}
	return nil
}

func getLogLevel(logTag string) logrus.Level {
	switch logTag {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel

	}
	return logrus.DebugLevel

}
