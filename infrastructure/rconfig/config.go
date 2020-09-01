package rconfig

type DataProvider struct {
	Name             string `json:"name"`
	ConnectionString string `json:"connection_string"`
	User             string `json:"user"`
	Password         string `json:"password"`
	Collection       string `json:"collection"`
}

type Endpoint struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Template struct {
		Provider string `json:"provider"`
		Key      string `json:"key"`
	} `json:"template"`
	Data struct {
		Provider string `json:"provider"`
		Key      string `json:"key"`
		Ttl      int    `json:"ttl"`
	} `json:"data"`
}

type Config struct {
	Endpoints []Endpoint     `json:"endpoints"`
	Providers []DataProvider `json:"providers"`
}
