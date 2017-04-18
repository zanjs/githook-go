package main

// Config is ...
var Config = struct {
	APPInfo struct {
		Version string `default:"1.0.9"`
		Port    int    `default:"12300"`
		Contact string `default:"anlasheng@gmail.com By Julian -anla"`
	}
	Shell struct {
		File string `default:"pull.bat"`
	}
}{}
