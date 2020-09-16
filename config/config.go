package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ProjectConfig struct {
	Database DataBase `json:"database"`
	Jwt      Jwt      `json:"jwt"`
}

type DataBase struct {
	Mongo Mongo `json:"mongo"`
}

type Mongo struct {
	DatabaseAddress string `json:"database_address"`
	// per second
	ClientExpTime     int64  `json:"client_exp_time"`
	DatabaseName      string `json:"database_name"`
	UserCollection    string `json:"user_collection"`
	ArticleCollection string `json:"article_collection"`
	CommentCollection string `json:"comment_collection"`
}

type Jwt struct {
	// per minute
	TokenLife     int64  `json:"token_life"`
	ContentKey    string `json:"content_key"`
	SigningSecret string `json:"signing_secret"`
}

var Config ProjectConfig

func init() {

	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteStream, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
