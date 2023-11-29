package iberiar

import (
	"github.com/snapcore/go-gettext"
	"embed"
)

// L10N is an localizable (l10n) text. You can translate
// it with "GetString()".
type L10N interface {
	GetUserMessage()     string
	GetDomainName()      string
	GetDomainLocaleDir() string
	GetDomainLocaleFS()  embed.FS
}

// GetString translates the message to one of the languages. If
// none is found the default message is returned.
func GetString(msg L10N, languages ...string) string {
	var catalog gettext.Catalog
	var domain  gettext.TextDomain
	domain = getDomain(msg)
	if len(languages) == 0 {
		catalog = domain.UserLocale()
	} else {
		catalog = domain.Locale(languages...)
	}
	return catalog.Gettext(msg.GetUserMessage())
}

// GetError returns the error string intended for the user.
func GetError(err error, languages ...string) string {
	e, ok := err.(L10N)
	if !ok { return "Internal error" }
	return GetString(e, languages...)
}

// GetField returns the error field, or "" if none.
func GetField(err error) string {
	type hasField interface {
		GetField() string
	}
	e, ok := err.(hasField)
	if !ok { return "" }
	return e.GetField()
}

func getDomain(msg L10N) gettext.TextDomain {
	return gettext.TextDomain{
		Name: msg.GetDomainName(),
		LocaleDir: msg.GetDomainLocaleDir(),
		LocaleFS: msg.GetDomainLocaleFS(),
	}
}
