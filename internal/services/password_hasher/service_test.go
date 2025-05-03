package password_hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestService_GenerateHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{password: "123456"},
		},
		{
			name: "empty",
			args: args{password: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				cost: bcrypt.DefaultCost,
			}
			_, err := s.GenerateHashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_CompareHashAndPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "success",
			password: "123456",
		},
		{
			name:     "empty",
			password: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				cost: bcrypt.DefaultCost,
			}
			passwordHash, err := s.GenerateHashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !s.CompareHashAndPassword(passwordHash, tt.password) {
				t.Errorf("CompareHashAndPassword() = %v, want %v", s.CompareHashAndPassword(passwordHash, tt.password), true)
			}
		})
	}
}
