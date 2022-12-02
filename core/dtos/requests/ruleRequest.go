package requests

type RuleRequest struct {
	Rule  string `json:"rule"`
	Value int64  `json:"value"`
}
