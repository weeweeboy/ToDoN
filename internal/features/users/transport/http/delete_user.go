package users_trnsport_http

import (
	core_logger "ToDoN/internal/core/logger"
	core_http_response "ToDoN/internal/core/transport/http/response"
	core_http_utils "ToDoN/internal/core/transport/http/utils"
	"net/http"
)

func (u *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	pathValue, err := core_http_utils.GetIntPathValue(r, "user_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get path value ")
		return
	}

	err = u.userService.DeleteUser(ctx, pathValue)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.JSONResponse("user deleted", http.StatusOK)
}
