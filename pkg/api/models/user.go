package models

type Server struct {
	LastActivity string            `json:"last_activity"`
	Name         string            `json:"name"`
	Pending      string            `json:"pending"`
	ProgressURL  string            `json:"progress_url"`
	Ready        bool              `json:"ready"`
	Started      string            `json:"started"`
	State        map[string]string `json:"state"`
	Stopped      bool              `json:"stopped"`
	Url          string            `json:"url"`
	UserOptions  map[string]string `json:"user_options"`
}

type User struct {
	Admin        bool              `json:"admin"`
	AuthState    map[string]string `json:"auth_state"`
	Groups       []string          `json:"groups"`
	LastActivity string            `json:"last_activity"`
	Name         string            `json:"name"`
	Pending      string            `json:"string"`
	Roles        []string          `json:"roles"`
	Server       string            `json:"server"`
	Servers      []Server          `json:"servers"`
}

type stateQuery string

const (
	StateQueryActive   stateQuery = "active"
	StateQueryInactive stateQuery = "inactive"
	StateQueryReady    stateQuery = "ready"
)

type UserListQuery struct {
	State                 stateQuery
	Offset                uint
	Limit                 uint
	IncludeStoppedServers bool
}
