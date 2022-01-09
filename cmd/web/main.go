package main

import (
	"fmt"
	"log"
	"os/exec"

	"omoshiroimg/web"
)

func main() {

	cmd := exec.Command("ls tmp", "-la")
	var result, err = cmd.Output()
	fmt.Printf("%s\n", result)
	fmt.Println("err : ", err)

	srv := web.NewServer()
	log.Fatal(srv.ListenAndServe())

	return
}
