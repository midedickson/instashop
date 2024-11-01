package token

import (
	"net/http"
	"strings"
)

func ExtractFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	if bearerTokenChars := strings.Split(bearerToken, " "); len(bearerTokenChars) == 2 {
		return bearerTokenChars[1]
	}

	return ""

}
