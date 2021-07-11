package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstans(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewsAccessToken(t *testing.T) {
	at := GetNewsAccessToken()
	assert.False(t, at.IsExpired(), "brand news access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "news access token should not have defined access token id")

	assert.True(t, at.UserId == 0, "news access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token created three hours from now should Not be expired")
}
