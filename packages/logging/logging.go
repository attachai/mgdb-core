package logging

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	cnst "github.com/attachai/mgdb-core/app/cnst"
	"github.com/attachai/mgdb-core/app/structs"
	"github.com/attachai/mgdb-core/utils"
	"github.com/sirupsen/logrus"
)

type LoggingServiceBackend struct{}

var Debug = cnst.Debug
var Info = cnst.Info
var Warning = cnst.Warning
var Error = cnst.Error
var Fatal = cnst.Fatal
var logconfig structs.LogConfiguration

func InitLog(logconfigInit structs.LogConfiguration) {
	logconfig = logconfigInit
	logLevel := logconfig.Level
	logrus.SetLevel(getLogLevel(logLevel))
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableColors: false,
	})
}

func sendLogger(data Jsonbodylog) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error("panic occurred:", err)
		}
	}()
	sendToServerLog(data)
}

func sendToServerLog(data Jsonbodylog) {
	path := logconfig.Server
	service := "api/logging"
	url := path + "/" + service

	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.Error("err", err)
	}
	var jsonStr = []byte(string(byteData))
	// log.Println("send log json :", string(byteData))
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
	// fmt.Printf("%+v\n", result)
}

func Logger(logLevel string, massage interface{}, saveLog_option ...bool) {
	saveLog := true
	if len(saveLog_option) > 0 {
		saveLog = saveLog_option[0]
	}
	isServerLog := logconfig.OnServerLog
	if logLevel == Debug {
		logrus.Debug(massage)
		if isServerLog && saveLog {
			verifyLogger(logLevel, massage)
		}
	} else if logLevel == Info {
		logrus.Info(massage)
		if isServerLog && saveLog {
			verifyLogger(logLevel, massage)
		}
	} else if logLevel == Warning {
		logrus.Warn(massage)
		if isServerLog && saveLog {
			verifyLogger(logLevel, massage)
		}
	} else if logLevel == Error {
		logrus.Error(massage)
		if isServerLog && saveLog {
			verifyLogger(logLevel, massage)
		}
	} else if logLevel == Fatal {
		logrus.Fatal(massage)
		if isServerLog && saveLog {
			verifyLogger(logLevel, massage)
		}
	}
}

func verifyLogger(logLevel string, massage interface{}) {
	var logObj Jsonbodylog
	logObj.App_id = logconfig.AppId
	logObj.App_name = logconfig.AppName
	logObj.Level = getLogLevel(logLevel).String()
	logObj.Message = massage

	isServerLog := logconfig.OnServerLog
	logLevelEnv := logconfig.Level
	if isServerLog {
		if logLevelEnv == logLevel && logLevelEnv == Debug {
			sendLogger(logObj)
		} else {
			state := getLogState(logLevelEnv)
			if utils.ContainInSlice(state, logLevel) {
				sendLogger(logObj)
			}
		}
	}

}
func getLogState(level string) []string {
	debugLvl := []string{Debug, Info, Warning, Error, Fatal}
	infoLvl := []string{Info, Warning, Error, Fatal}
	warningLvl := []string{Warning, Error, Fatal}
	errorLvl := []string{Error, Fatal}
	fatalLvl := []string{Fatal}
	switch level {
	case Debug:
		return debugLvl
	case Info:
		return infoLvl
	case Warning:
		return warningLvl
	case Error:
		return errorLvl
	case Fatal:
		return fatalLvl

	}
	return nil
}

func getLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case Debug:
		return logrus.DebugLevel
	case Info:
		return logrus.InfoLevel
	case Warning:
		return logrus.WarnLevel
	case Error:
		return logrus.ErrorLevel
	case Fatal:
		return logrus.FatalLevel
	}
	return logrus.DebugLevel

}
