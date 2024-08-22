package app

type RequestData struct {
	Project string `json:"project"`
	Key     string `json:"key"`
}


type PutRequestData struct {
	Project string `json:"project"`
	Key     string `json:"key"`
	Value   string `json:"value,omitempty"` // "omitempty" allows this field to be optional
}
