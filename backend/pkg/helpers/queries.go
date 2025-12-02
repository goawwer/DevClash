package helpers

import (
	"net/http"
)

type FilterParameters struct {
	Keys   []string
	Values []string
}

func GetQueryWithFilterParameters(r *http.Request) FilterParameters {
	var params FilterParameters

	baseQuery := r.URL.Query()

	for k, v := range baseQuery {
		if len(v) > 0 {
			params.Keys = append(params.Keys, k)
			params.Values = append(params.Values, v[0])
		}
	}

	return params
}
