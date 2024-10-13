package version

type VersionInfo struct {
	Version       string // Application version
	GitCommit     string // Git commit hash
	BuildTime     string // Time of the build
	GoVersion     string // Go version used
	TargetOS      string // Operating System
	TargetArch    string // System Architecture
	ContainerId   string
	ContainerName string
}

var BuildInfo VersionInfo

func CreateBuildInfo(version string, commit string, buildTime string, goVersion string, targetOS string, targetArch string, containerId string, containerName string) *VersionInfo {
	return &VersionInfo{
		Version:       version,
		GitCommit:     commit,
		BuildTime:     buildTime,
		GoVersion:     goVersion,
		TargetOS:      targetOS,
		TargetArch:    targetArch,
		ContainerId:   containerId,
		ContainerName: containerName,
	}
}
