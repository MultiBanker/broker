package config

import (
	"github.com/jessevdk/go-flags"

	"github.com/MultiBanker/broker/src/database/drivers"
)

type Config struct {
	*Database
	*Servers
	*WorkerConfigs
	*Token
	Version string
}

type Servers struct {
	HTTP            *HTTP
	GRPC            *GRPC
	VictoriaMetrics *VictoriaServer

	Dbg    bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
	JWTKey string `long:"jwt-key" env:"JWT_KEY" description:"JWT secret key" required:"false" default:"airba-secret"`
}

type HTTP struct {
	Client *Client
	Admin  *Admin
}

type Database struct {
	DSName string `long:"ds" env:"DATASTORE" description:"DataStore name (format: mongo/null)" required:"false" default:"mongo"`
	DSDB   string `long:"ds-db" env:"DATASTORE_DB" description:"DataStore database name (format: cascade)" required:"false" default:"broker"`
	DSURL  string `long:"ds-url" env:"DATASTORE_URL" description:"DataStore URL (format: mongodb://localhost:27017)" required:"false" default:"mongodb://localhost:27017/"`
}

type GRPC struct {
	ListenAddr string `long:"grpc-listen" env:"GRPC_LISTEN" description:"Grpc Listen Address (format: :4000|127.0.0.1:4000)" required:"false" default:":4000"`
}

type VictoriaServer struct {
	ListenAddr string `long:"victoria-metrics-listen" env:"VICTORIA_METRICS_LISTEN" description:"Victoria Metrics Listen Address (format: :9090|127.0.0.1:9090)" required:"false" default:":9090"`
}

type Client struct {
	ListenAddr string `long:"client-clienthttp-listen" env:"CLIENT_HTTP_LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
	BasePath   string `long:"client-base-path" env:"CLIENT_BASE_PATH" description:"base path of the host" required:"false" default:"/broker"`
	FilesDir   string `long:"client-files-directory" env:"CLIENT_FILES_DIR" description:"Directory where all static files are located" required:"false" default:"/usr/share/broker"`
	IsTesting  bool   `long:"client-testing" env:"CLIENT_APP_TESTING" description:"testing mode"`
	CertFile   string `long:"client-cert" env:"CLIENT_CERT_FILE" description:"Location of the SSL/TLS cert file" required:"false" default:""`
	KeyFile    string `long:"client-key" env:"CLIENT_KEY_FILE" description:"Location of the SSL/TLS key file" required:"false" default:""`
}

type Admin struct {
	ListenAddr string `long:"admin-clienthttp-listen" env:"ADMIN_HTTP_LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8090"`
	BasePath   string `long:"admin-base-path" env:"ADMIN_BASE_PATH" description:"base path of the host" required:"false" default:"/broker"`
	FilesDir   string `long:"admin-files-directory" env:"ADMIN_FILES_DIR" description:"Directory where all static files are located" required:"false" default:"/usr/share/broker"`
	IsTesting  bool   `long:"admin-testing" env:"ADMIN_APP_TESTING" description:"testing mode"`
	CertFile   string `long:"admin-cert" env:"ADMIN_CERT_FILE" description:"Location of the SSL/TLS cert file" required:"false" default:""`
	KeyFile    string `long:"admin-key" env:"ADMIN_KEY_FILE" description:"Location of the SSL/TLS key file" required:"false" default:""`
}

type WorkerConfigs struct {
}

type Token struct {
	AccessTokenTime  int `long:"access-token" env:"ACCESS_TOKEN_DURATION_HOURS" description:"Access Token Duration" required:"true" default:"2"`
	RefreshTokenTime int `long:"refresh-token" env:"REFRESH_TOKEN_DURATION_MONTH" description:"Refresh Token Duration" required:"true" default:"1"`
}

func (d Database) ToDataStore() drivers.DataStoreConfig {
	return drivers.DataStoreConfig{
		Engine: d.DSName,
		DBName: d.DSDB,
		URL:    d.DSURL,
	}
}

func ParseConfig() (*Config, error) {
	c := &Config{}
	p := flags.NewParser(c, flags.Default)
	if _, err := p.Parse(); err != nil {
		return nil, err
	}

	return c, nil
}
