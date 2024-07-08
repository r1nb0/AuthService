package constants

const (
	// Config
	ConfigPath string = "./configs"
	ConfigName string = "config-prod"
	ConfigType string = "yaml"

	// jwtClaims
	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserID"
	NicknameKey            string = "Nickname"
	EmailKey               string = "Email"
	ExpireTimeKey          string = "Exp"
)
