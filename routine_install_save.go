package main

import (
	"log"
	"time"
	//"fmt"
)


var internalProgress *ProgressInfo = &ProgressInfo{}

func routineInstallSave(config *InstallConfig) {
	outputProgress(10, "Begin deploy...")
	time.Sleep(1 * time.Second)

	outputProgress(30, "step1...")
	time.Sleep(3 * time.Second)

	outputProgress(80, "step2...")
	time.Sleep(2 * time.Second)
	/**
	//step1: verify configs
	if ok := verifyInstallConfig(config); !ok {
		log.Println("Failed to verify install configuration")
	}
	**/

	channelInstallQuit <- 1
}

func verifyInstallConfig(config *InstallConfig) bool {
	verifyItems := []func(config *InstallConfig)bool{
		verifyGeneralConfig,
	}

	for _, verifyFunc := range verifyItems {
		if !verifyFunc(config){
			return false
		}
	}

	return true
}

func verifyGeneralConfig(config *InstallConfig) bool {
	validDebug := map[string]byte{"enabled":1, "disabled":2}
	if _, ok := validDebug[config.General.Debug]; !ok {
		log.Println("Invalid general debug")
		return false
	}

	return true
}

func outputProgress(percentage uint, message string) {
	internalProgress.Percentage = percentage
	internalProgress.Message = message
	channelInstallProgress<- internalProgress
}