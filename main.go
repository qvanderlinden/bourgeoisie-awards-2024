package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	myClerk "github.com/bourgeoisie-awards-2024/clerk"
	"github.com/bourgeoisie-awards-2024/handlers"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: couldn't load env file")
	}

	clerkClient, err := clerk.NewClient(os.Getenv("CLERK_SECRET_KEY"))
	if err != nil {
		fmt.Println("Could not create the Clerk client")
		panic(err)
	}

	clerkFrontendConfig := myClerk.FrontendConfig{
		PublishableKey: os.Getenv("CLERK_PUBLISHABLE_KEY"),
		FrontendAPI:    os.Getenv("CLERK_FRONTEND_API"),
		Version:        os.Getenv("CLERK_VERSION"),
	}

	pool, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	landingPageHandler := handlers.NewLandingPageHandler(&clerkFrontendConfig)
	welcomePageHandler := handlers.NewWelcomePageHandler(&clerkFrontendConfig)
	votesPageHandler := handlers.NewVotesPageHandler(&clerkFrontendConfig, pool)
	thankYouPageHandler := handlers.NewThankYouPageHandler(&clerkFrontendConfig)

	injectActiveSession := clerk.WithSessionV2(clerkClient)
	requireActiveSession := clerk.RequireSessionV2(clerkClient)
	landingPageHandler = injectActiveSession(landingPageHandler)
	welcomePageHandler = requireActiveSession(welcomePageHandler)
	votesPageHandler = requireActiveSession(votesPageHandler)
	thankYouPageHandler = requireActiveSession(thankYouPageHandler)

	r := mux.NewRouter()
	r.Handle("/", landingPageHandler)
	r.Handle("/welcome", welcomePageHandler)
	r.Handle("/votes", votesPageHandler)
	r.Handle("/votes", votesPageHandler)
	r.Handle("/thank-you", thankYouPageHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Listening on address: %s\n", address)
	http.ListenAndServe(address, r)
}
