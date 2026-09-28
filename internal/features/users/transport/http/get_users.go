package users_trnsport_http

import (
	core_logger "ToDoN/internal/core/logger"
	core_http_response "ToDoN/internal/core/transport/http/response"
	core_http_utils "ToDoN/internal/core/transport/http/utils"
	"fmt"
	"net/http"
)

type GetUsersResponse []UserDTOResponse

func (u *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}
	userDomains, err := u.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit query param %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset query param %w", err)
	}
	return limit, offset, nil
}
