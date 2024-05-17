package helper

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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

func NewSuccess(data interface{}) error {
	return &Response{
		Code:    200,
		Status:  true,
		Message: "SUCCESS",
		Data: data,
		
	}
}

func NewCreated() error {
	return &Response{
		Code:    201,
		Status:  false,
		Message: "CREATED",
	}
}

func NewNotFoundError(message string) error {
	return &ErrResponse{
		Code:    404,
		Status:  false,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return &ErrResponse{
		Code:    400,
		Status:  false,
		Message: message,
	}
}

func NewInternalServerError() error {
	return &ErrResponse{
		Code:    500,
		Status:  false,
		Message: "INTERNAL SERVER ERROR",
	}
}
