package handlers

// TestGreet tests the GET /greet endpoint
// func TestGreet(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	router := gin.Default()     // Create a new router
// 	router.GET("/greet", Greet) // Registering the handler to the router

// 	req, _ := http.NewRequest("GET", "/greet", nil) // mock http request is prepared
// 	resp := httptest.NewRecorder()                  // response recorder

// 	router.ServeHTTP(resp, req) // process the request through the router and record the generated response in resp.

// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "Hello, how are you?")
// }

// // TestGreetToName tests the POST /greet endpoint with a name
// func TestGreetToName(t *testing.T) {
// 	router := gin.Default()
// 	router.POST("/greet", GreetToName)

// 	user := model.User{Name: "Rohan"}
// 	userJson, _ := json.Marshal(user) // returns the json encoding in []byte format

// 	// both bytes.NewBuffer(jsonValue) and strings.NewReader(string(jsonValue)) can be used
// 	req, _ := http.NewRequest("POST", "/greet", bytes.NewBuffer(userJson))
// 	resp := httptest.NewRecorder()

// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	assert.Contains(t, resp.Body.String(), "Hello Rohan, how are you?")
// }
