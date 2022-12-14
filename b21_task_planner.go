package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.ListenAndServe(":8888", nil)
}

//go:embed b21_task_planner

var embeddedFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("b21_task_planner"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embeddedFiles, "b21_task_planner")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
