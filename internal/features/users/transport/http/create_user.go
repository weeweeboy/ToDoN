package users_trnsport_http

import (
	"ToDoN/internal/core/domain"
	core_logger "ToDoN/internal/core/logger"
	core_http_request "ToDoN/internal/core/transport/http/request"
	core_http_response "ToDoN/internal/core/transport/http/response"
	"net/http"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=12,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("invoke CreateUser handler")
	var req CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(req)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)

}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
