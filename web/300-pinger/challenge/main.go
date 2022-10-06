package main

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

var (
	//go:embed static/*
	content      embed.FS
	contentFS, _ = fs.Sub(content, "static")
)

func staticFunc(m *mux.Router) {
	log.Println("Static file contents:")
	fs.WalkDir(contentFS, ".", func(path string, d fs.DirEntry, err error) error {
		log.Println(fmt.Sprintf("\t%s", path))
		return nil
	})
	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(contentFS))))
}

func rootStaticHandler(w http.ResponseWriter, r *http.Request) {
	fname := "index.html"
	f, _ := contentFS.Open(fname)
	fi, _ := fs.Stat(contentFS, fname)
	data := make([]byte, fi.Size())
	_, err := f.Read(data)
	if err != nil {
		log.Println("error:", err.Error())
	}
	http.ServeContent(w, r, "index.html", fi.ModTime(), bytes.NewReader(data))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	fmt.Println("host =>", host)
	pingString := fmt.Sprintf("ping -c 3 -w 3 %s", host)
	cmd := exec.Command("/bin/sh", "-c", pingString)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		w.WriteHeader(501)
		return
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Println(err)
		w.WriteHeader(502)
		return
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
		w.WriteHeader(503)
		return
	}

	w.Write(data)
	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	staticFunc(router)
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
