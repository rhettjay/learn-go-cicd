package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default config .env is unreadable: %v", err)
	}

	port := os.Getenv("PORT")

	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second}

	fmt.Println("server started on ", port)
	srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hmm... Need's more cowbell. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
