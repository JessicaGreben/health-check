package types

// Results contains the results of all the health checks.
type Results struct {
	Deploys []DeployResults
	Pods    []PodResults
}

// DeployResults has the results of the health checks for deploy.
type DeployResults struct {
	Name       string
	Namespace  string
	Labels     BaseResults
	Containers []ContainerResults
}

// PodResults has the results of the health checks for pods.
type PodResults struct {
	Name      string
	Namespace string
	HPA       bool
	Naked     bool
}

// ContainerResults has the results of the health checks for containers.
type ContainerResults struct {
	Name      string
	Requests  bool
	Limits    bool
	Live      bool
	Ready     bool
	HostPorts bool
	Tag       bool
}

type BaseResults struct {
	Passed bool
	ErrMsg string
}
