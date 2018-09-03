package base256

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var (
	decodeSpec = map[string]string{
		"x💀🍵👯233🎂💁D":       "Equim",
		"你好👾🍧🙆🍬🙇🕷👦🍯🚶🍬👱🕸世界": "Hello World!",
		"❤❤❤🦁🌐👈🚁🚴🏤🦁🗾🏍🚁🚴🏟🦁🗺👀🚁🚵🏔🦁🗾🖐🚁🏋🗺🐅🕍💓🚠💋🚈🦁🗺👉🚁🚴💒🐖🚇🚴🌺🤑🍠👟🏝👼🌼🙃🚁🏋🏗": "「さやかちゃん、大好きだ！(*^ω^*)」",
		"🐗🏰🖖🌅🐗🏰🖕🚆🐈🌐👉🛀👃🏡💬🕷🐗🏰🖖🌄🐗🏰🖕🚆🐈🌐👉🛀👃🏡💬🚠🤘💈🐯🚆💙🚠💥🕋🐯🚅💫🚡🙏🏝🐅🏞🙌🛎👭🏙😕": "👩🏻‍💻 👨🏻‍💻咱们工人有力量(",
		"🐏🏘💏🌾😮🛀👲🚘🏇🍏": "\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f",
		"":     "",
		"    ": "",
	}
)

func TestDecode(t *testing.T) {
	for k, v := range decodeSpec {
		if actual := string(DecodeString(k)); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func TestDecoder(t *testing.T) {
	for k, v := range decodeSpec {
		d := NewDecoder(bytes.NewReader([]byte(k)))
		actual, err := ioutil.ReadAll(d)
		if err != nil {
			t.Fatal(err)
		}

		if string(actual) != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}
