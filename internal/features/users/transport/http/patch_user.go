package users_trnsport_http

import (
	"ToDoN/internal/core/domain"
	core_logger "ToDoN/internal/core/logger"
	core_http_request "ToDoN/internal/core/transport/http/request"
	core_http_response "ToDoN/internal/core/transport/http/response"
	core_http_types "ToDoN/internal/core/transport/http/types"
	core_http_utils "ToDoN/internal/core/transport/http/utils"
	"fmt"
	"net/http"
	"strings"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (u *PatchUserRequest) Validate() error {
	if u.FullName.Set {
		if u.FullName.Value == nil {
			return fmt.Errorf("full_name is required")
		}
		fullNameLen := len([]rune(*u.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("full_name length must be between 3 and 100")
		}
	}

	if u.PhoneNumber.Set {
		if u.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*u.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("phone_number length must be between 10 and 15")
			}

			if !strings.HasPrefix(*u.PhoneNumber.Value, "+") {
				return fmt.Errorf("phone_number must start with '+'")
			}
		}
	}
	return nil
}

type PatchUserResponse UserDTOResponse

func (u *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "user_id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user_id path value")
		return
	}

	var request PatchUserRequest
	err = core_http_request.DecodeAndValidateRequest(r, &request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	userPatch := userPatchFromRequest(request)
	userDomain, err := u.userService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user in service")
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)

	return
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
