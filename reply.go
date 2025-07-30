package rpclient

import (
	"fmt"

	"gopkg.in/guregu/null.v4"
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

// HasError 回应数据中是否存在错误
func (r *Reply) HasError() bool {
	for _, result := range r.Results {
		if !result.Ok {
			return true
		}
	}
	return false
}

// Errors 错误集合
func (r *Reply) Errors() []error {
	if len(r.Results) == 0 {
		return nil
	}

	errs := make([]error, 0, len(r.Results))
	for _, result := range r.Results {
		if result.Ok {
			continue
		}

		if !result.Error.Valid {
			result.Error = null.StringFrom("rpclient: Unknown error")
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

// ErrorSummary 错误摘要
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
