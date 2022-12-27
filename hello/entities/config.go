package entities

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Config struct {
	TemplateDir     string          `json:"templatesDir"`
	AppgroupService string          `json:"appgroupService"`
	ImageService    string          `json:"imageService"`
	Zones           map[string]Zone `json:"zones"`
	DBs             []Db            `json:"dbs"`
	AuthKeyValueMap Auth            `json:"auth"`
}

type Auth struct {
	ClientId           string  `json:"clientId"`
	ClientSecret       string  `json:"clientSecret"`
	AuthNUrl           string  `json:"authnUrl"`
	ServiceUrl         string  `json:"serviceUrl"`
	SessionStoreType   string  `json:"sessionStoreType"`
	SessionStoreParams Session `json:"sessionStoreParams"`
}

type Session struct {
	CookieName     string `json:"cookieName"`
	Secret         string `json:"secret"`
	Duration       int    `json:"duration"`
	ActiveDuration int    `json:"activeDuration"`
}

type Zone struct {
	State        bool   `json:"active"`
	ZeusServer   string `json:"zeusServer"`
	IndunaServer string `json:"indunaServer"`
	RacktablesIP string `json:"racktables"`
}

type Db struct {
	Entity      string `json:"entity"`
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	MaxOpenConn int    `json:"maxOpenConnection"`
	MaxIdleConn int    `json:"maxIdleConnection"`
}

func ReadConfigFile(configFile string) {
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("unable to read ", configFile, err)
	}

	err = json.Unmarshal(jsonFile, &Config)
	if err != nil {
		log.Fatalln("unable to unmarshal configuration ", err)
	}
}
