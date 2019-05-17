package easyJson

//easyjson:json
type RequestPattern struct {
	Domain string `json:"domain"`
	Path   string `json:"path"`
}

//easyjson:json
type ResponsePattern struct {
	Location   string `json:"location"`
}
