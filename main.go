package main

import (
	"github.com/gorilla/mux"
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/mqtt"
	v1 "github.com/insektionen/IN-TV/v1"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strings"
)

var clients map[string]bool

func ClientSaver(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if clientName := r.PostFormValue("name"); clientName != "" {
			clients[clientName] = true
		} else {
			clients["other"] = true
		}
		next.ServeHTTP(w, r)
	})
}

func setupConfig() {
	viper.SetConfigName("in_tv")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/insektionen")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("http.listen", ":8081")

	viper.SetDefault("mqtt.broker", "tcp://domain.tld:port")
	viper.SetDefault("mqtt.username", "username")
	viper.SetDefault("mqtt.password", "password")
	viper.SetDefault("mqtt.client_id", "IN-TV")

	viper.SetDefault("paths.slideshow_storage", "/var/lib/insektionen/in_tv/slideshows")
	viper.SetDefault("paths.album_storage", "/var/lib/insektionen/in_tv/albums")

	if err := viper.ReadInConfig(); err != nil {
		_ = viper.SafeWriteConfig()
		log.Fatalln("Could not read config:", err)
	}
}

func main() {
	setupConfig()
	mqtt.Connect()

	err := os.MkdirAll(viper.GetString("paths.slideshow_storage"), 0777)
	if err != nil {
		log.Fatalln("Could not create slideshow storage:", err)
	}
	err = os.MkdirAll(viper.GetString("paths.album_storage"), 0777)
	if err != nil {
		log.Fatalln("Could not create album storage:", err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = api.NotFoundHandler
	r.MethodNotAllowedHandler = api.MethodNotAllowedHandler
	r.Use(api.RecoveryMiddleware)
	r.Use(api.LogMiddleware)
	r.Use(cors.AllowAll().Handler)

	log.Println("Registering middleware and routes")
	v1.RegisterRoutes(r.PathPrefix("/api/v1").Subrouter())

	fh := &frontendHandler{staticPath: "frontend/build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(fh)

	log.Println("Listening on", viper.GetString("http.listen"))
	if err := http.ListenAndServe(viper.GetString("http.listen"), r); err != nil {
		log.Fatalln("Could not listen for HTTP:", err)
	}
}
