package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gitlab.com/qr-through/entry/backend/infrastructure"
	"gitlab.com/qr-through/entry/backend/internal/core/domain"
	"gitlab.com/qr-through/entry/backend/pkg/errors"
	"gitlab.com/qr-through/entry/backend/pkg/util"
)

type lineHandler struct {
	serv domain.LineService
}

func NewLineHandler(serv domain.LineService) *lineHandler {
	return &lineHandler{
		serv: serv,
	}
}

func (h lineHandler) Webhook(c *fiber.Ctx) error {
	httpRequest := new(http.Request)
	if err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true); err != nil {
		log.Panic(err)
		return util.ResponseError(c, errors.NewInternalError("can't convert request from gofiber to http"))
	}

	events, err := infrastructure.LineBot.ParseRequest(httpRequest)
	if err != nil {
		log.Panic(err)
		return util.ResponseBadRequest(c)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "Get QR" {
					qrString := "https://api.qrserver.com/v1/create-qr-code/?size=320x320&margin=10&data=" + "QRTHROUGH:" + message.ID
					// Handle text messages
					messageId, err := strconv.ParseInt(message.ID, 10, 64)
					if err != nil {
						log.Panic(err)
						errMsg := "เกิดข้อผิดพลาด: ไม่สามารถแปลง message ID เป็น Integer ได้ (Bad Request)."
						if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errMsg)).Do(); err != nil {
							log.Println(err)
						}
						break
					}

					if err = h.serv.CreateQR(messageId, event.Source.UserID); err != nil {
						if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(err.Error())).Do(); err != nil {
							log.Println(err)
						}
						break
					}

					qrCode := linebot.NewImageMessage(qrString, qrString)
					if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, qrCode).Do(); err != nil {
						log.Panic(err)
					}
				}
			}
		}
	}
	return util.ResponseOK(c, nil)
}
