package tgbot

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/AnesVijay/glogger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotInit(token string, useProxy bool) *tgbotapi.BotAPI {

	httpClient := &http.Client{}

	if useProxy {
		proxyUrl, err := url.Parse("socks5://127.0.0.1:1080") // Example: local SOCKS5 proxy
		if err != nil {
			glogger.GetLogger().SendError(fmt.Sprintf("invalid proxy for tg bot: %v", err))
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}

	bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		glogger.GetLogger().SendError(fmt.Sprintf("failed to initialize tg bot: %v", err))
	}

	bot.Debug = false
	bot.Self.CanJoinGroups = false

	glogger.GetLogger().SendInfo(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	return bot
}
