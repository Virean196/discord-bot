package ping

func Handle(content string) (string, bool) {
	if content == "!ping" {
		return "Pong!", true
	}

	return "", false
}
