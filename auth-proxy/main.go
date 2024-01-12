package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

func main() {
	remote, err := url.Parse(os.Getenv("BFF_URL"))
	if err != nil {
		panic(err)
	}
	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			if introspect := retrospectToken(r); !introspect.active {
				w.WriteHeader(introspect.code)
				w.Write([]byte(introspect.err.Error()))
				return
			}
			r.Host = remote.Host
			w.Header().Set("x-user", "heyboonsong")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	http.ListenAndServe(":8080", nil)

}

type Introspect struct {
	active bool
	code   int
	err    error
}

func retrospectToken(r *http.Request) Introspect {
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		return Introspect{
			active: false,
			code:   http.StatusUnauthorized,
			err:    errors.New("Unauthorized"),
		}
	}
	token := strings.TrimSpace(strings.ReplaceAll(accessToken, "Bearer", ""))
	clientID := "admin-cli"
	clientSecret := "k12w9mq5TGLMvWqyWWAMAaVCtczKDhRo"
	realm := "master"
	client := gocloak.NewClient(os.Getenv("KEYCLOAK_URL"))
	ctx := context.Background()
	res, err := client.RetrospectToken(ctx, strings.TrimSpace(token), clientID, clientSecret, realm)
	if err != nil {
		fmt.Println("RetrospectToken: ", err.Error())
		return Introspect{
			active: false,
			code:   http.StatusForbidden,
			err:    errors.New("StatusForbidden"),
		}
	}
	if !*res.Active {
		return Introspect{
			active: false,
			code:   http.StatusUnauthorized,
			err:    errors.New("Unauthorized"),
		}
	}
	return Introspect{
		active: true,
		code:   http.StatusOK,
		err:    nil,
	}
}
