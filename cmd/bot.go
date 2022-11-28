package cmd

import (
	"gopkg.in/telebot.v3"
	"log"
	"recipe_bot/internal/models"
	"recipe_bot/internal/repository"
	"time"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *repository.UserModel
}

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Такой пользователь уже существует %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)
		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Hello, " + ctx.Sender().FirstName + "\nDon't know what to cook today? Reciper Teller is here to help you. \n\n 1. Get random recipe with /random \n 2. Get recipe by name /name {enter here recipe name}\n 3. Get random recipe with one ingredient /ingredient {enter ingredient name}")
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}
