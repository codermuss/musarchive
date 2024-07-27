package localization

import (
	"encoding/json"
	"fmt"

	"sync"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

type LocalizationManager struct {
	bundle     *i18n.Bundle
	localizers sync.Map // concurrent-safe map for localizers
}

var (
	instance *LocalizationManager
	once     sync.Once
	initErr  error

	supportedLangs = []string{"en"}
	defaultLang    = "en"
)

// Initialize initializes the LocalizationManager singleton
func Initialize(path string) error {
	once.Do(func() {
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

		instance = &LocalizationManager{
			bundle: bundle,
		}

		// Load all supported languages at startup
		for _, lang := range supportedLangs {
			if err := instance.loadLanguageFiles(path, lang); err != nil {
				initErr = fmt.Errorf("failed to load language files for %s: %v", lang, err)
				return
			}
			instance.localizers.Store(lang, i18n.NewLocalizer(instance.bundle, lang))
		}

		// Ensure the default language is loaded
		if err := instance.loadLanguageFiles(path, defaultLang); err != nil {
			initErr = fmt.Errorf("failed to load default language files: %v", err)
			return
		}
		instance.localizers.Store(defaultLang, i18n.NewLocalizer(instance.bundle, defaultLang))
	})
	return initErr
}

// GetInstance returns the singleton instance of LocalizationManager
func GetInstance() *LocalizationManager {
	return instance
}

// loadLanguageFiles loads the translation files for a given language
func (lm *LocalizationManager) loadLanguageFiles(path, lang string) error {
	_, err := lm.bundle.LoadMessageFile(path)
	log.Info().Msgf("%s%s%s", util.LocalizationPath, lang, util.LocalizationType)
	return err
}

// Translate retrieves the localized message for the given language and message ID
func (lm *LocalizationManager) Translate(lang, messageID string, args ...interface{}) string {
	localizer, _ := lm.localizers.Load(lang)

	if loc, ok := localizer.(*i18n.Localizer); ok {
		params := make(map[string]interface{})
		for i, arg := range args {
			params[fmt.Sprintf("arg%d", i)] = arg
		}

		message, err := loc.Localize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: params,
		})
		if err != nil {
			return fmt.Sprintf("[Error localizing message: %s]", err)
		}
		return message
	}

	return "[Localizer not found]"
}

// GetSupportedLanguages returns the list of supported languages
func GetSupportedLanguages() []string {
	return supportedLangs
}
