package controller

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jx2lee/nhncloud-instctl/pkg/instance"
)

// SSHConnect connect to NHN Cloud Instance
func SSHConnect(instanceName string) {
	instanceList := instance.GetInstanceList()

	var serverUser, floatingIP, privateKey string
	var isInstance bool
	for _, instance := range instanceList {
		if instance.InstanceName == instanceName {
			serverUser = setServerUser(instance.ImageName)
			floatingIP = setFloatingIP(instance.PublicIP)
			privateKey = instance.PrivateKey
			isInstance = true
		}
	}
	if isInstance == false {
		log.Fatal("The instance you entered does not exist. Please recheck the list and try to connect.")
	}
	keyPath := instance.GetEnvparser().GetPrivateKeyPath(privateKey[:len(privateKey)-4])
	log.Println("ssh", fmt.Sprintf("-i%s", keyPath), fmt.Sprintf("%s@%s", serverUser, floatingIP))

	cmd := exec.Command("ssh", fmt.Sprintf("-i%s", keyPath), fmt.Sprintf("%s@%s", serverUser, floatingIP))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// StartInstance start NHN Cloud Instance
func StartInstance(instanceName string) {
	serverID := getServerID(instanceName)
	log.Println("Instance UUID: ", serverID)
	requestStatusCode := instance.PostInstanceStatus("start", serverID)

	if requestStatusCode == 202 {
		log.Println("Instance startup succeeded.")
	} else if requestStatusCode == 409 {
		log.Fatal("Cannot start instance ", instanceName, " while it is in vm_state active.")
	} else {
		log.Fatal("Instance stoping failed.")
	}
}

// PauseInstance pause NHN Cloud Instance
func PauseInstance(instanceName string) {
	serverID := getServerID(instanceName)
	log.Println("Instance UUID: ", serverID)
	requestStatusCode := instance.PostInstanceStatus("stop", serverID)

	if requestStatusCode == 202 {
		log.Fatal("Instance stoping succeeded.")
	} else if requestStatusCode == 409 {
		log.Fatal("Cannot start instance ", instanceName, " while it is in vm_state stoped.")
	} else {
		log.Fatal("Instance stoping failed.")
	}
}

func setServerUser(imageName string) string {
	var serverUser string
	if strings.Contains(imageName, "CentOS") == true {
		serverUser = "centos"
	} else if strings.Contains(imageName, "Ubuntu") == true {
		serverUser = "ubuntu"
	} else if strings.Contains(imageName, "Debian") == true {
		serverUser = "debian"
	} else {
		log.Fatal("Windows server is not supported. Please connect to another instance.")
	}
	return serverUser
}

func setFloatingIP(publicIP string) string {
	if publicIP == "None" {
		log.Fatal("You cannot access the instance. After creating FloatingIP, add it to the instance and try.")
	}
	return publicIP
}

func getServerID(instanceName string) string {
	exitedInstanceList := instance.GetInstanceList()

	var serverID string
	var isInstance bool = false
	for _, instance := range exitedInstanceList {
		if instance.InstanceName == instanceName {
			serverID = instance.InstanceID
			isInstance = true
		}
	}

	if isInstance == false {
		log.Fatal("The instance you entered does not exist. Please recheck the list and try to connect.")
	}

	return serverID
}
