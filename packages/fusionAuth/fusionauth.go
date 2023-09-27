package fusionAuth

import (
	cnst "github.com/attachai/mgdb-core/app/cnst"
	"github.com/attachai/mgdb-core/app/structs"
	coreStructs "github.com/attachai/mgdb-core/app/structs"
	"github.com/attachai/mgdb-core/utils"

	"encoding/json"
	"fmt"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
)

type Fusionauth struct{}

var Auth *fusionauth.FusionAuthClient

// DBConnection ..
func FusionAuthConnection(a *fusionauth.FusionAuthClient) {
	Auth = a
}

// /API is used to validate a JSON
func (f Fusionauth) ValidateJwt(token string) coreStructs.Jsonresponse {
	// initFusionAuth()
	var result coreStructs.Jsonresponse
	valResponse, error := Auth.ValidateJWT(token)
	var responseJSON []byte
	if error != nil {
		responseJSON, _ = json.Marshal(error)
	} else {
		responseJSON, _ = json.Marshal(valResponse)
	}
	fmt.Println("ValidateJwt:: ", string(responseJSON))
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}

	jwtMap := jsonMap["jwt"].(map[string]interface{})
	// var roles []string
	// for _, value := range jwtMap["roles"].(map[string]interface{}) {
	// 	roles = append(roles, fmt.Sprint(value))
	// }
	result.Results = structs.Validateresponse{
		UserId: fmt.Sprint(jwtMap["sub"]),
	}
	// fmt.Println("jsonMap=statusCode = ", jsonMap["statusCode"])

	if fmt.Sprint(jsonMap["statusCode"]) == "200" {
		result.Message = "The request was successful."
		result.StatusCode = "200"

	} else if fmt.Sprint(jsonMap["statusCode"]) == "401" {
		result.Message = "The access token is not valid. A new access token may be obtained by authentication against the Login API, the Token Endpoint or by utilizing a Refresh Token."
		result.StatusCode = "401"

	}

	return result
}

func (f Fusionauth) SendValidateJwt(token string) coreStructs.Jsonresponse {
	// initFusionAuth()
	var result coreStructs.Jsonresponse
	valResponse, error := Auth.ValidateJWT(token)
	var responseJSON []byte
	if error != nil {
		responseJSON, _ = json.Marshal(error)
	} else {
		responseJSON, _ = json.Marshal(valResponse)
	}
	fmt.Println("ValidateJwt:: ", string(responseJSON))
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}

	jwtMap := jsonMap["jwt"].(map[string]interface{})
	// var roles []string
	// for _, value := range jwtMap["roles"].(map[string]interface{}) {
	// 	roles = append(roles, fmt.Sprint(value))
	// }
	result.Results = structs.Validateresponse{
		UserId: fmt.Sprint(jwtMap["sub"]),
	}
	// fmt.Println("jsonMap=statusCode = ", jsonMap["statusCode"])

	if fmt.Sprint(jsonMap["statusCode"]) == "200" {
		result.Message = "The request was successful."
		result.StatusCode = "200"

	} else if fmt.Sprint(jsonMap["statusCode"]) == "401" {
		result.Message = "The access token is not valid. A new access token may be obtained by authentication against the Login API, the Token Endpoint or by utilizing a Refresh Token."
		result.StatusCode = "401"

	}

	return result
}

func (f Fusionauth) GetFusionUserId(token string) string {
	// initFusionAuth()

	valResponse, error := Auth.ValidateJWT(token)
	var responseJSON []byte
	if error != nil {
		responseJSON, _ = json.Marshal(error)
	} else {
		responseJSON, _ = json.Marshal(valResponse)
	}
	fmt.Println("ValidateJwt:: ", string(responseJSON))
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}

	jwtMap := jsonMap["jwt"].(map[string]interface{})

	return fmt.Sprint(jwtMap["sub"])
}

// /get user profile
func (f Fusionauth) UserProfile(userId string) coreStructs.Jsonresponse {

	// initFusionAuth()
	var result coreStructs.Jsonresponse
	userResponse, errors, err := Auth.RetrieveUser(userId)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		result.StatusCode = cnst.HTTP404
		result.Message = err.Error()
		return result
	}

	// Write the response from the FusionAuth client as JSON
	var responseJSON []byte
	if errors != nil {
		responseJSON, err = json.Marshal(errors)
	} else {
		responseJSON, err = json.Marshal(userResponse)
	}
	fmt.Println("RetrieveUser response;", string(responseJSON))
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}
	userMap := jsonMap["user"].(map[string]interface{})
	result.Results = userMap

	// fmt.Println("jsonMap=statusCode = ", jsonMap["statusCode"])

	if fmt.Sprint(jsonMap["statusCode"]) == "200" {
		result.Message = "Retrieve user was successful."
		result.StatusCode = "200"
	}

	return result
}

// / userId from fusionauth
func GetUserIdFusionAuth(token string) string {
	var userId string

	result := new(Fusionauth).ValidateJwt(token)
	responseJSON, _ := json.Marshal(result)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}

	jwtMap := jsonMap["results"].(map[string]interface{})
	userId = fmt.Sprint(jwtMap["userId"])

	return userId
}

// / Refresh a JWT
func (f Fusionauth) CheckRefreshToken(email string) coreStructs.Jsonresponse {
	// initFusionAuth()
	var result coreStructs.Jsonresponse
	appId := utils.ViperEnvVariable("AppId")
	var forgotPassData fusionauth.ForgotPasswordRequest
	forgotPassData.ApplicationId = appId
	forgotPassData.LoginId = email
	forgotPassData.SendForgotPasswordEmail = true

	userResponse, errors, err := Auth.ForgotPassword(forgotPassData)
	fmt.Println("ForgotPassword Response = ", userResponse)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ForgotPassword err.Error() = ", err.Error())
		result.StatusCode = cnst.HTTP404
		result.Message = err.Error()
		return result
	}

	// Write the response from the FusionAuth client as JSON
	var responseJSON []byte
	if errors != nil {

		responseJSON, err = json.Marshal(errors)
	} else {

		responseJSON, err = json.Marshal(userResponse)
	}
	fmt.Println("ForgotPassword User response;", string(responseJSON))
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(responseJSON, &jsonMap)
	if err != nil {
		panic(err)
	}

	if fmt.Sprint(jsonMap["statusCode"]) == "200" {
		result.Message = "The request was successful."
		result.StatusCode = "200"
		result.SaveStatus = true
		result.Results = jsonMap["verificationId"]
	} else {
		result.Message = "The request was invalid and/or malformed."
		result.StatusCode = "404"
		result.SaveStatus = false
		result.Results = nil
	}

	return result
}
