package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

// Pull is
func Pull(w http.ResponseWriter, req *http.Request) {

	var str1, str2 string
	str1 = GetCurrentDirectory()

	str2 = GetParentDirectory(str1)

	fmt.Println(str2)

	cmd := exec.Command(Config.Shell.File)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)

		err = fmt.Errorf("%s", err)

		fmt.Println(err.Error())

	}
	Alog(string(out))
	fmt.Println(string(out))

	io.WriteString(w, "body")
}
