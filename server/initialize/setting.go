package initialize

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixURL string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	User string
	Password string
	Host string
	Port string
	Name string
	TablePrefix string
	MaxIdleConns int
	MaxOpenConns int
}
var DatabaseSetting = &Database{}

var cfg *ini.File

// Setup initialize the configuration instance
func SettingSetup() {
	var err error
	cfg, err = ini.Load("configs/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	//mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	//RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
