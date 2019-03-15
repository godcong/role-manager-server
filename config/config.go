package config

import (
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"os"
)

// Database ...
type Database struct {
	Prefix   string `toml:"prefix"`
	Type     string `toml:"type"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	Username string `toml:"username"`
	DB       string `toml:"db"`
}

// REST ...
type REST struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// Media ...
type Media struct {
	Upload      string `toml:"upload"`        //上传路径
	Transfer    string `toml:"transfer"`      //转换路径
	M3U8        string `toml:"m3u8"`          //m3u8文件名
	KeyURL      string `toml:"key_url"`       //default url
	KeyDest     string `toml:"key_dest"`      //key 文件输出目录
	KeyFile     string `toml:"key_file"`      //key文件名
	KeyInfoFile string `toml:"key_info_file"` //keyFile文件名
}

// IPFS ...
type IPFS struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// Queue ...
type Queue struct {
	Type     string `toml:"type"`
	HostPort string `toml:"host_port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type Manager struct {
	ManagerName    string `toml:"manager_name"`
	NodeName       string `toml:"node_name"`
	CensorName     string `toml:"censor_name"`
	EnableGRPC     bool   `toml:"enable_grpc"`
	EnableREST     bool   `toml:"enable_rest"`
	REST           REST   `toml:"rest"`
	RequestType    string `toml:"request_type"`
	KeyAddressRule string `toml:"key_address_rule"`
	Host           string `toml:"host"`
}

// Configure ...
type Configure struct {
	Database Database `toml:"database"`
	//Media    Media    `toml:"media"`
	Queue   Queue   `toml:"queue"`
	IPFS    IPFS    `toml:"ipfs"`
	Manager Manager `toml:"manager"`
}

var config *Configure

// Initialize ...
func Initialize(filePath ...string) error {
	if filePath == nil {
		filePath = []string{"config.toml"}
	}
	log.Println(filePath)
	cfg := LoadConfig(filePath[0])

	config = cfg

	return nil
}

// IsExists ...
func IsExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Panicln(err)
	}
	return true
}

// LoadConfig ...
func LoadConfig(filePath string) *Configure {
	cfg := DefaultConfigure()
	openFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Error(err)
		panic(err.Error())
	}
	decoder := toml.NewDecoder(openFile)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Error(err)
		panic(err.Error())
	}
	log.Infof("config: %+v", cfg)
	return cfg
}

// Config ...
func Config() *Configure {
	if config == nil {
		panic("nil config")
	}
	return config
}

// DefaultString ...
func DefaultString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}

func DefaultConfigure() *Configure {
	return &Configure{
		Database: Database{},
		Queue:    Queue{},
		IPFS:     IPFS{},
		Manager: Manager{
			ManagerName: "godcong.grpc.manager",
			NodeName:    "godcong.grpc.node",
			CensorName:  "godcong.grpc.censor",
			EnableGRPC:  true,
			EnableREST:  true,
			REST: REST{
				Enable: true,
				Type:   "",
				Path:   "",
				Port:   ":7780",
			},
			RequestType:    "grpc",
			Host:           "http://localhost:8080",
			KeyAddressRule: "/api/video/:id/auth",
		},
	}
}
