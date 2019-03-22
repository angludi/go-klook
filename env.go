package klook

type Env int8

const (
	_ Env = iota
	// Sandbox : represent login sandbox environment
	SandboxLogin

	// Production : represent login production environment
	ProductionLogin

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)

var typeString = map[Env]string{
	SandboxLogin:    "http://sandbox-login.klktech.com",
	ProductionLogin: "http://login.klktech.com",
	Sandbox:         "http://sandbox-api.klktech.com",
	Production:      "http://api.klktech.com",
}

// implement stringer
func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}