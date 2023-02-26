package common_model

import "regexp"

const URLPattern = `/(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})/gi`

var urlCheck, _ = regexp.Compile(URLPattern)

type URL string

func (u URL) IsValid() bool {
	return urlCheck.MatchString(u.String())
}

func (u URL) String() string {
	return string(u)
}
