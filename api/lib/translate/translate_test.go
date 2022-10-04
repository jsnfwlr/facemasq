package translate

import (
	"fmt"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TranslationTable struct {
	Language string
	Messages []Translation
}

type Translation struct {
	Config   *i18n.LocalizeConfig
	Expected string
}

func TestTranslation(t *testing.T) {
	tableSet := []TranslationTable{
		{
			Language: "en",
			Messages: []Translation{
				{
					Expected: "Unable to retreive Category data",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "RetreiveCategoryError",
							Other: "Unable to retreive Category data",
						},
					},
				},
				{
					Expected: "Default Message",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "NoTranslationAvailable",
							Other: "Default Message",
						},
					},
				},
				{
					Expected: "English Muffin",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "Food",
							Other: "English {{.Food}}",
						},
						TemplateData: map[string]string{
							"Food": "Muffin",
						},
					},
				},
			},
		},
		{
			Language: "es",
			Messages: []Translation{
				{
					Expected: "No se pueden recuperar los datos de la categoría",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "RetreiveCategoryError",
							Other: "Unable to retreive Category data",
						},
					},
				},
				{
					Expected: "Default Message",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "NoTranslationAvailable",
							Other: "Default Message",
						},
					},
				},
				{
					Expected: "Spanish Churros",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "Food",
							Other: "English {{.Food}}",
						},
						TemplateData: map[string]string{
							"Food": "Churros",
						},
					},
				},
			},
		},
		{
			Language: "fr",
			Messages: []Translation{
				{
					Expected: "Impossible de récupérer les données de catégorie",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "RetreiveCategoryError",
							Other: "Unable to retreive Category data",
						},
					},
				},
				{
					Expected: "Default Message",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "NoTranslationAvailable",
							Other: "Default Message",
						},
					},
				},
				{
					Expected: "French Baguette",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "Food",
							Other: "English {{.Food}}",
						},
						TemplateData: map[string]string{
							"Food": "Baguette",
						},
					},
				},
			},
		},
		{
			Language: "zh",
			Messages: []Translation{
				{
					Expected: "无法检索类别数据",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "RetreiveCategoryError",
							Other: "Unable to retreive Category data",
						},
					},
				},
				{
					Expected: "Default Message",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "NoTranslationAvailable",
							Other: "Default Message",
						},
					},
				},
				{
					Expected: "Chinese Dumplings",
					Config: &i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID:    "Food",
							Other: "English {{.Food}}",
						},
						TemplateData: map[string]string{
							"Food": "Dumplings",
						},
					},
				},
			},
		},
	}

	for ts := range tableSet {
		Language = tableSet[ts].Language
		err := Start()
		if err != nil {
			t.Fatal(err)
		}
		for m := range tableSet[ts].Messages {
			output, err := Localiser.Localize(tableSet[ts].Messages[m].Config)
			if err != nil {
				if err.Error() != fmt.Sprintf(`message "NoTranslationAvailable" not found in language "%s"`, tableSet[ts].Language) {
					t.Error(err)
				}
			}
			if output != tableSet[ts].Messages[m].Expected {
				t.Errorf("%s:%s@%d:%d - expected `%s` got `%s` ", tableSet[ts].Language, tableSet[ts].Messages[m].Config.DefaultMessage.ID, ts, m, tableSet[ts].Messages[m].Expected, output)
			}
		}
	}
	// name := "Bob"
	// helloPerson := Localiser.MustLocalize(&i18n.LocalizeConfig{
	// })
	// if helloPerson != "Baguette Bob" {
	// 	t.Error(helloPerson, Language)
	// }
}
