package base256

import (
	"bytes"
	"testing"
)

var (
	encodeSpec = map[string]string{
		"Equim":        "💀🍵👯🎂💁",
		"Hello World!": "👾🍧🙆🍬🙇🕷👦🍯🚶🍬👱🕸",
		"「さやかちゃん、大好きだ！(*^ω^*)」":                    "🦁🌐👈🚁🚴🏤🦁🗾🏍🚁🚴🏟🦁🗺👀🚁🚵🏔🦁🗾🖐🚁🏋🗺🐅🕍💓🚠💋🚈🦁🗺👉🚁🚴💒🐖🚇🚴🌺🤑🍠👟🏝👼🌼🙃🚁🏋🏗",
		"👩🏻‍💻 👨🏻‍💻咱们工人有力量(":                        "🐗🏰🖖🌅🐗🏰🖕🚆🐈🌐👉🛀👃🏡💬🕷🐗🏰🖖🌄🐗🏰🖕🚆🐈🌐👉🛀👃🏡💬🚠🤘💈🐯🚆💙🚠💥🕋🐯🚅💫🚡🙏🏝🐅🏞🙌🛎👭🏙😕",
		"\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f": "🐏🏘💏🌾😮🛀👲🚘🏇🍏",
		"": "",
	}
)

func TestEncode(t *testing.T) {
	for k, v := range encodeSpec {
		if actual := EncodeToString([]byte(k)); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func TestEncoder(t *testing.T) {
	for k, v := range encodeSpec {
		buf := new(bytes.Buffer)
		e := NewEncoder(buf)
		if _, err := e.Write([]byte(k)); err != nil {
			t.Fatal(err)
		}
		if err := e.Close(); err != nil {
			t.Fatal(err)
		}
		if actual := string(buf.Bytes()); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}
