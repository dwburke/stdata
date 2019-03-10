package api

import (
	"net/http"
	//"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	endpoint_ws "github.com/addictmud/mud/api/ws"
	endpoint_world_room "github.com/dwburke/stdata/api/endpoints"
)

func init() {
	viper.SetDefault("api.enabled", false)
	viper.SetDefault("api.https", false)
	viper.SetDefault("api.listen", ":8085")
	viper.SetDefault("api.ssl-cert", "cert.pem")
	viper.SetDefault("api.ssl-key", "key.pem")
}

func SetupRoutes(r *mux.Router) {
	endpoints.SetupRoutes(r)
}

func Run() {
	if !viper.GetBool("api.enabled") {
		log.Println("api.enabled == false; not starting")
		return
	}

	log.WithFields(log.Fields{
		"api.https":    viper.GetBool("api.https"),
		"api.listen":   viper.GetString("api.listen"),
		"api.ssl-cert": viper.GetString("api.ssl-cert"),
		"api.ssl-key":  viper.GetString("api.ssl-key"),
	}).Info("api: starting")

	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	SetupRoutes(r)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("html"))))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	//srv := &http.Server{
	//Handler:      r,
	//Addr:         viper.GetString("api.listen"),
	//WriteTimeout: 15 * time.Second,
	//ReadTimeout:  15 * time.Second,
	//}

	go func() {
		if viper.GetBool("api.https") == true {
			log.Fatal(
				http.ListenAndServeTLS(
					viper.GetString("api.listen"),
					viper.GetString("api.ssl-cert"),
					viper.GetString("api.ssl-key"),
					handlers.CORS(originsOk, headersOk, methodsOk)(r),
				))
		} else {
			log.Fatal(
				http.ListenAndServe(
					viper.GetString("api.listen"),
					handlers.CORS(originsOk, headersOk, methodsOk)(r),
				),
			)
		}
	}()

	log.Info("api: running")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Info(r.RemoteAddr, " ", r.RequestURI, " - [", r.ContentLength, "]")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
