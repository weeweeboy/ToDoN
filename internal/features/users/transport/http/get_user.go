package users_trnsport_http

import (
	core_logger "ToDoN/internal/core/logger"
	core_http_response "ToDoN/internal/core/transport/http/response"
	core_http_utils "ToDoN/internal/core/transport/http/utils"
	"net/http"
)

type GetUserResponse UserDTOResponse

func (u *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "user_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user_id path param")
		return
	}

	user, err := u.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
