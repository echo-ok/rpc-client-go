package rpclient

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
	if len(r.Results) == 0 {
		return nil
	}

	errs := make([]error, 0, len(r.Results))
	for _, result := range r.Results {
		if result.Ok || !result.Error.Valid {
			continue
		}

		var label string
		if result.Label.Valid {
			label = result.Label.String
		} else {
			label = result.StoreName
		}
		errs = append(errs, fmt.Errorf("%s: %s", label, result.Error.String))
	}

	return errs
}

func (r *Reply) ErrorSummary() []string {
	if len(r.Errors()) == 0 {
		return nil
	}

	messages := make([]string, len(r.Results))
	for i, e := range r.Errors() {
		messages[i] = e.Error()
	}

	return messages
}
