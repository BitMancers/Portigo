package resource

type SecretSpec struct {
	Provider string `json:"provider"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}
