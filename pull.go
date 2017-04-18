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

	cmd := exec.Command("pull.bat")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	io.WriteString(w, "body")
}
