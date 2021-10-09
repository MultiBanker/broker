package config

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/MultiBanker/broker/src/database/drivers"
)

type Config struct {
	*Database
	*Servers
	*Client
	*WorkerConfigs
	*Token
	Version string
}

type Servers struct {
	HTTP       *HTTPServer
	GRPC       *GRPCServer
	Prometheus *PrometheusServer

	Dbg    bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
	JWTKey string `long:"jwt-key" env:"JWT_KEY" description:"JWT secret key" required:"false" default:"airba-secret"`
}

type Database struct {
	DSName string `short:"n" long:"ds" env:"DATASTORE" description:"DataStore name (format: mongo/null)" required:"false" default:"mongo"`
	DSDB   string `short:"d" long:"ds-db" env:"DATASTORE_DB" description:"DataStore database name (format: cascade)" required:"false" default:"broker"`
	DSURL  string `short:"u" long:"ds-url" env:"DATASTORE_URL" description:"DataStore URL (format: mongodb://localhost:27017)" required:"false" default:"mongodb://localhost:27017/"`
}

type HTTPServer struct {
	ListenAddr string `short:"l" long:"listen" env:"LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
	BasePath   string `long:"base-path" env:"BASE_PATH" description:"base path of the host" required:"false" default:"/broker"`
	FilesDir   string `long:"files-directory" env:"FILES_DIR" description:"Directory where all static files are located" required:"false" default:"/usr/share/broker"`
	IsTesting  bool   `long:"testing" env:"APP_TESTING" description:"testing mode"`
	CertFile   string `short:"c" long:"cert" env:"CERT_FILE" description:"Location of the SSL/TLS cert file" required:"false" default:""`
	KeyFile    string `short:"k" long:"key" env:"KEY_FILE" description:"Location of the SSL/TLS key file" required:"false" default:""`
}

type GRPCServer struct {
	ListenAddr string `long:"grpc-listen" env:"GRPC_LISTEN" description:"Grpc Listen Address (format: :4000|127.0.0.1:4000)" required:"false" default:":4000"`
}

type PrometheusServer struct {
	ListenAddr string `long:"prom-listen" env:"PROM_LISTEN" description:"Listen Address (format: :9090|127.0.0.1:9090)" required:"false" default:":9090"`
}

type Client struct {
}

type WorkerConfigs struct {
}

type Token struct {
	AccessToken  int `long:"access-token" env:"ACCESS_TOKEN_DURATION_HOURS" description:"Access Token Duration" required:"true" default:"2"`
	RefreshToken int `long:"refresh-token" env:"REFRESH_TOKEN_DURATION_MONTH" description:"Refresh Token Duration" required:"true" default:"1"`
}


func (d Database) ToDataStore() drivers.DataStoreConfig {
	return drivers.DataStoreConfig{
		Engine: d.DSName,
		DBName: d.DSDB,
		URL:    d.DSURL,
	}
}

func ParseConfig() *Config {
	c := &Config{}
	p := flags.NewParser(c, flags.Default)
	if _, err := p.Parse(); err != nil {
		log.Println("[ERROR] Ошибка парсинга опций:", err)
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	return c
}
