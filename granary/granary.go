package granary

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/novliang/yh_user/utils"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var DbConfigurator *viper.Viper
var Db = map[string]*gorm.DB{}

func Load() error {
	err := LoadConfig()
	if err != nil {
		return err
	}

	dbConf := DbConfigurator.GetStringMap("db")

	for n, v := range dbConf {
		vMap := v.(map[string]interface{})
		var err error
		db, err := gorm.Open("mysql", vMap["username"].(string)+":"+vMap["password"].(string)+"@tcp("+vMap["host"].(string)+")/"+vMap["db"].(string)+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			return err
		}
		Db[n] = db
	}

	return nil
}

func LoadConfig() (error) {

	w, err := os.Getwd()
	if err != nil {
		return errors.New("Can't get path wd err")
	}
	a, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return errors.New("Can't get path wd err")
	}

	var configFile = "db.toml"
	appConfigPath := filepath.Join(w, "config", configFile)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(a, "config", configFile)
		if !utils.FileExists(appConfigPath) {
			return errors.New("Can't get db config file err")
		}
	}
	DbConfigurator = viper.New()

	DbConfigurator.SetConfigName("db")
	DbConfigurator.AddConfigPath(strings.TrimRight(appConfigPath, configFile))

	err = DbConfigurator.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("Can't get db config file err")
		} else {
			return err
		}
	}
	return nil;
}
