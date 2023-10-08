package main

import (
	"encoding/json"
	"fmt"
	"github.com/NiuStar/OpenAIAuth/auth"

	"os"
)

func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PASSWORD"), os.Getenv("PROXY"))
	auth.IsPlatform = true
	err := auth.Begin()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	// if os.Getenv("PROXY") != "" {
	/*_, err = auth.GetPUID()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}*/
	// }
	// JSON encode auth.GetAuthResult()
	result := auth.GetAuthResult()
	result_json, _ := json.Marshal(result)
	println(string(result_json))
}
