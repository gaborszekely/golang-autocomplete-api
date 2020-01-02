package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gaborszekely/golang-autocomplete-api/pkg/autocomplete"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Run - Start our API server
func Run() {
	client := getMongoClient()
	collection := client.Database(DBName).Collection(CollectionName)
	wordTrie := getRootNode(collection)

	router := mux.NewRouter()

	// Handle /api/type endpoint
	router.HandleFunc("/api/type", func(w http.ResponseWriter, r *http.Request) {
		handleType(w, r, wordTrie, collection)
	}).Methods("POST")

	// Handle /api/clear endpoint
	router.HandleFunc("/api/clear", func(w http.ResponseWriter, r *http.Request) {
		handleClear(w, r, wordTrie, collection)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func getMongoClient() *mongo.Client {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getRootNode(collection *mongo.Collection) *autocomplete.TrieNode {
	findResult := collection.FindOne(context.TODO(), bson.M{"value": ""})
	err := findResult.Err()

	if err != nil {
		newRootNode := createRootNode()
		collection.InsertOne(context.TODO(), newRootNode)
		return getRootNode(collection)
	}

	rootNode := autocomplete.TrieNode{}

	err = findResult.Decode(&rootNode)
	if err != nil {
		log.Fatal(err)
	}

	return &rootNode
}

func createRootNode() *autocomplete.TrieNode {
	return &autocomplete.TrieNode{
		Val:        "",
		Children:   make([](*autocomplete.TrieNode), 0),
		IsEnd:      false,
		Occurences: 0,
	}
}
