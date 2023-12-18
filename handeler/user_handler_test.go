package handeler

import (
	"context"
	"fmt"
	"hotel-reservation/db"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.UserStore
}

func setupMongo() *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, db.TESTDBNAME),
	}
}

func (tdb *testdb) tearDown(t *testing.T) {
	if err := tdb.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPost(t *testing.T) {
	tdb := setupMongo()
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)

}
