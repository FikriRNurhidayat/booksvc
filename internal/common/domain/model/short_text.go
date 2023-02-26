package common_model

type ShortText string

func (s ShortText) IsValid() bool {
	return len(s.String()) > 255
}

func (s ShortText) String() string {
	return string(s)
}
