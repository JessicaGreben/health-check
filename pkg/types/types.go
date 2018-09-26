package types

// Results contains the results of all the health checks.
type Results struct {
	Pods []PodResults
}

// PodResults has the results of the health checks for pods.
type PodResults struct {
	Name       string
	Namespace  string
	HPA        bool
	Naked      bool
	Containers []ContainerResults
}

// ContainerResults has the results of the health checks for containers.
type ContainerResults struct {
	Name     string
	Requests bool
	Limits   bool
	Live     bool
	Ready    bool
}
