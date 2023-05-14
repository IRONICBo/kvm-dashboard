package agent

// Get Vm Status
// 1. Init an agent
// 2. Start the agent
// 3. Read data from channel
// 4. ... Save to influxDB or other operations

// todo: check the agent status
type AgentInfo struct {
	// vm info
	UUID string

	// ssh info
	User     string
	Password string
	Ip       string
	Port     uint
}

type AgentInterface interface {
	// Start the agent
	Start() error
	// Stop the agent
	Stop() error
	// Restart the agent
	Restart() error
	// Get the agent status
	Status() error
}
