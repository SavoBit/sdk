package main

import (
	"testing"
	"os"
	"strings"
	"os/exec"
)

func init() {

	command := "nohup python app.py &"
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err1 := cmd.Run()
	return
	if err1 != nil {
		panic(err1)

	}
}

func TestCreateCloud(t *testing.T) {

	command := "go run create_cloud.go"
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err1 := cmd.Run()
	if err1 != nil {
		panic(err1)
		t.Fail()
	}
}


//func TestCreateTenant(t *testing.T) {
//
//	command := "go run create_tenant.go"
//	parts := strings.Fields(command)
//	cmd := exec.Command(parts[0], parts[1:]...)
//	cmd.Stderr = os.Stderr
//	cmd.Stdout = os.Stdout
//	err := cmd.Run()
//	if err != nil {
//		panic(err)
//		t.Fail()
//	}
//}
//
//
//func TestCreateHealthMonitor(t *testing.T) {
//
//	command := "go run create_healthmonitor.go"
//	parts := strings.Fields(command)
//	cmd := exec.Command(parts[0], parts[1:]...)
//	cmd.Stderr = os.Stderr
//	cmd.Stdout = os.Stdout
//	err := cmd.Run()
//	if err != nil {
//		panic(err)
//		t.Fail()
//	}
//}
//
//
//func TestCreateProfiles(t *testing.T) {
//
//	command := "go run create_profiles.go"
//	parts := strings.Fields(command)
//	cmd := exec.Command(parts[0], parts[1:]...)
//	cmd.Stderr = os.Stderr
//	cmd.Stdout = os.Stdout
//	err := cmd.Run()
//	if err != nil {
//		panic(err)
//		t.Fail()
//	}
//}
//
//
//func TestCreateVirtualservice(t *testing.T) {
//
//	command := "go run create_virtualservice.go"
//	parts := strings.Fields(command)
//	cmd := exec.Command(parts[0], parts[1:]...)
//	cmd.Stderr = os.Stderr
//	cmd.Stdout = os.Stdout
//	err := cmd.Run()
//	if err != nil {
//		panic(err)
//		t.Fail()
//	}
//}