package controller

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"github.com/jx2lee/nhncloud-instctl/pkg/instance"
)

// InstanceListOutput: print instance list output
func InstanceListOutput(region string) {
	instanceList := instance.GetInstanceList(region)

	// print all instance
	writer := tabwriter.NewWriter(os.Stdout, 16, 8, 2, '\t', 0)
	fmt.Fprintln(writer, "Instance ID\tInstance Name\tImage Name\tStatus\tPublic IP\tPrivate Key")
	for _, instance := range instanceList {
		formatting := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s",
			instance.InstanceID,
			instance.InstanceName,
			instance.ImageName,
			instance.Status,
			instance.PublicIP,
			instance.PrivateKey)
		fmt.Fprintln(writer, formatting)
	}
	writer.Flush()
}

// SSHConnect: connect to NHN Cloud Instance
func SSHConnect(instanceName string, region string) {
	instanceList := instance.GetInstanceList(region)

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
		logrus.Fatal("The instance you entered does not exist. Please recheck the list and try to connect.")
	}
	keyPath := instance.GetEnvparser().GetPrivateKeyPath(privateKey[:len(privateKey)-4])
	logrus.Info("ssh", fmt.Sprintf("-i%s", keyPath), fmt.Sprintf("%s@%s", serverUser, floatingIP))

	cmd := exec.Command("ssh", fmt.Sprintf("-i%s", keyPath), fmt.Sprintf("%s@%s", serverUser, floatingIP))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// StartInstance: start NHN Cloud Instance
func StartInstance(instanceName string, region string) {
	serverID := getServerID(instanceName, region)
	logrus.Info("Instance UUID: ", serverID)

	requestStatusCode := instance.PostInstanceStatus("start", serverID, region)

	if requestStatusCode == 202 {
		logrus.Info("Instance startup succeeded.")
	} else if requestStatusCode == 409 {
		logrus.Fatal("Cannot start instance ", instanceName, " while it is in vm_state active.")
	} else {
		logrus.Fatal("Failed to start instance.")
	}
}

// PauseInstance: pause NHN Cloud Instance
func PauseInstance(instanceName string, region string) {
	serverID := getServerID(instanceName, region)
	logrus.Info("Instance UUID: ", serverID)

	requestStatusCode := instance.PostInstanceStatus("stop", serverID, region)

	if requestStatusCode == 202 {
		logrus.Info("Instance stoping succeeded.")
	} else if requestStatusCode == 409 {
		logrus.Fatal("Cannot start instance ", instanceName, " while it is in vm_state stoped.")
	} else {
		logrus.Fatal("Failed to stop instance.")
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
		logrus.Fatal("Windows server is not supported. Please connect to another instance.")
	}
	return serverUser
}

func setFloatingIP(publicIP string) string {
	if publicIP == "None" {
		logrus.Fatal("You cannot access the instance. After creating FloatingIP, add it to the instance and try.")
	}
	return publicIP
}

func getServerID(instanceName string, region string) string {
	exitedInstanceList := instance.GetInstanceList(region)

	var serverID string
	var isInstance = false
	for _, instance := range exitedInstanceList {
		if instance.InstanceName == instanceName {
			serverID = instance.InstanceID
			isInstance = true
		}
	}

	if isInstance == false {
		logrus.Fatal("The instance you entered does not exist. Please recheck the list and try to connect.")
	}

	return serverID
}
