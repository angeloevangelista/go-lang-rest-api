package models

type Cat struct {
	Id    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Owner *Person `json:"owner,omitempty"`
}
