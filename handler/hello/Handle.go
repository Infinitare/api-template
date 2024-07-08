package hello

import (
	"github.com/Infinitare/types-template/metrics"
	"github.com/Infinitare/types-template/requests"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		requests.MustAdmin(metrics.TraceHandler(post))(w, r)
	}

}
