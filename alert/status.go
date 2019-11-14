package alert

import (
	"github.com/sahith-narahari/otis/config"
	"net/http"
)

func GetStatus() error {
	localIp := "http://localhost:"
	portNo := config.NewApp.Port
	url := localIp + portNo
	response, err := http.Get(url)
	if err != nil && response == nil {
		return err
	}
	return nil
}
