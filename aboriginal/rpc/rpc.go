package rpc

import (
	"errors"
	"github.com/novliang/compass/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var RpcConfigurator *viper.Viper

type Rpc struct {
	*grpc.Server
}

func (r *Rpc) LoadConfig() error {

	w, err := os.Getwd()
	if err != nil {
		return errors.New("Can't get path wd err")
	}
	a, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return errors.New("Can't get path wd err")
	}

	var configFile = "rpc.toml"
	appConfigPath := filepath.Join(w, "config", configFile)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(a, "config", configFile)
		if !utils.FileExists(appConfigPath) {
			return errors.New("Can't get db config file err")
		}
	}

	RpcConfigurator = viper.New()
	RpcConfigurator.SetConfigName("rpc")
	RpcConfigurator.AddConfigPath(strings.TrimRight(appConfigPath, configFile))
	err = RpcConfigurator.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("Can't get db config file err")
		} else {
			return err
		}
	}
	return nil;
}

func (r *Rpc) Start(address string) error {
	listen, err := net.Listen("tcp", address) //监听所有网卡8028端口的TCP连接
	if err != nil {
		return err
	}
	return r.Serve(listen)
}

func Engine() interface{} {
	s := grpc.NewServer()
	return &Rpc{s}
}
