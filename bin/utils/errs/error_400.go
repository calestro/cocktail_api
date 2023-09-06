package errs

import (
	"net/http"

	"github.com/go-chi/render"
)

type Err400 struct {
	Message string
	Error   int
}

func PostError(message string, r *http.Request, w http.ResponseWriter) {

	model := Err400{Message: message, Error: 400}

	res := map[string]*Err400{
		"data": &model,
	}

	render.JSON(w, r, res)
}
