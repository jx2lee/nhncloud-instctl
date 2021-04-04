package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/jx2lee/nhncloud-instctl/pkg/instance"
)

type ToastConfig = instance.ToastConfig
type PayloadAttachFIP struct {
	Floatingip Floatingip `json:"floatingip"`
}
type Floatingip struct {
	PortID string `json:"port_id"`
}


func requestToken() (*ToastConfig, error) {
	tenantID, userName, passWord, _ := instance.GetEnvparser("rds").GetPasswordCredentials()
	tokenURL := "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

	data := instance.Payload{
		Auth: instance.Auth{
			TenantID: tenantID,
			PasswordCredentials: instance.PasswordCredentials{
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
		logrus.Fatal("Request Failed: " + tokenURL)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatal("Unable to use client")
	}

	defer resp.Body.Close()

	allRespBytes, _ := ioutil.ReadAll(resp.Body)
	allRespMap := make(map[string]interface{})
	if err := json.Unmarshal(allRespBytes, &allRespMap); err != nil {
		logrus.Fatal(err)
	}

	tokenInfo := allRespMap["access"].(interface{}).(map[string]interface{})["token"]

	toastConfig := instance.ToastConfig{}
	toastConfig.TokenID = fmt.Sprint(tokenInfo.(interface{}).(map[string]interface{})["id"])
	toastConfig.ExpireDate = fmt.Sprint(tokenInfo.(interface{}).(map[string]interface{})["expires"])
	return &toastConfig, nil
}

func GetIaasInstancePortInfo(iaasInstanceId string) {
	tokenInfo, err := requestToken()
	logrus.Info("Token-ID: ", tokenInfo.TokenID)
	if err != nil {
		logrus.Fatal("Failed Get Token.")
	}

	var baseURL = "http://iaas.tcc1.cloud.toastoven.net/network/v2.0/ports?device_id=" + iaasInstanceId
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		logrus.Fatal("Request Failed: " + baseURL)
	}

	req.Header.Set("X-Auth-Token", tokenInfo.TokenID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatal("Request Failed.")
	}

	defer resp.Body.Close()
	allRespBytes, _ := ioutil.ReadAll(resp.Body)
	allRespMap := make(map[string]interface{})
	if err := json.Unmarshal(allRespBytes, &allRespMap); err != nil {
		logrus.Fatal(err)
	}

	portInfo, err := prettyPrintJSON(allRespBytes)
	if err != nil {
		logrus.Fatal("Failed to change to json format.")
	}

	fmt.Printf("%s", portInfo)
}

func AttachFip(floatingIpId string, portId string) {
	tokenInfo, err := requestToken()
	if err != nil {
		logrus.Fatal("Failed Get Token.")
	}

	data := PayloadAttachFIP{
		Floatingip{
			PortID: portId,
		},
	}
	payloadAttachFIPBytes, err := json.Marshal(data)
	if err != nil {
		logrus.Fatal("Hmm..")
	}
	body := bytes.NewReader(payloadAttachFIPBytes)

	var baseURL = "http://iaas.tcc1.cloud.toastoven.net/network/v2.0/floatingips/" + floatingIpId
	req, err := http.NewRequest("PUT", baseURL, body)
	if err != nil {
		logrus.Fatal("Request Failed: " + baseURL)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", tokenInfo.TokenID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatal("Request Failed.")
	}

	defer resp.Body.Close()
	allRespBytes, _ := ioutil.ReadAll(resp.Body)
	allRespMap := make(map[string]interface{})
	if err := json.Unmarshal(allRespBytes, &allRespMap); err != nil {
		logrus.Fatal(err)
	}

	responseAttachInfo, err := prettyPrintJSON(allRespBytes)
	if err != nil {
		logrus.Fatal("Failed to change to json format.")
	}

	fmt.Printf("%s", responseAttachInfo)
}

func prettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
