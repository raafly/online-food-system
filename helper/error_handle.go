package helper

type Response struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Err     interface{} `json:"err,omitempty"`
}

func (e *ErrResponse) Error() string {
	return e.Message
}

func (e *Response) Error() string {
	return e.Message
}

func NewSuccess() error {
	return &Response{
		Code: 200,
		Status: true,
		Message: "SUCCESS",
	}
}

func NewNotFoundError() error {
	return &ErrResponse{
		Code:  404,
		Status: false,
		Message: "id not found",
	}
}

func NewInternalServerError() error {
	return &ErrResponse{
		Code:    500,
		Status:  false,
		Message: "INTERNAL SERVER ERROR",
	}
}
