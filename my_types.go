package swag

type AuthConfig struct {
	AppVerifiedRequired  string `json:"auth_type"`
	UserVerifiedRequired string `json:"auth_key"`
}

type Backend struct {
	Method  string `json:"method"`
	Url     string `json:"url"`
	Timeout int    `json:"timeout"`
	Type    string `json:"type"`
}

type BKApigatewayResource struct {
	AuthConfig AuthConfig `json:"auth_config"`
	Backend    Backend    `json:"backend"`
	IsPublic   bool       `json:"is_public"`
}

type BKApigateway struct {
	AppVerifiedRequired  string `json:"auth_type"`
	UserVerifiedRequired string `json:"auth_key"`
	IsPublic             bool   `json:"is_public"`
}
