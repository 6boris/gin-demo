package config

import (
	"fmt"
	"testing"
)

func TestSetup(t *testing.T) {

	fmt.Println("server:")
	fmt.Println("server.RunMode", ServerSetting.RunMode)
	fmt.Println("server.HttpPort", ServerSetting.HttpPort)
	fmt.Println("server.ReadTimeout", ServerSetting.ReadTimeout)
	fmt.Println("server.WriteTimeout", ServerSetting.WriteTimeout)

}
