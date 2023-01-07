package test

import (
	"context"
	"log"
	"math/rand"
	"testing"
	"time"

        "pilot_server/ent" // !!! please change correct root module name

	_ "github.com/lib/pq" // SHOULD import postgresql driver
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func makeRandStr(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getClient() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=agentdb password=agentdb dbname=agentdb sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgresql: %v", err)
	}
	//defer client.Close()

	// Run the auto migration tool.
	ctx := context.Background()
	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func TestCreateNode(t *testing.T) {
	ctx := context.Background()
	client := getClient()
	defer client.Close()

	n, err := client.Node.
		Create().
		SetID(makeRandStr(10)).
		SetTimestamp("11:11:11").
		SetPodInfo("mypod").
		Save(ctx)

	if err != nil {
		t.Error("failed creating node: ", err)
	}
	t.Log("node was created: ", n)
}

func TestReadNode(t *testing.T) {
	ctx := context.Background()
	client := getClient()
	defer client.Close()

	nodes, err := client.Node.Query().All(ctx)

	if err != nil {
		t.Error("failed read node: ", err)
	}

	for _, n := range nodes {
		log.Printf("%v\n", n)
	}
}

func TestUpdateNode(t *testing.T) {
	ctx := context.Background()
	client := getClient()
	defer client.Close()

	nodes, err := client.Node.Query().All(ctx)

	if err != nil {
		t.Error("failed creating node: ", err)
	}

	for _, n := range nodes {
		log.Printf("%v\n", n)
	}
}
