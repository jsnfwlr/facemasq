package translate

import (
	"fmt"
	"os"

	"facemasq/lib/files"
	"facemasq/lib/logging"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
)

var (
	Language  = "en"
	Bundle    *i18n.Bundle
	Localiser *i18n.Localizer
)

func init() {
	Bundle = i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	Language = os.Getenv("LOCALE")
	if Language == "" {
		Language = "en"
	}
}

func Start() (err error) {
	var dir string

	dir, err = files.GetDir("i18n")
	if err != nil {
		logging.Errorf("error getting i18n folder: %v", err)
		return
	}
	translationFile := fmt.Sprintf("%[2]s%[1]capi%[1]cfacemasq.%[3]s.toml", os.PathSeparator, dir, Language)
	// translationFile := "active.fr.toml"

	if !files.FileExists(translationFile) {
		err = fmt.Errorf("could not find %s", translationFile)
		return
	}
	Bundle.MustLoadMessageFile(translationFile)
	logging.Processf("Translations loaded from %s", translationFile)
	Localiser = i18n.NewLocalizer(Bundle, Language, fmt.Sprintf("%s;q=0.9", Language))
	return
}

func Message(id, msg string) (translation string) {
	translation = Localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    id,
			Other: msg,
		},
	})
	return
}

// func Translatef() {}
