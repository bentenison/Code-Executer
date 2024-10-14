package executorbus

import "time"

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

type Submission struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	LanguageID      string    `json:"language_id"`
	CodeSnippet     string    `json:"code_snippet"`
	SubmissionTime  time.Time `json:"submission_time"`
	ExecutionStatus string    `json:"execution_status"`
	ResultID        string    `json:"result_id,omitempty"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PerformanceMetric struct {
	ID            int           `json:"id"`
	SubmissionID  string        `json:"submission_id"`
	ExecutionTime time.Duration `json:"execution_time"`
	MemoryUsage   int           `json:"memory_usage"`
	Status        string        `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
type CodeExecutionStat struct {
	ID            string        `json:"id"`
	UserID        string        `json:"user_id"`
	LanguageID    string        `json:"language_id"`
	ExecutionTime time.Duration `json:"execution_time"`
	MemoryUsage   int           `json:"memory_usage"`
	Status        string        `json:"status"`
	ErrorMessage  string        `json:"error_message,omitempty"`
	CodeSnippet   string        `json:"code_snippet"`
	ContainerID   string        `json:"container_id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
