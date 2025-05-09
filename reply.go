package client

import (
	"fmt"
	"github.com/samber/lo"
	"go.uber.org/multierr"
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

func (r *Reply) Error() error {
	var err error
	for _, result := range r.Results {
		if result.Ok || !result.Error.Valid {
			continue
		}

		s := result.Label.ValueOrZero()
		if s == "" {
			s = result.StoreName
		}
		err = multierr.Append(err, fmt.Errorf("%s: %s", s, result.Error.String))
	}

	return err
}

func (r *Reply) ErrorStrings() []string {
	errs := make([]string, 0, len(r.Results))
	for _, result := range r.Results {
		if result.Ok || !result.Error.Valid {
			continue
		}

		errs = append(errs, fmt.Sprintf("%s: %s", result.StoreName, result.Error.String))
	}

	return lo.Uniq(errs)
}
