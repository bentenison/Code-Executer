package executorbus

type Result struct {
}

type ContainerSpec struct {
	ID      string `json:"Id"`
	Names   []string
	Image   string
	ImageID string
	Command string
	Created int64
	Ports   []Port
	State   string
	Status  string
}
type Port struct {
	// Host IP address that the container's port is mapped to
	IP string `json:"IP,omitempty"`
	// Port on the container
	PrivatePort uint16 `json:"PrivatePort"`
	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`
	Type       string `json:"Type"`
}

type Stats struct {
	MemoryStats struct {
		Usage uint64 `json:"usage"`
		Limit uint64 `json:"limit"`
	} `json:"memory_stats"`
	CPUStats struct {
		CPUUsage struct {
			TotalUsage uint64 `json:"total_usage"`
		} `json:"cpu_usage"`
	} `json:"cpu_stats"`
}
