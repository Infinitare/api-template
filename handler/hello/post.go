package hello

import (
	"encoding/json"
	"github.com/Infinitare/types-template/console"
	"github.com/Infinitare/types-template/responses"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {

	var req postRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		console.ErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	responses.SendJson(postResponse{
		World: req.Hello,
	}, http.StatusCreated, w, r)

}

type postRequest struct {
	Hello string `json:"hello"`
}

type postResponse struct {
	World string `json:"wold"`
}
