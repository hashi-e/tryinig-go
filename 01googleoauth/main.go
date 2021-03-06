package main

import (
	"net/http"
	"log"
  "fmt"
  "context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  oauthapi "google.golang.org/api/oauth2/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/loginr", LoginRHandler)
	log.Println("Server has started")
	http.ListenAndServe(":5001", mux)
}

func LoginHandler (w http.ResponseWriter, r *http.Request) {
  var url = conf.AuthCodeURL("yourStateUUID", oauth2.AccessTypeOffline)
  fmt.Fprintf(w, "Visit here : %s", url)
}

func LoginRHandler(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query()["code"]
  if code == nil ||  len(code) == 0 {
    fmt.Fprint(w,"Invalid Parameter")
  }
  ctx := context.Background()
  tok, err := conf.Exchange(ctx, code[0])
  if err != nil {
  fmt.Fprintf(w,"OAuth Error:%v", err)
  }
  client := conf.Client(ctx, tok)
  svr, err := oauthapi.New(client)
  ui, err := svr.Userinfo.Get().Do()
  if err != nil {
    fmt.Fprintf(w,"OAuth Error:%v", err)
  } else {
    fmt.Fprintf(w, "Your are logined as : %s",  ui.Email)
  }
}

var conf = &oauth2.Config{
  ClientID:     "Your Client ID",
  ClientSecret: "Your Client Secret",
  Scopes:       []string{oauthapi.UserinfoEmailScope},
  Endpoint: google.Endpoint,
  RedirectURL : "http://localhost:5001/loginr",
}
