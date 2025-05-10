package jwt_token_generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	secretKey         = "secret_key"
	defaultExpireTime = time.Minute * 10
)

func TestService_GenerateAndParseToken(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name  string
		args  args
		token string
	}{
		{
			name: "success 1",
			args: args{
				userId: 1,
			},
		},
		{
			name: "success 123",
			args: args{
				userId: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := NewService(secretKey, defaultExpireTime)

			token, err := j.GenerateToken(tt.args.userId)
			require.NoError(t, err)
			require.True(t, len(token) > 0)

			userID, err := j.ParseToken(token)
			require.NoError(t, err)
			require.Equal(t, tt.args.userId, userID)
		})
	}
}

func TestService_ParseTokenBadToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bad token",
			args: args{
				token: "123123123123123123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := NewService(secretKey, defaultExpireTime)

			userID, err := j.ParseToken(tt.args.token)
			require.ErrorIs(t, err, InvalidTokenErr)
			require.Equal(t, 0, userID)
		})
	}
}

func TestService_ParseTokenExpire(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		token   string
		wantErr bool
	}{
		{
			name: "success 1",
			args: args{
				userId: 1,
			},
		},
		{
			name: "success 123",
			args: args{
				userId: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expireAfter := time.Millisecond * 1

			j := NewService(secretKey, expireAfter)

			token, err := j.GenerateToken(tt.args.userId)
			require.NoError(t, err)

			time.Sleep(expireAfter * 2)

			_, err = j.ParseToken(token)
			require.ErrorIs(t, err, InvalidTokenErr)
		})
	}
}
