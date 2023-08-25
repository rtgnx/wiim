package res

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed index.html
var FS embed.FS

func GetFileSystem() http.FileSystem {
	log.Print("using embed mode")
	fsys, err := fs.Sub(FS, ".")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
