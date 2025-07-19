package utils

import "url-shortener/models"

type Cfg models.Config

var globalCfg Cfg


func InitializeGlobalConfig(port int, host string) {
	globalCfg.Localhost = host
	globalCfg.ServicePort = port
}

func GetGlobalConfig() Cfg {
	return globalCfg
}