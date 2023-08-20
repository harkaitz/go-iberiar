package iberiar

import (
	"regexp"
	"strings"
)

// Rune .
type Rune struct {
	From      string
	ToLatin   string
	ToIberian string
}

// SillabaryNorthEastNonDual .
var SillabaryNorthEastNonDual []Rune = []Rune {
	{"  *"         ,""    , ""},
	{"[,.;]"       ,"·"   , ""},
	{"h"           ,""    , ""},
	/* ===========================[ORIG]===== */
	{"a"     ,"A"     , ""},
	{"e"     ,"E"     , ""},
	{"i"     ,"I"     , ""},
	{"o"     ,"O"     , ""},
	{"u"     ,"U"     , ""},
	{"[gk]a" ,"(G/K)A", ""},
	{"[gk]e" ,"(G/K)E", ""},
	{"[gk]i" ,"(G/K)I", ""},
	{"[gk]o" ,"(G/K)O", ""},
	{"[gk]u" ,"(G/K)U", ""},
	{"ba"    ,"BA"    , ""},
	{"be"    ,"BE"    , ""},
	{"bi"    ,"BI"    , ""},
	{"bo"    ,"BO"    , ""},
	{"bu"    ,"BU"    , ""},
	{"[dt]a" ,"(D/T)A", ""},
	{"[dt]e" ,"(D/T)E", ""},
	{"[dt]i" ,"(D/T)I", ""},
	{"[dt]o" ,"(D/T)O", ""},
	{"[dt]u" ,"(D/T)U", ""},
	{"t?z"   ,"Z"     , ""},
	{"s"     ,"S"     , ""},
	{"x"     ,"S"     , ""},
	{"rr"    ,"RR"    , ""},
	{"r"     ,"R"     , ""},
	{"l"     ,"L"     , ""},
	{"m"     ,"M"     , ""},
	{"n|ñ"   ,"N"     , ""},
	{"p"     ,"P"     , ""},/* (m') */
	/* ============================[EXTRA]==== */
	{"k"     ,"[K]", ""},
	{"t"     ,"[T]", ""},
}

// SillabaryNorthEastDual .
var SillabaryNorthEastDual []Rune = []Rune {
	{"  *"         ,""    , ""},
	{"[,.;]"       ,"·"   , ""},
	{"h"           ,""    , ""},
	/* ===========================[ORIG]===== */
	{"a"     ,"A"  , ""},
	{"e"     ,"E"  , ""},
	{"i"     ,"I"  , ""},
	{"o"     ,"O"  , ""},
	{"u"     ,"U"  , ""},
	{"ga"    ,"GA" , ""},
	{"ge"    ,"GE" , ""},
	{"gi"    ,"GI" , ""},
	{"go"    ,"GO" , ""},
	{"ka"    ,"KA" , ""},
	{"ke"    ,"KE" , ""},
	{"ki"    ,"KI" , ""},
	{"ko"    ,"KO" , ""},
	{"[gk]u" ,"KU" , ""},
	{"ba"    ,"BA" , ""},
	{"be"    ,"BE" , ""},
	{"bi"    ,"BI" , ""},
	{"bo"    ,"BO" , ""},
	{"bu"    ,"BU" , ""},
	{"da"    ,"DA" , ""},
	{"de"    ,"DE" , ""},
	{"di"    ,"DI" , ""},
	{"do"    ,"DO" , ""},
	{"du"    ,"DU" , ""},
	{"ta"    ,"TA" , ""},
	{"te"    ,"TE" , ""},
	{"ti"    ,"TI" , ""},
	{"to"    ,"TO" , ""},
	{"tu"    ,"TU" , ""},
	{"t?z"   ,"Z"  , ""},
	{"s"     ,"S"  , ""},
	{"x"     ,"S"  , ""},
	{"rr"    ,"RR" , ""},
	{"r"     ,"R"  , ""},
	{"l"     ,"L"  , ""},
	{"m"     ,"M"  , ""},
	{"n"     ,"N"  , ""},
	{"p"     ,"P"  , ""},
	/* ============================[EXTRA]==== */
}

// ToIberian writes the script in `txt` using the runes in
// `runes`.
func ToIberian(txt string, runes []Rune) (iberian string, latin string, err error) {
	/* Compile regular expressions. */
	regexes := make([]*regexp.Regexp, len(runes));
	for n, r := range runes {
		regexes[n], err = regexp.Compile("\\A"+r.From)
		if err != nil {
			return
		}
	}
	/* Match, shift loop. */
	s := strings.ToLower(strings.Trim(txt, " \n.;,"))
	for len(s)>0 {
		found := false
		if s[0:1] == "\r" {
			s = s[1:]
			continue
		} else if s[0:1] == "\n" {
			latin += "\n"
			iberian += "\n"
			s = s[1:]
			continue
		}
		for n, r := range regexes {
			idx := r.FindStringIndex(s)
			if idx == nil {
				continue
			}
			latin   += runes[n].ToLatin
			iberian += runes[n].ToIberian
			found = true
			s = s[idx[1]:]
			break
		}
		if !found {
			latin += "(" + s[0:1] + ")"
			s = s[1:]
		}
	}
	/* Return the result. */
	return
}
