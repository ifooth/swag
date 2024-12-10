package gen

type backend struct {
	Type    string `json:"type"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Timeout int    `json:"timeout"`
}

type authConfig struct {
	AppVerifiedRequired  bool `json:"appVerifiedRequired"`
	UserVerifiedRequired bool `json:"userVerifiedRequired"`
}

type bkAPIGwResource struct {
	IsPublic   bool       `json:"isPublic"`
	AuthConfig authConfig `json:"authConfig"`
	Backend    backend    `json:"backend"`
}
