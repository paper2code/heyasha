package platforms

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"io/ioutil"
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
	}
	var msg = messenger.Receive{}
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s", string(x))
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
				"EAAIkwCguBLUBAAVB0PQJDtfn7rI8HLMb17ouGxfkXDhRFvYsWTwcj0UfJ1mKXRrVJSRSZBlpbYVIZCzRpazEZCoSdtWyQROPhaogz4qNROZCiVVNfS178CQFjgQ3dolApZAaareEmo9yXIQ0zzD6rzrU36efHyXI3gY3OoZCgVagZDZD",
				messenger.Recipient{info.Sender.ID},
			}

			//ai_msg := engine.BotReply(strconv.FormatInt(info.Message.Sender.ID, 10), info.Message.Text)
			resp.Text(info.Message.Text)
		}
	}
	c.String(http.StatusOK, "Webhook Success!!!")

}
