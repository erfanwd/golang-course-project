package helpers

import "github.com/erfanwd/golang-course-project/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"result_code"`
	ValidationErrors *[]validations.ValidationError `json:"validation_errors"`
	Error            any                            `json:"error"`
}

func GenerateBaseHttpResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result: result,
		Success: success ,
		ResultCode: resultCode,
	}
}

func GenerateBaseHttpResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse{
	return &BaseHttpResponse{
		Result: result,
		Success: success ,
		Error: err.Error(),
	}
}

func GenerateBaseHttpResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse{
	return &BaseHttpResponse{
		Result: result,
		Success: success ,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode int, err any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}
