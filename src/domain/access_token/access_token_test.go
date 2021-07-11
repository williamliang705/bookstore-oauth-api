package access_token

import "testing"

func TestGetNewsAccessToken(t *testing.T) {
	at := GetNewsAccessToken()
	if at.IsExpired() {
		t.Error("brand news access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("news access token should not have defined access token id")
	}

	if at.UserId != 0 {
		t.Error("news access token should not have an associated user id")
	}
}