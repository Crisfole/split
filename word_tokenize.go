package word

import "regexp"

type replacer struct {
	regex *regexp.Regexp
	repl  string
}

func (r *replacer) replace(str string) string {
	return r.regex.ReplaceAllString(str, r.repl)
}

// A list of replacer{  applied in sequenc }.
// The resulting string is split on whitespace.
// (Adapted from splitta, which adapted it from
// the Punkt Word Tokenizer)
var tokenizers []replacer = []replacer{
	// uniform quotes
	replacer{regexp.MustCompile(`''`), `"`},
	replacer{regexp.MustCompile(`´´`), `"`},
	replacer{regexp.MustCompile("``"), `"`},

	// Separate punctuation (except period) from words:
	replacer{regexp.MustCompile(`(^|\s)(')`), `$1$2 `},
	replacer{regexp.MustCompile(`([\(\"\` + "`" + `{\[:;&\#\*@])(.)`), `$1 $2`},

	replacer{regexp.MustCompile(`(.)([?!\)";}\]*:@'])`), `$1 $2`},
	replacer{regexp.MustCompile(`([\)}\]])(.)`), `$1$2 `},
	replacer{regexp.MustCompile(`(.)([({\[])`), `$1 $2`},
	replacer{regexp.MustCompile(`((^|\s)\-)([^\-])`), `$1 $2`},

	// Treat double-hyphen as one token:
	replacer{regexp.MustCompile(`([^-])(\-\-+)([^-])`), `$1 $2 $3`},
	replacer{regexp.MustCompile(`(\s|^)(,)(\S)`), `$1$2 $3`},

	// Only separate comma if space follows:
	replacer{regexp.MustCompile(`(.)(,)(\s|$)`), `$1 $2$3`},

	// Combine dots separated by whitespace to be a single token:
	replacer{regexp.MustCompile(`\.\s\.\s\.`), `...`},

	// Separate "No.6"
	replacer{regexp.MustCompile(`([A-Za-z]\.)(\d+)`), `$1 $2`},

	// Separate words from ellipses
	replacer{regexp.MustCompile(`([^\.]|^)(\.{2,})(.?)`), `$1 $2 $3`},
	replacer{regexp.MustCompile(`(^|\s)(\.{2,})([^\.\s])`), `$1$2 $3`},
	replacer{regexp.MustCompile(`([^\.\s])(\.{2,})($|\s)`), `$1 $2$3`},

	//# adding a few things here:

	// fix %, $, &
	replacer{regexp.MustCompile(`(\d)%`), `$1 %`},
	replacer{regexp.MustCompile(`\$(\.?\d)`), `$ $1`},
	replacer{regexp.MustCompile(`(\w)& (\w)`), `$1&$2`},
	replacer{regexp.MustCompile(`(\w\w+)&(\w\w+)`), `$1 & $2`},

	// fix (n `t) --> ( n`t)
	replacer{regexp.MustCompile(`n 't( |$)`), " n't$1"},
	replacer{regexp.MustCompile(`N 'T( |$)`), " N'T$1"},

	// treebank tokenizer special words
	// Tweaked because Go's Regex parser has a bizzare behavior
	replacer{regexp.MustCompile(`([Cc]an)(not)`), `$1 $2`},

	replacer{regexp.MustCompile(`\s+`), ` `},
}

func Tokenize(str string) string {
	for _, r := range tokenizers {
		str = r.replace(str)
	}
	return str
}
