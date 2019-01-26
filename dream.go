package dream

import (
	"fmt"
	"log"
	"net/http"
)

func RunApp(args []string){

	flags := ParseFlags(args)

	config := ParseConfig(flags.ConfigPath)

	fmt.Println(config)

	http.HandleFunc("/",handler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err)
	}
}

func handler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	t ,err := response.Write([]byte("success"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t)
}
