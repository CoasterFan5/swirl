package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

type RouteData struct {
	path string
	data []byte
}

var routeList []RouteData

func registerRoute(path string, data []byte) {
	routeList = append(routeList, RouteData{
		path: path,
		data: data,
	})
}

func discoverChildren(basePath string, addPath string) {
	concatPath := path.Join(basePath, addPath)
	entries, _ := os.ReadDir(concatPath)
	for _, item := range entries {
		if item.IsDir() {
			dirName := item.Name()
			discoverChildren(basePath, path.Join(addPath, dirName))
		} else {
			filePath := path.Join(concatPath, item.Name())
			if item.Name() == "+page.html" {
				// read the file cause we have to expose it.
				fileData, _ := os.ReadFile(filePath)
				registerRoute(path.Join(addPath), fileData)

			}
			fmt.Println(item)
		}
	}
}

func main() {
	// get the routes
	srcPath := path.Join("./src/routes")
	discoverChildren(srcPath, "/")

	mux := http.NewServeMux()

	for _, item := range routeList {
		fmt.Println(item.path)
		if item.path == "/" {
			item.path = "/{$}" // Signals the end of a path

		}
		mux.HandleFunc(item.path, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write(item.data)
		})
	}

	server := http.Server{
		Addr:    ":5177",
		Handler: http.Handler(mux),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(err)
	}
}
