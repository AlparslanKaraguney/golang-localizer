package localizer

import (
	// Import the internal/translations so that it's init() function
	// is run. It's really important that we do this here so that the
	// default message catalog is updated to use our translations
	// *before* we initialize the message.Printer instances below.
	_ "localizer/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Define a Localizer type which stores the relevant locale ID (as used
// in our URLs) and a (deliberately unexported) message.Printer instance
// for the locale.
type Localizer struct {
	ID      string
	printer *message.Printer
}

// Initialize a slice which holds the initialized Localizer types for
// each of our supported locales.
// IETF BCP 47 language tags - https://en.wikipedia.org/wiki/IETF_language_tag
// ISO 3166-1 alpha-2 country codes - https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
// Example ID:      "en-GB", // United Kingdom
var locales = []Localizer{
	// First we add the default locale which is English (United Kingdom)
	{
		// United Kingdom
		ID:      "en-GB",
		printer: message.NewPrinter(language.MustParse("en-GB")),
	},
	{
		// Turkey
		ID:      "tr-TR",
		printer: message.NewPrinter(language.MustParse("tr-TR")),
	},
}

func defaultLocale() Localizer {
	return locales[0]
}

// The Get() function accepts a locale ID and returns the corresponding
// Localizer for that locale. If the locale ID is not supported then
// this returns default locale.
func Get(id string) Localizer {
	for _, locale := range locales {
		if id == locale.ID {

			return locale
		}
	}

	return defaultLocale()
}

// We also add a Translate() method to the Localizer type. This acts
// as a wrapper around the unexported message.Printer's Sprintf()
// function and returns the appropriate translation for the given
// message and arguments.
func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
