package main

import (
	"log"
	"net/http"
)

func main() {

	// ハンドラーを貼るパスの指定です。今回は他のディレクトリもないので、ルートを指定します。
	rootPath := "/"
	// 返却するファイルが配置されている場所のパスです。
	docDir := "./"

	http.HandleFunc(
		rootPath,
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, docDir)
		},
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
	//appengine.Main()
}
