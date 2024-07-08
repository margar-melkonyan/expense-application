package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"reflect"
	"slices"
	"strconv"
)

const expense = "expense"
const income = "income"

var (
	mainOptions = []string{
		"/income",
		"/expense",
		"/incomes",
		"/expenses",
	}

	periodOptions = []string{
		"/day",
		"/week",
		"/month",
		"/custom",
		"/menu",
	}

	registrationProfileOptions = []string{
		"/confirm",
		"/cancel",
	}

	selectedCategory  = ""
	selectedPeriod    = ""
	selectedType      = ""
	enterBudgetTitle  = false
	enterBudgetAmount = false
	userExists        = true
	user              = model.User{}
	budget            = model.Budget{}
	categoriesName    = []string{}
)

type TgService struct {
	categoryRepository repository.Category
	budgetRepository   repository.Budget
	userRepository     repository.User
}

func NewTgService(
	categoryRepository repository.Category,
	budgetRepository repository.Budget,
	userRepository repository.User,
) *TgService {
	return &TgService{
		categoryRepository: categoryRepository,
		budgetRepository:   budgetRepository,
		userRepository:     userRepository,
	}
}

func (s *TgService) CommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	categoriesName = s.categoryRepository.GetCategoriesName(expense)
	categoriesName = append(categoriesName, "/menu")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "start", "register":
		userExists = false

		_, err := s.userRepository.CurrentTgUser(update.Message.From.ID)

		if err == nil {
			userExists = true
		}

		if !userExists {
			msg.Text = "To register, enter your email address:"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			msg.Text = "Tap on button menu for start working."
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					[]string{
						"/menu",
					},
					1,
				)...,
			)
		}
	case "confirm":
		if user.Email != "" {
			err := s.userRepository.SignUpByTg(&user)
			user.Email = ""

			if err != nil {
				return err
			}
			userExists = true

			msg.Text = "You was successfully registered!"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					mainOptions,
					2,
				)...,
			)
			break
		}
		msg.Text = "Email couldn't be empty!"
	case "cancel":
		msg.Text = "If you changed mind you can write command /register"
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	if userExists {
		switch update.Message.Command() {
		case "menu":
			selectedPeriod = ""
			selectedCategory = ""

			msg.Text = "Select button:"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					mainOptions,
					2,
				)...,
			)
		case "expense":
			msg.Text = "Select expense categories"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					categoriesName,
					len(categoriesName)/2,
				)...,
			)
			selectedType = expense
			categoriesName = s.categoryRepository.GetCategoriesName(selectedType)
			categoriesName = append(categoriesName, "/menu")
		case "day":
			selectedPeriod = "day"
		case "week":
			selectedPeriod = "week"
		case "month":
			selectedPeriod = "month"
		case "custom":
			selectedPeriod = "custom"
		}
	}

	if !userExists && !update.Message.IsCommand() {
		user.Name = fmt.Sprintf("%s %s", update.Message.From.FirstName, update.Message.From.LastName)
		user.Email = update.Message.Text
		user.TgId = update.Message.From.ID
		msg.Text = "Confirm registration"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			s.CreateKeyboard(
				registrationProfileOptions,
				2,
			)...,
		)
	}

	if slices.Contains(categoriesName[:len(categoriesName)-1], update.Message.Text) && categoriesName != nil {
		selectedCategory = update.Message.Text
		msg.Text = fmt.Sprintf("Category \"%s\" was selected", selectedCategory)

		err := s.SendMessage(bot, msg, update)
		if err != nil {
			slog.Error(err.Error())
		}

		msg.Text = "Enter title what was spent on:"

		err = s.SendMessage(bot, msg, update)
		if err != nil {
			return err
		}
		enterBudgetTitle = true
	}

	if enterBudgetTitle {
		budget.User, _ = s.userRepository.CurrentTgUser(update.Message.From.ID)
		budget.Type = selectedType
		budget.Title = update.Message.Text
		msg.Text = "Enter how much was spent on:"
		enterBudgetAmount = true
		enterBudgetTitle = false
	}

	if enterBudgetAmount {
		budget.Amount, _ = strconv.ParseFloat(update.Message.Text, 64)
		err := s.budgetRepository.Create(&budget)
		enterBudgetAmount = false
		if err != nil {
			slog.Error(err.Error())
		}
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
	if commandsPerRow < 1 {
		return nil
	}

	chunkSize := (len(commands) - 1) / commandsPerRow

	if chunkSize < 2 {
		chunkSize += 1
	}

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
