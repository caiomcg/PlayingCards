package helpers

type httpError struct {
    code    int
    Key     string `json:"error"`
    Message string `json:"message"`
}

func NewHTTPError(code int, key string, msg string) *httpError {
    return &httpError{
        code:    code,
        Key:     key,
        Message: msg,
    }
}

func (e *httpError) Error() string {
    return e.Key + ": " + e.Message
}
