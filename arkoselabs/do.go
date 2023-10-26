package arkoselabs

import (
	"context"
	"fmt"
	_ "github.com/rafaelcalleja/OpenAIAuth/env"
	"os"
	"sync"
)

var (
	instance *Arkoselabs
	once     sync.Once
)

func init() {
	once.Do(func() {
		instance = &Arkoselabs{}
		instance.login_api = os.Getenv("login_api")
		instance.platformLoginApi = os.Getenv("platform_login_api")
		fmt.Println("api", instance.login_api)
	})
}

func Instance() *Arkoselabs {
	return instance
}

type Arkoselabs struct {
	ctx              context.Context
	login_api        string
	platformLoginApi string
}
