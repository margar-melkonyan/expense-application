package helper

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log/slog"
	"regexp"
	"slices"
	"strings"
)

type ErrorResponse struct {
	Message string              `json:"message"`
	Errors  []map[string]string `json:"errors"`
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func FormatValidationError(err interface{}) ErrorResponse {
	var errorResponse ErrorResponse
	errorResponse.Message = "Check your entered data!"
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, fileLoadErr := bundle.LoadMessageFile("translation/en.json")
	if fileLoadErr != nil {
		slog.Debug(fileLoadErr.Error())
	}

	var additionalValue string
	var additionField string
	var additionalFields = []string{"min", "max", "date", "gt", "gte", "eqfield", "oneof"}
	locale := i18n.NewLocalizer(bundle, "en")
	for _, ver := range err.(validator.ValidationErrors) {
		additionField = ""
		additionalValue = ""

		if slices.Contains(additionalFields, ver.Tag()) {
			additionField = strings.ToLower(ver.ActualTag())
			additionalValue = ver.Param()
		}

		word, wordErr := locale.Localize(&i18n.LocalizeConfig{
			MessageID: fmt.Sprintf("%s", ver.ActualTag()),
			TemplateData: map[string]string{
				"attribute":   ToSnakeCase(ver.StructField()),
				additionField: additionalValue,
			},
		})

		if wordErr != nil {
			slog.Debug(wordErr.Error())
		}

		errorResponse.Errors = append(errorResponse.Errors, map[string]string{
			ToSnakeCase(ver.StructField()): word,
		})
	}

	return errorResponse
}
