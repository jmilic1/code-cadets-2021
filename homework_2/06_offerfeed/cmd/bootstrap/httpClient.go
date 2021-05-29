package bootstrap

import (
	stdhttp "net/http"
	"time"
)

func HttpClient() *stdhttp.Client {
	return &stdhttp.Client{Timeout: time.Second * 5}
}
