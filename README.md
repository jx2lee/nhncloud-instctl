# NHN Cloud Instance CTL

This repository is a cli program that manages nhncloud instances and was developed as Golang.

![](gif/intro.gif)

# Features

The following features are provided by this script:

* Check the list of instances.
* Connect to the instance.
* Start and stop the instance.
* Check the script version.

# How to Install

## Download Binary

```bash
# in macos..
$ wget https://github.com/jx2lee/nhncloud-instctl/releases/download/v0.1/nhncloud-instctl-darwin && \
       chmod +x nhncloud-instctl-darwin && \
       mv nhncloud-instctl-darwin /usr/local/bin/nhncloud-instctl

# in linux(x86_64)..
$ wget https://github.com/jx2lee/nhncloud-instctl/releases/download/v0.1/nhncloud-instctl-Linux-x86-64 && \
       chmod +x nhncloud-instctl-Linux-x86-64 && \
       mv nhncloud-instctl-Linux-x86-64 /usr/local/bin/nhncloud-instctl
```

## Set Configuration

Create configuration file as `/etc/toast-config.yaml` like below.
```yaml
version: v1
spec:
  credentials:
    tenantID: "{your_tenantID}"
    accessID: "{your_nhncloud_ID}"
    accessPassword: "{your_nhncloud_PW}"
  privateKeys:
    - keyName: "{your_privatekey1_name}"
      keyPath: "{yoyr_privatekey1_path}"    
    - keyName: "{your_privatekey2_name}"
      keyPath: "{yoyr_privatekey2_path}"
    ...
    if more..
    ...
```

Configuration file consists of two part.
* spec.credentails.tennantID,accessID and accessPassword : NHN Cloud credentials ([https://docs.toast.com/ko/Compute/Compute/ko/identity-api/](https://docs.toast.com/ko/Compute/Compute/ko/identity-api/))
* spec.privateKeys : Default SSH private key when using nhncloud-instctl connect for a NHN Cloud instance.

# How to Use

## Print the list of instances

The ability to check the list of instances.  
```bash
$ nhncloud-instctl list
```

```bash
❯ nhncloud-instctl list
Instance ID				Instance Name				Image Name				Status		Public IP		Private Key
19d55cbe-86e0-40cb-a501-09f2ff6d1db4	rancher-test-03				CentOS 7.5 (2020.12.22)			SHUTOFF		None			jx2lee.pem
6b26d3fb-9659-43f5-839f-9b9dd235c2e2	rancher-test-02				CentOS 7.5 (2020.12.22)			SHUTOFF		None			jx2lee.pem
b06d397a-fa84-4be2-9d4a-69533c582bd3	rancher-test-01				CentOS 7.5 (2020.12.22)			SHUTOFF		133.186.223.27		jx2lee.pem
db9c01c4-a8ba-45cb-8df6-1af8841566a0	jx2lee-default-w-am34o3vbq5wg-node-0	CentOS 7.5 - Container (2020.01.26)	ACTIVE		133.186.217.62		jx2lee.pem
25ad6d7d-d184-4fb5-bceb-596453714c3a	master002				CentOS 7.5 (2020.12.22)			SHUTOFF		None			jx2lee.pem
50c43dbd-c408-4f09-8c6e-1ddcb13c2643	master003				CentOS 7.5 (2020.12.22)			SHUTOFF		None			jx2lee.pem
afbae9c2-e788-42bc-895c-fa0165570547	master001				CentOS 7.5 (2020.12.22)			SHUTOFF		133.186.213.49		jx2lee.pem
d0c17dd0-09f8-4ae0-a22f-55116280a3ab	jx2lee					CentOS 7.5 (2020.12.22)			ACTIVE		133.186.241.218		jx2lee.pem
```

## Connect to the instance

The ability to connect to the existing instance.  
```bash
$ nhncloud-instctl connect {exiting_instance_name}
```

```bash
❯ nhncloud-instctl connect jx2lee
2021/03/08 14:33:35 ssh -i/Users/nhn/workspace/toast-cloud/jx2lee.pem centos@133.186.241.218
Last login: Mon Mar  8 10:16:15 2021 from 211.178.107.54
[centos@singlenode ~]$
```

## Start & Stop the instance

The ability to start or stop the existing instance.  
```bash
$ nhncloud-instctl start(or stop) {exiting_instance_name}
```

```bash
# start
❯ nhncloud-instctl start master001
2021/03/08 14:34:14 Instance UUID:  afbae9c2-e788-42bc-895c-fa0165570547
2021/03/08 14:34:15 Instance startup succeeded.
# stop
❯ nhncloud-instctl stop master001
2021/03/08 14:34:41 Instance UUID:  afbae9c2-e788-42bc-895c-fa0165570547
2021/03/08 14:34:42 Instance stoping succeeded.
```

## Print Version

Print binary verison.
```bash
$ nhncloud-instctl version
```

# Reference

* [https://github.com/alicek106/go-ec2-ssh-autoconnect](https://github.com/alicek106/go-ec2-ssh-autoconnect)