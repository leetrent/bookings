package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/leetrent/bookings/internal/config"
)

var appConfig *config.AppConfig

func NewHelpers(ac *config.AppConfig) {
	appConfig = ac
}

func ClientError(rw http.ResponseWriter, statusCode int) {
	appConfig.InfoLog.Println("Client error with status of ", statusCode)
	http.Error(rw, http.StatusText(statusCode), statusCode)
}

func ServerError(rw http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLog.Println(errMsg)
	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}