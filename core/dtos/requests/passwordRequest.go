package requests

type PasswordRequest struct {
	Password string        `json:"password"`
	Rules    []RuleRequest `json:"rules"`
}
