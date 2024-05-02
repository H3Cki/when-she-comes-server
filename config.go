package wscsrv

type Config struct {
	NetworkCfg NetworkConfig

	MouseSens  float64
	ScrollSens float64
	Actions    []Action
}

type NetworkConfig struct {
	WebRTCMaxConnections   int
	WebRTCWhitelistedNames []string
	WebRTCSDPHubEnabled    bool
	WebRTCSDPHubport       int
	Port                   int
}

type Action struct {
	Name    string
	Command string
}
