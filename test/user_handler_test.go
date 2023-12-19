package test

// type testdb struct {
// 	db.UserStore
// }

// func setupMongo() *testdb {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &testdb{
// 		Store: {UserStore:db.NewMongoUserStore(client)}
// 	}
// }

// func (tdb *testdb) tearDown(t *testing.T) {
// 	if err := tdb.Drop(context.TODO()); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestPost(t *testing.T) {
// 	tdb := setupMongo()
// 	defer tdb.tearDown(t)

// 	app := fiber.New()
// 	userHandler := handler.NewUserHandler(tdb.UserStore)
// 	app.Post("/", userHandler.HandlePostUser)
// 	req := httptest.NewRequest("POST", "/", nil)
// 	req.Header.Add("Content-Type", "application/json")
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(resp)

// }
