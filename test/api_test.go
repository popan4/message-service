package test

import (
  "bytes"
  "encoding/json"
  _ "encoding/json"
  "github.com/gofiber/fiber/v2"
  "github.com/stretchr/testify/assert"
  "message-service/controller"
  "message-service/logger"
  "message-service/model"
  "message-service/repository"
  "message-service/router"
  "message-service/service"
  "net/http"
  "net/http/httptest"
  "testing"
)

func setupApp() *fiber.App {
  logger.Init()
  msgRepo := repository.NewMessageRepository()
  msgService := service.NewMessageService(msgRepo)
  ctrl := controller.NewMessageController(msgService)

  app := fiber.New()
  router.SetupRoutes(app, ctrl)
  return app
}

func createMessage(t *testing.T, app *fiber.App, text string) model.Message {
  body := []byte(`{"messageText":"` + text + `"}`)
  req := httptest.NewRequest(http.MethodPost, "/createMsg", bytes.NewReader(body))
  req.Header.Set("Content-Type", "application/json")
  resp, err := app.Test(req, -1)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusCreated, resp.StatusCode)

  var msg model.Message
  err = json.NewDecoder(resp.Body).Decode(&msg)
  resp.Body.Close()
  assert.NoError(t, err)
  return msg
}
func TestMessage_Create(t *testing.T) {
  app := setupApp()
  testMessageData := []struct {
	text         string
	isPalindrome bool
  }{
	{"pojop", true},
	{"pooja", false},
  }
  for _, tmd := range testMessageData {
	msg := createMessage(t, app, tmd.text)
	assert.Equal(t, tmd.text, msg.Text)
	assert.Equal(t, tmd.isPalindrome, msg.IsPalindrome)
  }
}

func TestMessage_GetAllMessages(t *testing.T) {
  app := setupApp()
  // create 2 messages
  createMessage(t, app, "1234321")
  createMessage(t, app, "hello")
  req := httptest.NewRequest(http.MethodGet, "/getAllMsg", nil)
  resp, err := app.Test(req, -1)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusOK, resp.StatusCode)

  var messages []model.Message
  err = json.NewDecoder(resp.Body).Decode(&messages)
  resp.Body.Close()
  assert.NoError(t, err)
  assert.GreaterOrEqual(t, len(messages), 2)
}

func TestMessage_Update(t *testing.T) {
  app := setupApp()

  message := createMessage(t, app, "madam")
  updateBody := []byte(`{"messageText":"hello"}`)
  req := httptest.NewRequest(http.MethodPut, "/updateMsg/"+message.Id, bytes.NewReader(updateBody))
  req.Header.Set("Content-Type", "application/json")
  resp, err := app.Test(req, -1)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusOK, resp.StatusCode)

  var updatedMessage model.Message
  err = json.NewDecoder(resp.Body).Decode(&updatedMessage)
  resp.Body.Close()
  assert.NoError(t, err)
  assert.Equal(t, "hello", updatedMessage.Text)
  assert.False(t, updatedMessage.IsPalindrome)
}

func TestMessage_Delete(t *testing.T) {
  app := setupApp()

  message := createMessage(t, app, "test")
  req := httptest.NewRequest(http.MethodDelete, "/DeleteMsg/"+message.Id, nil)
  resp, err := app.Test(req, -1)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestMessage_GetByID(t *testing.T) {
  app := setupApp()

  message := createMessage(t, app, "1234321")
  req := httptest.NewRequest(http.MethodGet, "/getMsg/"+message.Id, nil)
  resp, err := app.Test(req, -1)
  assert.NoError(t, err)
  assert.Equal(t, http.StatusOK, resp.StatusCode)

  var existingMsg model.Message
  err = json.NewDecoder(resp.Body).Decode(&existingMsg)
  resp.Body.Close()
  assert.NoError(t, err)

  assert.Equal(t, message.Id, existingMsg.Id)
  assert.Equal(t, message.Text, existingMsg.Text)
  assert.True(t, existingMsg.IsPalindrome)
}
