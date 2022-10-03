package translate

import (
	"fmt"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestTranslation(t *testing.T) {
	Language = "fr"
	err := Start()
	if err != nil {
		t.Error(err)
	}
	err = fmt.Errorf("error message: %d", 1)
	devError := Localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DeviceErrors",
			Other: "Dev Err English {{.Err}}",
		},
		TemplateData: map[string]string{
			"Err": err.Error(),
		},
	})
	if devError != "Dev Err French error message: 1" {
		t.Log(devError, Language)
		t.Error("Translation not working")
	}

	name := "Bob"
	helloPerson := Localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HelloPerson",
			Other: "Hello {{.Name}}",
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})
	if helloPerson != "Baguette Bob" {
		t.Error(helloPerson, Language)
	}
}
