package config

import (
	"time"

	"github.com/jessevdk/go-flags"

	"github.com/MultiBanker/broker/src/database/drivers"
)

type Config struct {
	*Database
	*Servers
	*Token
	*NotifyConfig

	Version string
}

type Servers struct {
	HTTP            *HTTP
	VictoriaMetrics *VictoriaServer

	Dbg    bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
	JWTKey string `long:"jwt-key" env:"JWT_KEY" description:"JWT secret key" required:"false" default:"airba-secret"`
}

type HTTP struct {
	Client     *Client
	Admin      *Admin
	HealthPort string `long:"healthcheck-listen" env:"HEALTH_CHECK_LISTEN" description:"Health check Listen Address (format: :9090|127.0.0.1:9090)" required:"false" default:":2026"`
	BasePath   string `long:"base-path" env:"BASE_PATH" description:"base path of the host" required:"false" default:"/broker"`
	FilesDir   string `long:"files-directory" env:"FILES_DIR" description:"Directory where all static files are located" required:"false" default:"/usr/share/broker"`
}

type Database struct {
	DSName string `long:"ds" env:"DATASTORE" description:"DataStore name (format: mongo/null)" required:"false" default:"mongo"`
	DSDB   string `long:"ds-db" env:"DATASTORE_DB" description:"DataStore database name (format: cascade)" required:"false" default:"broker"`
	DSURL  string `long:"ds-url" env:"DATASTORE_URL" description:"DataStore URL (format: mongodb://localhost:27017)" required:"false" default:"mongodb://localhost:27017/"`
}

type VictoriaServer struct {
	ListenAddr string `long:"victoria-metrics-listen" env:"VICTORIA_METRICS_LISTEN" description:"Victoria Metrics Listen Address (format: :9090|127.0.0.1:9090)" required:"false" default:":9090"`
}

type Client struct {
	ListenAddr string `long:"client-clienthttp-listen" env:"CLIENT_HTTP_LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
	IsTesting  bool   `long:"client-testing" env:"CLIENT_APP_TESTING" description:"testing mode"`
}

type Admin struct {
	ListenAddr string `long:"broker-clienthttp-listen" env:"ADMIN_HTTP_LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8090"`
	IsTesting  bool   `long:"broker-testing" env:"ADMIN_APP_TESTING" description:"testing mode"`
}

type Token struct {
	AccessTokenTime  time.Duration `long:"access-token" env:"ACCESS_TOKEN_DURATION_HOURS" description:"Access Token Duration" required:"true" default:"2h"`
	RefreshTokenTime time.Duration `long:"refresh-token" env:"REFRESH_TOKEN_DURATION_MONTH" description:"Refresh Token Duration" required:"true" default:"1h"`
}

type NotifyConfig struct {
	URL  string `long:"kaz_info_url" env:"NOTIFY_URL"`
	User string `long:"kaz_info_user" env:"NOTIFY_USER"`
	Pass string `long:"kaz_info_pass" env:"NOTIFY_PASSWORD"`
}

func (d Database) ToDataStore() drivers.DataStoreConfig {
	return drivers.DataStoreConfig{
		Engine: d.DSName,
		DBName: d.DSDB,
		URL:    d.DSURL,
	}
}

func ParseConfig() (*Config, error) {
	var c Config
	p := flags.NewParser(&c, flags.Default)
	if _, err := p.Parse(); err != nil {
		return nil, err
	}

	return &c, nil
}
