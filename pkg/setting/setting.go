package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

// 编写与配置项保持一致的结构体（App、Server、Database）

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//load app setting
	err = Cfg.Section("app").MapTo(AppSetting) //使用 MapTo 将配置项映射到结构体上
	if err != nil {
		log.Fatalf("Fail to get AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024 //对一些需特殊设置的配置项进行再赋值

	//load server setting
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Fail to get ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	//load db setting
	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Fail to get DatabaseSetting err: %v", err)
	}
}
