package coreinstruction

type RequestSpecification struct {
	Id                string    `json:"Id"`
	TypeDotFilter     string    `json:"TypeDotFilter"`
	Tags              *[]string `json:"Tags,omitempty"`
	IsRunAll          bool      `json:"IsRunAll"`
	IsContinueOnError bool      `json:"IsContinueOnError"`
	IsGlobal          bool      `json:"IsGlobal"`
}
