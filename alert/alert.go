package alert

import (
	"fmt"
	"github.com/sahith-narahari/otis/config"
)

func SendAlert() error {
	fmt.Println("this is token",config.NewApp.Token)
	fmt.Println("this is port",config.NewApp.Port)
	return nil
}