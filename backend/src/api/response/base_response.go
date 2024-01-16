package response

import (
	"github.com/Stefan923/go-estate-market/api/validator"
)

type BaseHttpResponse struct {
	Result           any  `json:"result"`
	Success          bool `json:"success"`
	ValidationErrors *[]validator.ValidationError
	Error            any `json:"error"`
}

func GenerateResponse(result any, success bool) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Result:  result,
	}
}

func GenerateResponseWithError(result any, success bool, error error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success: success,
		Error:   error.Error(),
	}

}

func GenerateResponseWithAnyError(result any, success bool, error any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success: success,
		Error:   error,
	}
}

func GenerateResponseWithValidationError(result any, success bool, error error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:          success,
		ValidationErrors: validator.ParseValidationErrors(error),
	}
}
