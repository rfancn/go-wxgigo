package main

import (
	"log"
	"time"
)

func (self *InstallModule) createDelployRoutine(config *InstallConfig) <-chan uint {
	channelQuit := make(chan uint)
	go func() {
		self.updateProgress(10, "Begin deploy...")
		time.Sleep(1 * time.Second)

		self.updateProgress(30, "step1...")
		time.Sleep(3 * time.Second)

		self.updateProgress(80, "step2...")
		time.Sleep(5 * time.Second)
		/**
		//step1: verify configs
		if ok := verifyInstallConfig(config); !ok {
			log.Println("Failed to verify install configuration")
		}
		**/

		//send signal to channelQuit to indicate routine has been finished
		self.updateProgress(100, "Deploy done!")
		channelQuit <- 1
	}()
	//return the channel to the caller
	return channelQuit
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
