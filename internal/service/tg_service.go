package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gosimple/slug"
	"log/slog"
	"math"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

const expense = "expense"
const income = "income"

var (
	registrationProfileOptions = []string{
		"/confirm",
		"/cancel",
	}

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
		"/menu",
	}

	selectedCategory = ""
	selectedType     = expense
	userExists       = true
	budget           = model.Budget{}
	budgetStatus     = ""
	user             = model.User{}
)

type TgService struct {
	categoryRepository repository.Category
	budgetRepository   repository.Budget
	userRepository     repository.User
	pdfService         PDF
	xlsxService        XLSX
}

func NewTgService(
	categoryRepository repository.Category,
	budgetRepository repository.Budget,
	userRepository repository.User,
	pdfService PDF,
	xlsxService XLSX,
) *TgService {
	return &TgService{
		categoryRepository: categoryRepository,
		budgetRepository:   budgetRepository,
		userRepository:     userRepository,
		pdfService:         pdfService,
		xlsxService:        xlsxService,
	}
}

func (s *TgService) CommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	categoriesName := s.categoryRepository.GetCategoriesName(selectedType)
	categoriesName = append(categoriesName, "/menu")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	_, err := s.userRepository.CurrentTgUser(update.Message.From.ID)

	switch update.Message.Command() {
	case "start", "register":
		userExists = false

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
			err = s.userRepository.CreateByTg(&user)
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
			selectedCategory = ""

			msg.Text = "Select button:"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					mainOptions,
					2,
				)...,
			)
		case "expense":
			selectedType = expense
			categoriesName = s.categoryRepository.GetCategoriesName(selectedType)
			categoriesName = append(categoriesName, "/menu")

			msg.Text = fmt.Sprintf("Select %s categories:", selectedType)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					categoriesName,
					len(categoriesName)/2,
				)...,
			)
		case "income":
			selectedType = income
			categoriesName = s.categoryRepository.GetCategoriesName(selectedType)
			categoriesName = append(categoriesName, "/menu")

			msg.Text = fmt.Sprintf("Select %s categories:", selectedType)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					categoriesName,
					2,
				)...,
			)
		case "incomes":
			selectedType = income
			msg.Text = "Select period:"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					periodOptions,
					2,
				)...,
			)
		case "expenses":
			selectedType = expense
			msg.Text = "Select period:"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					periodOptions,
					2,
				)...,
			)
		case "day":
			pdf := s.pdfService.GenDayReport(selectedType, user.Id).GetBytes()
			xlsx := s.xlsxService.GenDayReport(selectedType, user.Id).Bytes()

			_, err := bot.SendMediaGroup(tgbotapi.NewMediaGroup(
				update.Message.Chat.ID, []interface{}{
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.pdf",
						Bytes: pdf,
					}),
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.xlsx",
						Bytes: xlsx,
					}),
				},
			))
			if err != nil {
				slog.Error(err.Error())
			}
		case "week":
			pdf := s.pdfService.GenWeekReport(selectedType, user.Id).GetBytes()
			xlsx := s.xlsxService.GenWeekReport(selectedType, user.Id).Bytes()

			_, err := bot.SendMediaGroup(tgbotapi.NewMediaGroup(
				update.Message.Chat.ID, []interface{}{
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.pdf",
						Bytes: pdf,
					}),
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.xlsx",
						Bytes: xlsx,
					}),
				},
			))
			if err != nil {
				slog.Error(err.Error())
			}
		case "month":
			pdf := s.pdfService.GenMonthReport(selectedType, user.Id).GetBytes()
			xlsx := s.xlsxService.GenMonthReport(selectedType, user.Id).Bytes()

			_, err := bot.SendMediaGroup(tgbotapi.NewMediaGroup(
				update.Message.Chat.ID, []interface{}{
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.pdf",
						Bytes: pdf,
					}),
					tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
						Name:  "report.xlsx",
						Bytes: xlsx,
					}),
				},
			))
			if err != nil {
				slog.Error(err.Error())
			}
		}
	}

	if !userExists && !update.Message.IsCommand() {
		fmt.Println(update.Message.Text)
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

	switch budgetStatus {
	case "write_title":
		budget.User, _ = s.userRepository.CurrentTgUser(update.Message.From.ID)
		budget.Type = selectedType
		budget.Title = update.Message.Text
		msg.Text = "Enter amount:"

		budgetStatus = "write_amount"
	case "write_amount":
		amount, _ := strconv.ParseFloat(update.Message.Text, 64)

		if (!strings.Contains(update.Message.Text, ".") ||
			len(strings.Split(update.Message.Text, ".")[1]) <= 2) && amount != 0.0 {

			category, _ := s.categoryRepository.GetBySlug(slug.Make(selectedCategory))

			budget.Amount = amount
			budget.Amount = math.Round(budget.Amount * 100)
			err := s.budgetRepository.Store(&budget, &category)
			if err != nil {
				slog.Error(err.Error())
			}
			budget = model.Budget{}
			budgetStatus = ""
			selectedCategory = ""

			msg.Text = "Menu:"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				s.CreateKeyboard(
					mainOptions,
					2,
				)...,
			)

		} else {
			msg.Text = "Amount has invalid value"
		}
	}

	if slices.Contains(categoriesName[:len(categoriesName)-1], update.Message.Text) && categoriesName != nil && selectedCategory == "" {
		selectedCategory = update.Message.Text
		msg.Text = fmt.Sprintf("Category \"%s\" was selected", selectedCategory)

		err := s.SendMessage(bot, msg, update)
		if err != nil {
			slog.Error(err.Error())
		}

		msg.Text = "Enter title:"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			s.CreateKeyboard(
				[]string{
					"/menu",
				},
				1,
			)...,
		)

		budgetStatus = "write_title"
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
