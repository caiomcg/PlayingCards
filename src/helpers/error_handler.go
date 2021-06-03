package helpers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)


func ErrorHandler(err error, c echo.Context) {
    var (
        code = http.StatusInternalServerError
        key  = "ServerError"
        msg  string
    )

    if he, ok := err.(*httpError); ok {
        code = he.code
        key = he.Key
        msg = he.Message
    } else {
        msg = http.StatusText(code)
    }

    if !c.Response().Committed {
        if c.Request().Method == echo.HEAD {
            err := c.NoContent(code)
            if err != nil {
                c.Logger().Error(err)
            }
        } else {
            err := c.JSON(code, NewHTTPError(code, key, msg))
            if err != nil {
                c.Logger().Error(err)
            }
        }
    }
}

