package responses

type PasswordResponse struct {
	Verification bool        `json:"verify"`
	NoMatched    interface{} `json:"noMatch"`
}
