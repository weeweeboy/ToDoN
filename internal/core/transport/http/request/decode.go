package core_http_request

import (
	core_errors "ToDoN/internal/core/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil { // под капотом вызывает json.Unmarshal(), который у нас "переписан"
		return fmt.Errorf("decode request body %v:%w", err, core_errors.ErrInvalidArgument)
	}

	var err error

	v, ok := dest.(validatable)
	if ok {
		err = v.Validate()
	} else {
		err = validate.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("validate request body %v:%w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}
