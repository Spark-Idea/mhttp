package server

type StringKVSpec struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type StringTVSpec struct {
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type RequestSpec struct {
	Verb string `yaml:"verb"`
	Path string `yaml:"path"`
}

type ResponseSpec struct {
	Code    int            `yaml:"code"`
	Headers []StringKVSpec `yaml:"headers"`
	Body    StringTVSpec   `yaml:"body"`
}

type RestRequestResponseSpec struct {
	SpecType     string       `yaml:"type"`
	Id           string       `yaml:"id"`
	RequestSpec  RequestSpec  `yaml:"request"`
	ResponseSpec ResponseSpec `yaml:"response"`
}

type SpecBoundle struct {
	Spec []RestRequestResponseSpec `yaml:"spec"`
}
