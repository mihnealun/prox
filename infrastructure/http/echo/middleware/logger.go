package middleware

import (
	"errors"
	"fmt"
	"net/http/httputil"

	"github.com/labstack/echo/v4"

	"github.com/mihnealun/prox/infrastructure/container"
)

// Logger is a dedicated middleware where you can add your logs
func Logger(c container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// let all middleware run (this included the controller also)
			err := next(ctx)
			// pass the error to the HTTPErrorHandler in order to commit the response
			// otherwise response status code would be the default one (200 OK) in case of an error
			if err != nil {
				ctx.Error(err)
			}

			logger, e := c.GetLogger(ctx.Request().Context())
			if e != nil {
				return e
			}

			if c.GetConfig().Debug {
				logDebugInfo(logger, ctx)
			}
			logErrors(logger, err, ctx)

			return err
		}
	}
}

func logErrors(logger container.Logger, err error, ctx echo.Context) {
	if err == nil {
		return
	}

	req := ctx.Request()
	res := ctx.Response()

	reqID := req.Header.Get(echo.HeaderXRequestID)
	if reqID == "" {
		reqID = res.Header().Get(echo.HeaderXRequestID)
	}

	additionalFields := map[string]interface{}{
		"method":               req.Method,
		"host":                 req.Host,
		"uri":                  req.RequestURI,
		"id":                   reqID,
		"response_status_code": res.Status,
	}

	var loggedErr interface{} = err
	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		loggedErr = echoErr.Message
		if echoErr.Internal != nil {
			loggedErr = fmt.Sprintf("%v, %v", echoErr.Message, echoErr.Internal)
		}
	}

	logger.WithFields(additionalFields).Error(loggedErr)
}

func logDebugInfo(logger container.Logger, ctx echo.Context) {
	req := ctx.Request()
	rsp := ctx.Response()

	rq, _ := httputil.DumpRequest(req, true)
	dumpedReq := string(rq)

	additionalFields := map[string]interface{}{
		"request":              dumpedReq,
		"response_status_code": rsp.Status,
	}

	logger.WithFields(additionalFields).Debug(fmt.Sprintf("incoming HTTP request from %s(%s)", req.Referer(), req.RemoteAddr))
}
