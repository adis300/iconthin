package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func httpError400(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	// TODO: Use custom views to handle such error
	switch status {
	case http.StatusBadRequest:
		fmt.Fprint(w, "Bad Request 400")
	case http.StatusUnauthorized:
		fmt.Fprint(w, "Unauthorized 401")
	case http.StatusForbidden:
		fmt.Fprint(w, "Forbidden 403")
	case http.StatusNotFound:
		fmt.Fprint(w, "Not found 404")
	case http.StatusMethodNotAllowed:
		fmt.Fprint(w, "Method not allowed 405")
	case http.StatusNotAcceptable:
		fmt.Fprint(w, "Not acceptable 406")
	case http.StatusExpectationFailed:
		fmt.Fprint(w, "Expectation failed 417")
	case http.StatusTooManyRequests:
		fmt.Fprint(w, "Too many requests 429")
	case http.StatusUnavailableForLegalReasons:
		fmt.Fprint(w, "Unavailable 451")
	default:
		fmt.Fprint(w, "Unknown error "+strconv.Itoa(status))
	}
}

func httpError400Json(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// TODO: Use custom views to handle such error
	switch status {
	case http.StatusBadRequest:
		fmt.Fprint(w, `{"code":400,"msg":"Bad Request"`)
	case http.StatusUnauthorized:
		fmt.Fprint(w, `{"code":401,"msg":"Unauthorized"`)
	case http.StatusForbidden:
		fmt.Fprint(w, `{"code":403,"msg":"Not found"`)
	case http.StatusNotFound:
		fmt.Fprint(w, `{"code":404,"msg":"Not found"`)
	case http.StatusMethodNotAllowed:
		fmt.Fprint(w, `{"code":405,"msg":"Method not allowed"`)
	case http.StatusNotAcceptable:
		fmt.Fprint(w, `{"code":406,"msg":"Not acceptable"`)
	case http.StatusTooManyRequests:
		fmt.Fprint(w, `{"code":429,"msg":"Too many requests"`)
	case http.StatusUnavailableForLegalReasons:
		fmt.Fprint(w, `{"code":451,"msg":"Unavailable"`)
	default:
		fmt.Fprint(w, `{"code":`+strconv.Itoa(status)+`,"msg":"Unknown error"`)
	}
}

func httpError500(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	// TODO: Use custom views to handle such error
	switch status {
	case http.StatusInternalServerError:
		fmt.Fprint(w, "Internal server error 500")
	case http.StatusNotImplemented:
		fmt.Fprint(w, "Not implemented 501")
	case http.StatusBadGateway:
		fmt.Fprint(w, "Bad gateway 502")
	case http.StatusServiceUnavailable:
		fmt.Fprint(w, "Service unavailable 503")
	case http.StatusGatewayTimeout:
		fmt.Fprint(w, "Gateway timeout 504")
	default:
		fmt.Fprint(w, "Unknown error "+strconv.Itoa(status))
	}
}

func httpError500Json(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// TODO: Use custom views to handle such error
	switch status {
	case http.StatusInternalServerError:
		fmt.Fprint(w, `{"code":500,"msg":"Internal server error"`)
	case http.StatusNotImplemented:
		fmt.Fprint(w, `{"code":501,"msg":"Not implemented"`)
	case http.StatusBadGateway:
		fmt.Fprint(w, `{"code":502,"msg":"Bad gateway"`)
	case http.StatusServiceUnavailable:
		fmt.Fprint(w, `{"code":503,"msg":"Service unavailable"`)
	case http.StatusGatewayTimeout:
		fmt.Fprint(w, `{"code":504,"msg":"Gateway timeout 504"`)
	default:
		fmt.Fprint(w, `{"code":`+strconv.Itoa(status)+`,"msg":"Unknown error"`)
	}
}

func httpErrorRaw(w http.ResponseWriter, errorContent string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, errorContent)
}

func httpErrorCustom(w http.ResponseWriter, status int, errorContent string) {
	w.WriteHeader(status)
	fmt.Fprint(w, errorContent)
}
