package v1

import "time"

type CPUMetrics struct {
	CPU string `json:"CPU"`
	Metrics
}

type MemoryMetrics struct {
	Memory string `json:"Memory"`
	Metrics
}

type NetworkMetrics struct {
	Network string `json:"Network"`
	Metrics
}

type FileSystemMetrics struct {
	FileSystem string `json:"Filesystem"`
	Metrics
}

type UptimeMetrics struct {
	Uptime string `json:"Uptime"`
	Metrics
}

type Metrics struct {
	Metrics         []Metric  `json:"metrics"`
	LatestTimestamp time.Time `json:"latestTimestamp"`
}

type Metric struct {
	Timestamp time.Time `json:"timestamp"`
	Value     int64     `json:"value"`
}
