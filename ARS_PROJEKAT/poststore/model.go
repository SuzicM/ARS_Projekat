package poststore

type Config struct {
	Id string `json:"id"`
	Version string            `json:"version"`
	Entries map[string]string `json:"entries"`
}


type ConfigGroup struct {
	id string 
	Version     string   `json:"version"`
	Group []Config `json:"group"`
}
