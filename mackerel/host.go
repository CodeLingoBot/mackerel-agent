package mackerel

import mkr "github.com/mackerelio/mackerel-client-go"

// Host XXX
type Host struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Type             string `json:"type"` // TODO ENUM
	Status           string `json:"status"`
	CustomIdentifier string `json:"customIdentifier"`
}

// HostSpec is host specifications sent Mackerel server per hour
type HostSpec struct {
	Name             string            `json:"name"`
	Meta             mkr.HostMeta      `json:"meta"`
	Interfaces       interface{}       `json:"interfaces"`
	RoleFullnames    []string          `json:"roleFullnames"`
	Checks           []mkr.CheckConfig `json:"checks"`
	DisplayName      string            `json:"displayName,omitempty"`
	CustomIdentifier string            `json:"customIdentifier,omitempty"`
}
