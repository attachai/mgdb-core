package model_name_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"

	modelCtrl "github.com/attachai/mgdb-core/app/models/model_name_db/db/controllers"
	"github.com/attachai/mgdb-core/app/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceModel struct{}

// requestType(POST,GET,PUT,Delete)
func (s ServiceModel) SendApi(data []byte, url string, requestType string, header string) string {

	// url := "http://phr.mch.mfu.ac.th/servicedev/rest/rsservice/dataOperatePost"
	log.Println("URL:>", url)
	var jsonStr = []byte(string(data))
	log.Println("jsonStr post:", string(data))
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	if len(header) != 0 {
		req.Header.Add("token", header)
	} else {
		req.Header.Add("token", "")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log.Println("response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// log.Println("response Body:", string(body))

	return string(body)
}

func GetDocumentId(filterName string, filterVal interface{}, Collection string) string {
	var docId string
	var result structs.Jsonresponse
	var jsonPost structs.JsonService

	jsonPost.Collection = Collection
	conMap := make(map[string]interface{})
	conMap[filterName] = filterVal
	jsonPost.Condition = conMap
	proMap := make(map[string]interface{})
	proMap["id"] = 1
	jsonPost.Projection = proMap
	byteArray, err := json.Marshal(jsonPost)
	if err != nil {

		log.Fatal(err)
	}
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// Restore the io.ReadCloser to its original state
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteArray))
	modelCtrl := new(modelCtrl.ReadController)
	saveStatus, resultData := modelCtrl.FindDocument(c, false) //// read document
	log.Println(resultData)

	if saveStatus == true {
		result.Message = "ok"
		result.StatusCode = "200"

		result.Results = resultData
		for _, v := range resultData.([]primitive.M) {
			d := reflect.ValueOf(v)
			for _, value := range d.MapKeys() {
				reflectVal := d.MapIndex(value).Interface()
				docId = fmt.Sprint(reflectVal)
			}

		}

	}

	return docId
}

// / send image to upload
func (s ServiceModel) SendPostFileRequest(url string, filename string, filetype string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func (s ServiceModel) SendApiHeaders(data []byte, url string, requestType string, headers map[string]interface{}) (string, string) {

	// url := "http://phr.mch.mfu.ac.th/servicedev/rest/rsservice/dataOperatePost"
	log.Println("URL:>", url)
	var jsonStr = []byte(string(data))
	log.Println("jsonStr post:", string(data))
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	if len(headers) != 0 {
		for k, v := range headers {
			req.Header.Add(k, fmt.Sprint(v))
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// return "" , ""
		panic(err)
	}
	defer resp.Body.Close()

	// log.Println("response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("3")
		panic(err)
	}
	// log.Println("response Body:", string(body))

	return string(body), string(resp.Status)
}

func (s ServiceModel) SendApiBasicAuten(data []byte, url string, requestType string, headers map[string]interface{}) (string, string) {

	// url := "http://phr.mch.mfu.ac.th/servicedev/rest/rsservice/dataOperatePost"
	log.Println("URL:>", url)
	var jsonStr = []byte(string(data))
	log.Println("jsonStr post:", string(data))
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	//userName := "apikey"
	pass := "VFTzZpLC4eZnbfoRYPnoWl9ptpsJVgsW"
	//req.Header.Add("apikey","")
	//req.SetBasicAuth(userName, pass)
	// req.Header.Values(userName)
	// req.Header.Add()
	// req.Header.Add("Authorization", "API-key")
	// req.Header.Add("Add-to", "Header")
	req.Header.Add("apikey", pass)
	//req.Header.Add("Value", pass)

	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")
	if len(headers) != 0 {
		for k, v := range headers {
			req.Header.Add(k, fmt.Sprint(v))
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// return "" , ""
		panic(err)
	}
	defer resp.Body.Close()

	//log.Println("response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("3")
		panic(err)
	}
	//log.Println("response Body:", string(body))

	return string(body), string(resp.Status)
}

func (s ServiceModel) SendApiNotify(data []byte, url string, requestType string, ChannelToken string) (string, string) {

	// url := "https://api.line.me/v2/bot/message/push"
	log.Println("URL:>", url)

	var jsonStr = []byte(string(data))
	log.Println("jsonStr post:", string(data))
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)
	//req.Header.Set("X-Line-Retry-Key","Ueeee552cb3ec9f2c5f617b84600003c3")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return string(body), string(resp.Status)
}

func (s ServiceModel) SendApiKey(data []byte, url string, requestType string, headers map[string]interface{}) (string, string) {

	// url := "http://phr.mch.mfu.ac.th/servicedev/rest/rsservice/dataOperatePost"
	log.Println("URL:>", url)
	var jsonStr = []byte(string(data))
	log.Println("jsonStr post:", string(data))
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	//userName := "apikey"

	pass := ""
	if headers["password"] != nil {
		pass = fmt.Sprint(headers["password"])
	}

	if pass != "" {
		req.Header.Add("apikey", pass)
	}
	//req.Header.Add("Value", pass)

	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")
	if len(headers) != 0 {
		for k, v := range headers {
			req.Header.Add(k, fmt.Sprint(v))
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// return "" , ""
		panic(err)
	}
	defer resp.Body.Close()

	//log.Println("response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("3")
		panic(err)
	}
	//log.Println("response Body:", string(body))

	return string(body), string(resp.Status)
}
