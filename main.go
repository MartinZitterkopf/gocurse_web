package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MartinZitterkopf/gocurse_web/internal/curse"
	"github.com/MartinZitterkopf/gocurse_web/internal/enrollment"
	"github.com/MartinZitterkopf/gocurse_web/internal/user"
	"github.com/MartinZitterkopf/gocurse_web/pkg/bootstrap"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLogger()

	instanceDB, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepo(l, instanceDB)
	userService := user.NewService(l, userRepo)
	userEndpoint := user.MakeEndpoints(userService)

	curseRepo := curse.NewRepo(l, instanceDB)
	curseService := curse.NewService(l, curseRepo)
	curseEndpoint := curse.MakeEndpoints(curseService)

	enrollmentRepo := enrollment.NewRepo(l, instanceDB)
	enrollmentService := enrollment.NewService(l, userService, curseService, enrollmentRepo)
	enrollmentEndpoint := enrollment.MakeEndpoints(enrollmentService)

	router.HandleFunc("/users", userEndpoint.Create).Methods("POST")
	router.HandleFunc("/users", userEndpoint.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEndpoint.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEndpoint.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEndpoint.Delete).Methods("DELETE")

	router.HandleFunc("/curses", curseEndpoint.Create).Methods("POST")
	router.HandleFunc("/curses", curseEndpoint.GetAll).Methods("GET")
	router.HandleFunc("/curses/{id}", curseEndpoint.GetByID).Methods("GET")
	router.HandleFunc("/curses/{id}", curseEndpoint.Update).Methods("PATCH")
	router.HandleFunc("/curses/{id}", curseEndpoint.Delete).Methods("DELETE")

	router.HandleFunc("/enrollments", enrollmentEndpoint.Create).Methods("POST")

	srv := &http.Server{
		// Handler:      http.TimeoutHandler(router, 5*time.Second, "Timeout"), // http.TimeoutHandler() se utiliza para forzar una respues en el tiempo previsto por nosotros, para que no se quede esperando
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	// MANERA DE COMUNICARNOS POR MEDIO DEL PAQUETE ESTANDAR HTTP/NET
	// port := ":3333"												// para http
	// http.HandleFunc("/users", getUsersHttp)						// para http
	// http.HandleFunc("/courses", getCoursesHttp)					// para http
	// err := http.ListenAndServe(port, nil)						// para http
	// if err != nil {												// para http
	// 	fmt.Println(err)											// para http
	// }															// para http
	// -------------------------------------------------------------------------
}

// MANERA DE COMUNICARNOS POR MEDIO DEL PAQUETE ESTANDAR HTTP/NET
// func getUsersHttp(w http.ResponseWriter, r *http.Request) {		// para http
// 	fmt.Println("got /users")										// para http
// 	io.WriteString(w, "This is my user endpoint\n")					// para http
// }																// para http

// func getCoursesHttp(w http.ResponseWriter, r *http.Request) {	// para http
// 	fmt.Println("got /courses")										// para http
// 	io.WriteString(w, "This is my course endpoint\n")				// para http
// }																// para http
// -----------------------------------------------------------------------------

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUsers")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getCourses")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
