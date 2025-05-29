package client

import (
	"fmt"
)

type Reply struct {
	RequestId string   `json:"request_id"`
	Results   []Result `json:"results"`
}

func (r *Reply) Reset() *Reply {
	r.RequestId = ""
	r.Results = make([]Result, 0)
	return r
}

func (r *Reply) Errors() []error {
	errs := make([]error, 0, len(r.Results))
	for _, result := range r.Results {
		if result.Ok || !result.Error.Valid {
			continue
		}

		label := result.Label.ValueOrZero()
		if label == "" {
			label = result.StoreName
		}
		errs = append(errs, fmt.Errorf("%s: %s", label, result.Error.String))
	}

	return errs
}

func (r *Reply) ErrorSummary() []string {
	if len(r.Errors()) == 0 {
		return []string{}
	}

	messages := make([]string, len(r.Results))
	for i, e := range r.Errors() {
		messages[i] = e.Error()
	}

	return messages
}
