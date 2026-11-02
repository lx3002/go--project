package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){
     fmt.Println("hello world")
     godotenv.Load() 
     portstring:= os.Getenv("PORT")
     if portstring == ""{
        log.Fatal("Port is not found in the environment")
	  }

   router := chi.NewRouter()
   
   router.Use(cors.Handler(cors.Options{
		
		
		AllowedOrigins:  []string{"https://*", "https://*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"*"},
		ExposedHeaders:  []string{"Link"},
		AllowCredentials:false,
		MaxAge: 300,

	}))

   srv := &http.Server{
          Handler: router,
	  Addr: ":" + portstring,	
	}
	log.Printf("Server starting on port %v", portstring)
	err := srv.ListenAndServe()
	if err != nil{
           log.Fatal(err)
	}

   fmt.Println("Port:", portstring)

}


