package sock5

const (
	AuthNone byte = iota
	AuthGssApi
	AuthPassword
	AuthIana
	AuthOther

	AuthPass   = 0x0
	AuthRefuse = 0xff
)

var (
	ServerConfig *serverConfig
)

type (
	AuthUser struct {
		Name     string
		Password string
	}

	serverConfig struct {
		User          *AuthUser
		Port          int
		TcpBufferSize int
	}
)

func InitConfig(name, password string, port, buffsize int) {
	ServerConfig = &serverConfig{
		User: &AuthUser{
			Name:     name,
			Password: password,
		},
		Port:          port,
		TcpBufferSize: buffsize,
	}
}

func init() {
	ServerConfig = &serverConfig{
		User:          nil,
		Port:          1080,
		TcpBufferSize: 2048,
	}
}
