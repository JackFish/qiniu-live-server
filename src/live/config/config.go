package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiniu/log"
	"os"
)

var (
	App *AppConfig
)

const (
	DEFAULT_READ_TIMEOUT = 60
	DEFAULT_WRITE_TIMEOUT = 60
	DEFAULT_MAX_HEADER_BYTES = 1 << 12 //4KB

	DEFAULT_LOG_FILE = "run.log"
)

type SyncConfig struct {
	Server ServerConfig `json:"server"`
	App    AppConfig    `json:"app"`
	Orm    OrmConfig    `json:"orm"`
}

//server config
type ServerConfig struct {
	LocationPrefix string `json:"location_prefix"`
	ListenHost     string `json:"listen_host"`
	ListenPort     int    `json:"listen_port"`
	ReadTimeout    int    `json:"read_timeout,omitempty"`
	WriteTimeout   int    `json:"write_timeout,omitempty"`
	MaxHeaderBytes int    `json:"max_header_bytes,omitempty"`
}

//app config
type AppConfig struct {
	LiveHub   string `json:"live_hub"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	LogFile   string `json:"log_file,omitempty"`
	LogLevel  string `json:"log_level,omitempty"`
	QLogLevel int
}

//orm config
type OrmConfig struct {
	DriverName  string `json:"driver_name"`
	DataSource  string `json:"data_source"`
	MaxIdleConn int    `json:"max_idle_conn,omitempty"`
	MaxOpenConn int    `json:"max_open_conn,omitempty"`
	DebugMode   bool   `json:"debug_mode,omitempty"`
}

func LoadConfig(confFile string) (cfg *SyncConfig, err error) {
	cfgFh, openErr := os.Open(confFile)
	if openErr != nil {
		err = openErr
		return
	}
	defer cfgFh.Close()
	cfg = &SyncConfig{}
	decoder := json.NewDecoder(cfgFh)
	decodeErr := decoder.Decode(&cfg)
	if decodeErr != nil {
		err = errors.New(fmt.Sprintf("parse config error, %s", decodeErr))
		return
	}

	//check server defaults
	if cfg.Server.ReadTimeout <= 0 {
		cfg.Server.ReadTimeout = DEFAULT_READ_TIMEOUT
	}
	if cfg.Server.WriteTimeout <= 0 {
		cfg.Server.WriteTimeout = DEFAULT_WRITE_TIMEOUT
	}
	if cfg.Server.MaxHeaderBytes <= 0 {
		cfg.Server.MaxHeaderBytes = DEFAULT_MAX_HEADER_BYTES
	}

	if cfg.App.LogFile == "" {
		cfg.App.LogFile = DEFAULT_LOG_FILE
	}

	//check log level
	switch cfg.App.LogLevel {
	case "debug":
		cfg.App.QLogLevel = log.Ldebug
	case "info":
		cfg.App.QLogLevel = log.Linfo
	case "warn":
		cfg.App.QLogLevel = log.Lwarn
	case "error":
		cfg.App.QLogLevel = log.Lerror
	case "panic":
		cfg.App.QLogLevel = log.Lpanic
	case "fatal":
		cfg.App.QLogLevel = log.Lfatal
	default:
		cfg.App.QLogLevel = log.Ldebug
	}

	return
}
