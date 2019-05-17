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
	SandboxLogin:    "https://agent-login.sandbox-b.klook.io",
	ProductionLogin: "https://login.klktech.com",
	Sandbox:         "https://agent-api.sandbox-b.klook.io",
	Production:      "https://api.klktech.com",
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
