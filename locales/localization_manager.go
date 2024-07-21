package localization

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	instance *LocalizationManager
	once     sync.Once
)

// GetInstance returns the global instance of LocalizationManager.
func Instance() *LocalizationManager {

	return instance
}

func InitLocalization(defaultLang string) {
	once.Do(func() {
		instance = newLocalizationManager(defaultLang)
	})
}

// LocalizationManager handles loading and fetching translations.
type LocalizationManager struct {
	bundle      *i18n.Bundle
	localizers  map[string]*i18n.Localizer
	loadedLangs map[string]bool
	mu          sync.RWMutex
}

// NewLocalizationManager creates a new LocalizationManager.
func newLocalizationManager(defaultLang string) *LocalizationManager {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	manager := &LocalizationManager{
		bundle:      bundle,
		localizers:  make(map[string]*i18n.Localizer),
		loadedLangs: make(map[string]bool),
	}

	// Load default language files
	if err := manager.LoadLanguage(defaultLang); err != nil {
		log.Fatalf("Failed to load default language files: %v", err)
	}

	return manager
}

// LoadLanguage loads translation files for a specific language.
func (lm *LocalizationManager) LoadLanguage(lang string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if lm.loadedLangs[lang] {
		return nil
	}

	if err := lm.loadLanguageFiles(lang); err != nil {
		return err
	}

	lm.localizers[lang] = i18n.NewLocalizer(lm.bundle, lang)
	lm.loadedLangs[lang] = true
	return nil
}

// loadLanguageFiles loads translation files for a specific language.
func (lm *LocalizationManager) loadLanguageFiles(lang string) error {
	_, err := lm.bundle.LoadMessageFile(fmt.Sprintf("%s%s%s", util.LocalizationPath, lang, util.LocalizationType))
	return err
}

// Translate fetches a localized string for a given language, message ID, and positional parameters.
func (lm *LocalizationManager) Translate(lang, messageID string, args ...interface{}) string {
	lm.mu.RLock()
	localizer, exists := lm.localizers[lang]
	lm.mu.RUnlock()

	// Load language if not already loaded
	if !exists {
		if err := lm.LoadLanguage(lang); err != nil {
			return fmt.Sprintf("[Failed to load language %s: %s]", lang, err)
		}
		lm.mu.RLock()
		localizer = lm.localizers[lang]
		lm.mu.RUnlock()
	}

	// Prepare parameter map
	params := make(map[string]interface{})
	for i, arg := range args {
		params[fmt.Sprintf("arg%d", i)] = arg
	}

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: params,
	})
	if err != nil {
		return fmt.Sprintf("[Error localizing message: %s]", err)
	}
	return message
}
