package instance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ToastConfig struct
type ToastConfig struct {
	TokenID    string
	ExpireDate string
}

// Payload struct
type Payload struct {
	Auth Auth `json:"auth"`
}

// PasswordCredentials struct
type PasswordCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Auth struct
type Auth struct {
	TenantID            string              `json:"tenantId"`
	PasswordCredentials PasswordCredentials `json:"passwordCredentials"`
}

// RequestToken return token info
func RequestToken() (*ToastConfig, error) {
	tenantID, userName, passWord, _ := GetEnvparser().GetPasswordCredentials()
	var tokenURL = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"
	data := Payload{
		Auth: Auth{
			TenantID: tenantID,
			PasswordCredentials: PasswordCredentials{
				Username: userName,
				Password: passWord,
			},
		},
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logrus.Fatal("Unable to encode data: ", data)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", tokenURL, body)
	if err != nil {
		logrus.Fatal("Failed to generate request: " + tokenURL)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatal("Request Failed.")
	}

	defer resp.Body.Close()

	allRespBytes, _ := ioutil.ReadAll(resp.Body)
	allRespMap := make(map[string]interface{})
	if err := json.Unmarshal(allRespBytes, &allRespMap); err != nil {
		logrus.Fatal("Decode failed: ", allRespBytes)
	}

	tokenInfo := allRespMap["access"].(interface{}).(map[string]interface{})["token"]

	toastConfig := ToastConfig{}
	toastConfig.TokenID = fmt.Sprint(tokenInfo.(interface{}).(map[string]interface{})["id"])
	toastConfig.ExpireDate = fmt.Sprint(tokenInfo.(interface{}).(map[string]interface{})["expires"])
	return &toastConfig, nil
}

// RequestInstanceDetails return instanceDetail map[string]interface{}
func RequestInstanceDetails(token *ToastConfig, tenantID string, region string) (map[string]interface{}, error) {

	var baseURL = "https://" + region + "-api-instance.infrastructure.cloud.toast.com/v2/" + tenantID + "/servers/detail"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		logrus.Fatal("Failed to generate request: " + baseURL)
	}

	req.Header.Set("X-Auth-Token", token.TokenID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatal("No region have been selected. Please select kr1, kr2, or jp1.")
	}

	defer resp.Body.Close()
	allRespBytes, _ := ioutil.ReadAll(resp.Body)
	allRespMap := make(map[string]interface{})
	if err := json.Unmarshal(allRespBytes, &allRespMap); err != nil {
		logrus.Fatal("Decode failed: ", allRespBytes)
	}
	return allRespMap, nil
}

// var region string = cmd.Region

//PostInstanceStatus request Post to start(or stop) instance
func PostInstanceStatus(action string, serverID string, region string) int {
	tenantID, _, _, _ := GetEnvparser().GetPasswordCredentials()
	tokenInfo, _ := RequestToken()

	var requestURL = "https://" + region + "-api-instance.infrastructure.cloud.toast.com/v2/" + tenantID + "/servers/" + serverID + "/action"

	client := &http.Client{}
	var data *strings.Reader
	if action == "start" {
		data = strings.NewReader(`{ "os-start" : null } `)
	} else {
		data = strings.NewReader(`{ "os-stop" : null } `)
	}

	req, err := http.NewRequest("POST", requestURL, data)
	if err != nil {
		logrus.Fatal("Failed to generate request: " + requestURL)
	}
	req.Header.Set("X-Auth-Token", tokenInfo.TokenID)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("The request to stop or start the instance failed. Check the header. ")
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
