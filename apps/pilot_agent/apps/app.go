package apps

import (
	"bytes"
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	handler *echo.Echo
	Logs    echo.Logger
)

func GetEcho() *echo.Echo {
	return handler
}

func init() {
	handler = echo.New()

	Logs = handler.Logger
	Logs.SetHeader("[${time_rfc3339}][${short_file}:${line}][${level}]")
}

// StatusOK 기본 Json 형식 리턴
func GetHttpOk(arg ...interface{}) (int, string) {
	/*
		http.StatusCreated              = 201 // RFC 7231, 6.3.2
		http.StatusAccepted             = 202 // RFC 7231, 6.3.3
		http.StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
		http.StatusNoContent            = 204 // RFC 7231, 6.3.5
		http.StatusResetContent         = 205 // RFC 7231, 6.3.6
		http.StatusPartialContent       = 206 // RFC 7233, 4.1
		http.StatusMultiStatus          = 207 // RFC 4918, 11.1
		http.StatusAlreadyReported      = 208 // RFC 5842, 7.1
		http.StatusIMUsed               = 226 // RFC 3229, 10.4.1
	*/

	var code int = http.StatusOK
	if len(arg) > 0 {
		switch arg[0].(type) {
		case int:
			code = arg[0].(int)
		} // end of switch
	} // end of if

	return code, "{}" // 200
}

// HTTP 오류 처리 Json 리턴
func GetHttpError(code int, message string) (int, string) {
	var status string

	switch code {
	// 1xx: informational
	case
		http.StatusContinue,           // 100
		http.StatusSwitchingProtocols, // 101
		http.StatusProcessing,         // 102
		http.StatusEarlyHints:         // 103
		status = http.StatusText(code)

	// 3xx: Redirection
	case
		http.StatusMultipleChoices,   // 300
		http.StatusMovedPermanently,  // 301
		http.StatusFound,             // 302
		http.StatusSeeOther,          // 303
		http.StatusNotModified,       // 304
		http.StatusUseProxy,          // 305
		http.StatusTemporaryRedirect, // 307
		http.StatusPermanentRedirect: // 308
		status = http.StatusText(code)

	// 4xx: Client Error
	case
		http.StatusBadRequest,                   // 400
		http.StatusUnauthorized,                 // 401
		http.StatusPaymentRequired,              // 402
		http.StatusForbidden,                    // 403
		http.StatusNotFound,                     // 404
		http.StatusMethodNotAllowed,             // 405
		http.StatusNotAcceptable,                // 406
		http.StatusProxyAuthRequired,            // 407
		http.StatusRequestTimeout,               // 408
		http.StatusConflict,                     // 409
		http.StatusGone,                         // 410
		http.StatusLengthRequired,               // 411
		http.StatusPreconditionFailed,           // 412
		http.StatusRequestEntityTooLarge,        // 413
		http.StatusRequestURITooLong,            // 414
		http.StatusUnsupportedMediaType,         // 415
		http.StatusRequestedRangeNotSatisfiable, // 416
		http.StatusExpectationFailed,            // 417
		http.StatusTeapot,                       // 418
		http.StatusMisdirectedRequest,           // 421
		http.StatusUnprocessableEntity,          // 422
		http.StatusLocked,                       // 423
		http.StatusFailedDependency,             // 424
		http.StatusTooEarly,                     // 425
		http.StatusUpgradeRequired,              // 426
		http.StatusPreconditionRequired,         // 428
		http.StatusTooManyRequests,              // 429
		http.StatusRequestHeaderFieldsTooLarge,  // 431
		http.StatusUnavailableForLegalReasons:   // 451
		status = http.StatusText(code)

	// 5xx: Server Error
	case
		http.StatusInternalServerError,           // 500
		http.StatusNotImplemented,                // 501
		http.StatusBadGateway,                    // 502
		http.StatusServiceUnavailable,            // 503
		http.StatusGatewayTimeout,                // 504
		http.StatusHTTPVersionNotSupported,       // 505
		http.StatusVariantAlsoNegotiates,         // 506
		http.StatusInsufficientStorage,           // 507
		http.StatusLoopDetected,                  // 508
		http.StatusNotExtended,                   // 510
		http.StatusNetworkAuthenticationRequired: // 511
		status = http.StatusText(code)

	default:
		status = fmt.Sprintf("code %d", code)
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(message)

	data := fmt.Sprintf(`{"status": "%s", "message": %s}`, status, buffer.String())

	return code, data
}
