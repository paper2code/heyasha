package platforms

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"net/http"
)

func MessengerBot(c *gin.Context) {
	verify_token := "er7Wq4yREXBKpdRKjhAg"

	hub_mode := c.Query("hub.mode")
	hub_challenge := c.Query("hub.challenge")
	hub_verify_token := c.Query("hub.verify_token")
	if hub_verify_token == verify_token && hub_challenge != "" {
		c.Header("Hub Mode", hub_mode)
		c.String(http.StatusOK, hub_challenge)
	} else {

	var msg = messenger.Receive{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Errorln("Something wrong: %s\n", err.Error())
		return
	}

	for _, entry := range msg.Entry {
		for _, info := range entry.Messaging {
			a := messenger.Classify(info, entry)
			if a == messenger.UnknownAction {
				fmt.Println("Unknown action:", info)
				continue
			}

			resp := &messenger.Response{
				"EAAQCVhc9mrYBAHhsJZCMwXBal3pRUfAbnbK11vDjZBOPhbZBh1PZA0NVJzZB56gDkwVaE6Tn9rgVera6r8JsB3M2ZBuZC5tLJOCYZBrdsc2oJQzlVPNZAzFWzX2G7SqyqHpOfZBCxsHiA6adVEV7hQKxBh5XLD2hsrska3W7LZBPDRj7QZDZD",
				messenger.Recipient{info.Sender.ID},
			}

			//ai_msg := engine.BotReply(strconv.FormatInt(info.Message.Sender.ID, 10), info.Message.Text)
			resp.Text(info.Message.Text)
		}
	}
		c.String(http.StatusOK,"Webhook Success!!!")
	}
}
