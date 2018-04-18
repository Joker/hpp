package hpp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

var TabStr = []byte("    ")

func isInline(tag []byte) bool {
	switch string(tag) {
	case "a_", "b", "i", "em", "strong", "code", "span", "ins",
		"big", "small", "tt", "abbr", "acronym", "cite", "dfn",
		"kbd", "samp", "var", "bdo", "map", "q", "sub", "sup":
		return true
	default:
		return false
	}
}

func isVoid(tag []byte) bool {
	switch string(tag) {
	case "input", "link", "meta", "hr", "img", "br", "area", "base", "col",
		"param", "command", "embed", "keygen", "source", "track", "wbr":
		return true
	default:
		return false
	}
}

func Print(r io.Reader) []byte {
	var (
		b        = new(bytes.Buffer)
		tokenize = html.NewTokenizer(r)
		depth    = 0
		LongText = false
		prevType html.TokenType
		tagName  []byte
		prvName  []byte
		rb       = regexp.MustCompile(`^\s+\S`)
		re       = regexp.MustCompile(`\S\s+$`)
	)
Loop:
	for {
		nowType := tokenize.Next()

		if nowType != html.TextToken {
			prvName = tagName
			tagName, _ = tokenize.TagName()
		}

		switch nowType {
		case html.StartTagToken:
			if !(isInline(tagName) && prevType == html.TextToken) {
				b.WriteByte('\n')
				b.Write(bytes.Repeat(TabStr, depth))
			}
			b.Write(tokenize.Raw())
			if !isVoid(tagName) {
				depth += 1
			}

		case html.SelfClosingTagToken, html.CommentToken, html.DoctypeToken:
			b.WriteByte('\n')
			b.Write(bytes.Repeat(TabStr, depth))
			b.Write(tokenize.Raw())

		case html.EndTagToken:
			depth -= 1
			switch {
			case !bytes.Equal(prvName, tagName),
				prevType == html.SelfClosingTagToken,
				prevType == html.CommentToken,
				prevType == html.DoctypeToken,
				prevType == html.EndTagToken,
				prevType == html.TextToken && LongText:

				b.WriteByte('\n')
				b.Write(bytes.Repeat(TabStr, depth))
			}
			b.Write(tokenize.Raw())

		case html.TextToken:
			t := bytes.Replace(tokenize.Raw(), []byte{'\t'}, TabStr, -1)
			text := bytes.Trim(t, "\n\r ")
			if re.Match(t) {
				text = append(text, ' ')
			}
			LongText = false
			if len(text) > 0 {
				if bytes.Contains(text, []byte{'\n'}) {
					if !(prevType == html.EndTagToken && isInline(tagName)) {
						b.WriteByte('\n')
						b.Write(bytes.Repeat(TabStr, depth))
					} else {
						if rb.Match(t) {
							text = append([]byte{' '}, text...)
						}
					}
					b.Write(txtFmt(text, depth))
					LongText = true
				} else {
					switch {
					case utf8.RuneCount(text) > 80, prevType != html.StartTagToken:

						if !(prevType == html.EndTagToken && isInline(tagName)) {
							b.WriteByte('\n')
							b.Write(bytes.Repeat(TabStr, depth))
							LongText = true
						} else {
							if rb.Match(t) {
								text = append([]byte{' '}, text...)
							}
						}
					}
					b.Write(text)
				}
			}

		case html.ErrorToken:
			err := tokenize.Err()
			if err.Error() == "EOF" {
				break Loop
			}
			log.Panicln(err)

		}
		prevType = nowType
	}

	b.WriteByte('\n')
	return bytes.TrimLeft(b.Bytes(), "\n\r\t ")
}

func txtFmt(txt []byte, depth int) []byte {
	var (
		min = 1000
		ln  = 0
		f   = func(c rune) bool { return '\n' != c && ' ' != c }
	)
	for _, v := range bytes.FieldsFunc(txt, f) {
		ln = len(bytes.TrimLeft(v, " "))
		if ln > 0 && ln < min {
			min = ln
		}
	}
	var re = regexp.MustCompile(fmt.Sprintf(`\n\s{%d}`, min-1))
	return re.ReplaceAllLiteral(txt, append([]byte{'\n'}, bytes.Repeat(TabStr, depth)...))
}

func PrPrint(in string) string {
	return string(Print(strings.NewReader(in)))
}
