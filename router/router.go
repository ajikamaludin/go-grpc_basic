package router

import (
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
)

func IgnoreErr(err error) bool {
	switch {
	case err == nil || err == http.ErrServerClosed || err == cmux.ErrListenerClosed:
		return true
	}

	if opErr, ok := err.(*net.OpError); ok {
		return opErr.Err.Error() == "use of closed network connection"
	}

	return false
}
