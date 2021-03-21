package instance

import (
	"log"
)

// ListInfo struct
type ListInfo struct {
	InstanceID   string
	InstanceName string
	ImageName    string
	Status       string
	PublicIP     string
	PrivateKey   string
}

// GetInstanceList return ListInfo struct
func GetInstanceList(region string) []ListInfo {
	// get tenantID
	tenantID, _, _, err := GetEnvparser().GetPasswordCredentials()
	if err != nil {
		log.Fatal("Unable Get TennantID.")
	}

	// get token
	tokenInfo, err := RequestToken()
	if err != nil {
		log.Fatal("Failed Get Token.")
	}

	// get list for instance
	responseAPI, err := RequestInstanceDetails(tokenInfo, tenantID, region)
	if err != nil {
		log.Fatal("Failed Request")
	}
	instanceDetails := responseAPI["servers"]
	list, err := GetInstanceDetails(instanceDetails)

	return list
}

// GetInstanceDetails return struct for instance details
func GetInstanceDetails(servers interface{}) ([]ListInfo, error) {
	var InstanceDetailList []ListInfo
	for _, server := range servers.([]interface{}) {
		// Check task_state to pass if instance is being created
		if server.(map[string]interface{})["OS-EXT-STS:task_state"] != nil {
			continue
		}
		var InstanceInfo = ListInfo{
			InstanceID:   server.(map[string]interface{})["id"].(string),
			InstanceName: server.(map[string]interface{})["name"].(string),
			ImageName:    server.(map[string]interface{})["metadata"].(map[string]interface{})["description"].(string),
			Status:       server.(map[string]interface{})["status"].(string),
			PublicIP:     "None",
			PrivateKey:   server.(map[string]interface{})["key_name"].(string) + ".pem"}

		if len(server.(map[string]interface{})["addresses"].(map[string]interface{})["Default Network"].([]interface{})) == 2 {
			InstanceInfo.PublicIP = server.(map[string]interface{})["addresses"].(map[string]interface{})["Default Network"].([]interface{})[1].(map[string]interface{})["addr"].(string)
		}
		InstanceDetailList = append(InstanceDetailList, InstanceInfo)
	}
	return InstanceDetailList, nil
}
