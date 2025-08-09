package test

import (
  "bytes"
  _ "encoding/json"
  "github.com/gofiber/fiber/v2"
  "github.com/stretchr/testify/assert"
  "message-service/controller"
  "message-service/router"
  "message-service/service"
  "net/http"
  "net/http/httptest"
  "testing"
)

func setupApp() *fiber.App {

  msgService := service.NewMessageService()
  ctrl := controller.NewMessageController(msgService)

  app := fiber.New()
  router.SetupRoutes(app, ctrl)
  return app
}

func TestMessageCRUD(t *testing.T) {
  app := setupApp()
  // Create
  body := []byte(`{"messageText":"PoojaTest"}`)
  req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
  req.Header.Set("Content-Type", "application/json")
  resp, _ := app.Test(req)
  assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
