package core_http_utils

import (
	core_errors "ToDoN/internal/core/errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path value %w", key, core_errors.ErrInvalidArgument)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("invalid path value = '%s' by key = '%s' %v:%w",
			pathValue,
			key,
			core_errors.ErrInvalidArgument,
			err)
	}

	return val, nil
}
