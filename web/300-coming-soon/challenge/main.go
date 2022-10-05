package main

import (
	"bytes"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
)

var (
	//go:embed static/*
	content      embed.FS
	contentFS, _ = fs.Sub(content, "static")
	upgrader     = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	s              Specification
	waiter         WSWaiter
	fifteenFromNow = time.Now().Add(15 * time.Second)
)

const (
	user     = "jim"
	password = "admin"
)

type Specification struct {
	Port     int    `default:"1337"`
	Flag     string `default:"DEFAULT_FLAG"`
	FlagTime time.Time
}

type WSWaiter struct {
	sockets []*websocket.Conn
	lock    sync.Mutex
}

func (w *WSWaiter) Add(ws *websocket.Conn) {
	w.lock.Lock()
	defer w.lock.Unlock()

	w.sockets = append(w.sockets, ws)
}

func (w *WSWaiter) Run() {
	time.Sleep(time.Until(time.Time(s.FlagTime)))
	log.Println("it's time")

	for {
		w.lock.Lock()
		for _, ws := range w.sockets {
			flag := fmt.Sprintf("BSidesPDX{%s}", s.Flag)
			err := ws.WriteMessage(1, []byte(flag))
			if err != nil {
				log.Println(err)
			}
		}
		w.lock.Unlock()

		time.Sleep(1 * time.Minute)
	}
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

func staticFunc(m *mux.Router) {
	log.Println("Static file contents:")
	fs.WalkDir(contentFS, ".", func(path string, d fs.DirEntry, err error) error {
		log.Println(fmt.Sprintf("\t%s", path))
		return nil
	})
	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(contentFS))))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	login := struct {
		User     string
		Password string
	}{}
	success := map[string]string{"login": "success"}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&login); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if login.User == user && login.Password == password {
		h := w.Header()
		c := map[string]string{"authentication": "successful", "role": "user"}
		cv, _ := json.Marshal(c)
		dgst := GetMD5Hash(string(cv))
		cookieString := fmt.Sprintf("%s;%s", string(cv), dgst)
		h.Add("Cookie", cookieString)
		h.Add("Location", "admin")
		respondWithJSON(w, 302, success)
	} else {
		respondWithError(w, http.StatusForbidden, "login failed")
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func loginRequired(w http.ResponseWriter, r *http.Request) (success bool) {
	auth := struct {
		Authentication string `json:"authentication"`
		Role           string `json:"role"`
	}{}
	cookie := r.Header.Get("Cookie")
	parts := strings.Split(cookie, ";")
	if len(parts) != 2 {
		respondWithError(w, http.StatusBadRequest, "parse error")
		return
	}
	jsonPart := []byte(parts[0])
	calculated := GetMD5Hash(string(jsonPart))
	if parts[1] != calculated {
		respondWithError(w, http.StatusBadRequest, "mismatch")
		return
	}
	err := json.Unmarshal(jsonPart, &auth)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed")
		return
	}
	if auth.Authentication != "successful" {
		respondWithError(w, http.StatusForbidden, "unauthenticated")
		return
	}
	if auth.Role != "admin" {
		respondWithError(w, http.StatusForbidden, "permission denied")
		return
	}
	success = true
	return
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	auth := loginRequired(w, r)
	if auth != true {
		return
	}
	c := map[string]string{"good": "job"}
	respondWithJSON(w, 200, c)
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	waiter.Add(ws)
	err = ws.WriteMessage(1, []byte("FIXME: Jim- fix the selector race"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func sshFP(w http.ResponseWriter, r *http.Request) {
	response := `
		{
		    "hosts": {
		        "prod": [
		            {
		                "public_key": "VGhlIGZsYWcgY291bGQgYmUgaGlkZGVuIGluIHRoaXMgYmFzZTY0IGdhcmJhZ2UsIGJ1dCBpdCBpcyBub3QuIE1heWJlIGxhdGVyPwo=",
		                "algo": "ecdsa-sha2-nistp256"
		            },
		            {
		                "algo": "ssh-rsa",
		                "fp": "SHA256:cSQ6NTZEIXMpTDpEIEU5NiA0OUBENj8gQT0yOj9FSUU=",
		                "port": 22
		            }
		        ]
		    }
		}`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(response))

}

func main() {
	err := envconfig.Process("challenge", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	var blankTime time.Time
	if s.FlagTime == blankTime {
		s.FlagTime = fifteenFromNow
	}

	log.Printf("spec: %+v", s)

	go waiter.Run()

	router := mux.NewRouter()
	router.HandleFunc("/ws", wsEndpoint)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/admin", adminHandler).Methods("GET")
	router.HandleFunc("/.well-known/sshfp", sshFP).Methods("GET")
	staticFunc(router)

	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), router)
}
