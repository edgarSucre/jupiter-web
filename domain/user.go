package domain

// type Server struct {
// 	LastActivity string            `json:"last_activity"`
// 	Name         string            `json:"name"`
// 	Pending      string            `json:"pending"`
// 	ProgressURL  string            `json:"progress_url"`
// 	Ready        bool              `json:"ready"`
// 	Started      string            `json:"started"`
// 	State        map[string]string `json:"state"`
// 	Stopped      bool              `json:"stopped"`
// 	Url          string            `json:"url"`
// 	UserOptions  map[string]string `json:"user_options"`
// }

type JupyterUser struct {
	Admin        bool              `json:"admin"`
	AuthState    map[string]string `json:"auth_state"`
	Groups       []string          `json:"groups"`
	LastActivity string            `json:"last_activity"`
	Name         string            `json:"name"`
	Pending      string            `json:"string"`
	Roles        []string          `json:"roles"`
	Server       string            `json:"server"`
	//Servers      []Server          `json:"servers"`
}

type (
	User struct {
		Admin    bool
		Email    string
		ID       int64
		Name     string
		Password string
	}

	CreateUserParams struct {
		Admin    bool
		Email    string
		Name     string
		Password string
	}

	UpdateUserParams struct {
		Admin    bool
		ID       int
		Name     string
		Password *string
	}
)
