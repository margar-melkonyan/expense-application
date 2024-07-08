package service

import (
	"expense-application/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"reflect"
)

type TgService struct {
	repository repository.Category
}

func NewTgService(repository repository.Category) *TgService {
	return &TgService{
		repository: repository,
	}
}

func (s *TgService) CommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "start":
		msg.Text = "Buttons:"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			s.CreateKeyboard(
				[]string{
					"/income",
					"/expense",
					"/incomes",
					"/expenses",
					"/categories",
				},
				2,
			)...,
		)
	case "categories":
		msg.Text = "All categories:"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			s.CreateKeyboard(
				[]string{
					"/start",
				},
				2,
			)...,
		)
	}

	return s.SendMessage(bot, msg, update)
}

func (s *TgService) SendMessage(
	bot *tgbotapi.BotAPI,
	message tgbotapi.MessageConfig,
	update tgbotapi.Update,
) error {
	if reflect.TypeOf(message.Text).Kind() == reflect.String && update.Message.Text == "" {
		message.Text = "Use commands!"
	}

	if _, err := bot.Send(message); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func (s *TgService) CreateKeyboard(commands []string, commandsPerRow int) [][]tgbotapi.KeyboardButton {
	chunkSize := (len(commands) - 1) / commandsPerRow
	var buttons []tgbotapi.KeyboardButton
	var buttonRows [][]tgbotapi.KeyboardButton

	for i := 0; i < len(commands); i++ {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(commands[i]))
	}

	for i := 0; i < len(commands); i += chunkSize {
		end := i + chunkSize

		if end > len(commands) {
			end = len(commands)
		}

		buttonRows = append(buttonRows, tgbotapi.NewKeyboardButtonRow(buttons[i:end]...))
	}

	return buttonRows
}
