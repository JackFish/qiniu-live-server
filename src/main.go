package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qiniu/log"
	"live/config"
	"live/module/model"
	"live/server"
	"os"
	"runtime"
)

const (
	VERSION = "1.0.1"
)

func initLog(logLevel int, logFile string) (err error) {
	log.Info("init log")
	log.SetOutputLevel(logLevel)

	var logFp *os.File
	if logFile == "stdout" {
		logFp = os.Stdout
	} else {
		var openErr error
		logFp, openErr = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if openErr != nil {
			err = openErr
			return
		}
	}
	log.SetOutput(logFp)

	return
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var confFile string
	flag.StringVar(&confFile, "c", "", "config file for the service")
	flag.Usage = func() {
		fmt.Println(`
Usage of qlived:
    -c="": config file for the service
version ` + VERSION)
	}
	flag.Parse()

	if confFile == "" {
		fmt.Println("no config file specified")
		os.Exit(1)
	}
	_, statErr := os.Stat(confFile)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			fmt.Println("config file not found")
		} else {
			fmt.Println(statErr)
		}
		os.Exit(1)
	}

	//load config
	cfg, cfgErr := config.LoadConfig(confFile)
	if cfgErr != nil {
		fmt.Println(cfgErr)
		os.Exit(1)
	}

	config.App = &cfg.App

	//init log
	lErr := initLog(cfg.App.QLogLevel, cfg.App.LogFile)
	if lErr != nil {
		fmt.Println("init log error,", lErr)
		os.Exit(1)
	}

	//init orm
	ormErr := model.InitOrm(&cfg.Orm)
	if ormErr != nil {
		fmt.Println(ormErr)
		os.Exit(1)
	}

	//start server
	serv := server.NewServer(&cfg.Server)
	servErr := serv.Listen()
	if servErr != nil {
		fmt.Println(servErr)
		os.Exit(1)
	}
}
