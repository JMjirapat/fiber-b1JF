package handler

import (
	"encoding/json"
	"fiber/config"
	"fiber/infrastructure"
	"fiber/internal/core/domain"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Token struct {
	Iss string   `json:"iss"`
	Sub string   `json:"sub"`
	Aud string   `json:"aud"`
	Exp int64    `json:"exp"`
	Iat int64    `json:"iat"`
	Amr []string `json:"amr"`
}

type lineHandler struct {
	serv domain.LineService
}

func NewLineHandler(serv domain.LineService) *lineHandler {
	return &lineHandler{
		serv: serv,
	}
}

func (h lineHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := h.serv.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(*result)
}

func (h lineHandler) Register(c *fiber.Ctx) error {
	cfg := config.Config

	var body domain.LineRegisterBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	values := url.Values{}
	values.Set("id_token", body.TokenID)
	values.Set("client_id", cfg.ChannelID)

	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/verify", values)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(string(bodyBytes))
	}

	decoder := json.NewDecoder(resp.Body)
	var token Token
	if err := decoder.Decode(&token); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := h.serv.RegisterUser(body, token.Sub); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h lineHandler) Webhook(c *fiber.Ctx) error {
	httpRequest := new(http.Request)
	if err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true); err != nil {
		panic(err)
	}

	events, err := infrastructure.LineBot.ParseRequest(httpRequest)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "Get QR" {
					qrString := "https://api.qrserver.com/v1/create-qr-code/?size=320x320&margin=10&data=" + "QRTHROUGH:" + message.ID
					// Handle text messages
					id, err := strconv.ParseInt(message.ID, 10, 64)
					if err != nil {
						log.Panic(err)
						return c.SendStatus(fiber.StatusBadRequest)
					}

					if err = h.serv.CreateQR(id, event.Source.UserID); err != nil {
						errMsg := linebot.NewTextMessage(err.Error())
						if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, errMsg).Do(); err != nil {
							log.Panic(err)
							return c.SendStatus(fiber.StatusInternalServerError)
						}
						return c.SendStatus(200)
					}

					qrMsg := linebot.NewImageMessage(qrString, qrString)
					if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, qrMsg).Do(); err != nil {
						log.Panic(err)
						return c.SendStatus(fiber.StatusInternalServerError)
					}
				}
			}
		}
	}
	return c.SendStatus(fiber.StatusOK)
}
