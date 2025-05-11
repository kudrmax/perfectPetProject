package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/services/auth/mocks"
)

type MockFactory struct {
	jwtMock                   *mocks.MockjwtProviderService
	userServiceMock           *mocks.MockuserService
	passwordHasherServiceMock *mocks.MockpasswordHasherService
}

func createMocks(t *testing.T) *MockFactory {
	t.Helper()

	ctrl := gomock.NewController(t)
	return &MockFactory{
		jwtMock:                   mocks.NewMockjwtProviderService(ctrl),
		userServiceMock:           mocks.NewMockuserService(ctrl),
		passwordHasherServiceMock: mocks.NewMockpasswordHasherService(ctrl),
	}
}

func TestService_Register(t *testing.T) {
	const (
		name         = "Max Kudryashov"
		username     = "kudrmax"
		password     = "qwerty"
		passwordHash = "qwertyHash"
		token        = "someToken"
	)

	type args struct {
		name     string
		username string
		password string
	}
	tests := []struct {
		name            string
		args            args
		setMockBehavior func(m *MockFactory)
		wantAccessToken string
		wantErr         bool
	}{
		{
			name: "success",
			args: args{
				name:     name,
				username: username,
				password: password,
			},
			wantAccessToken: token,
			setMockBehavior: func(m *MockFactory) {
				m.userServiceMock.EXPECT().
					GetByUsername(username).
					Return(nil)
				m.passwordHasherServiceMock.EXPECT().
					GenerateHashPassword(password).
					Return(passwordHash, nil)
				m.userServiceMock.EXPECT().
					Create(
						&models.User{
							Name:         name,
							Username:     username,
							PasswordHash: passwordHash,
						},
					).
					Return(
						&models.User{
							Id:           1,
							Name:         name,
							Username:     username,
							PasswordHash: passwordHash,
						}, nil,
					)
				m.jwtMock.EXPECT().GenerateToken(1).Return(token, nil)
			},
		},
		{
			name: "error UserAlreadyExistsErr",
			args: args{
				name:     name,
				username: username,
				password: password,
			},
			wantErr: true,
			setMockBehavior: func(m *MockFactory) {
				m.userServiceMock.EXPECT().
					GetByUsername(username).
					Return(
						&models.User{
							Name:         name,
							Username:     username,
							PasswordHash: passwordHash,
						},
					)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockFactory := createMocks(t)
			tt.setMockBehavior(mockFactory)

			s := NewService(
				mockFactory.userServiceMock,
				mockFactory.jwtMock,
				mockFactory.passwordHasherServiceMock,
			)

			gotAccessToken, err := s.Register(tt.args.name, tt.args.username, tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.wantAccessToken, gotAccessToken)
		})
	}
}

func TestService_Login(t *testing.T) {
	const (
		name         = "Max Kudryashov"
		username     = "kudrmax"
		password     = "qwerty"
		passwordHash = "qwertyHash"
		token        = "someToken"
	)

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name            string
		args            args
		wantAccessToken string
		wantErr         bool
		setMockBehavior func(m *MockFactory)
	}{
		{
			name: "success",
			args: args{
				username: username,
				password: password,
			},
			wantAccessToken: token,
			setMockBehavior: func(m *MockFactory) {
				m.userServiceMock.EXPECT().
					GetByUsername(username).
					Return(
						&models.User{
							Id:           1,
							Name:         name,
							Username:     username,
							PasswordHash: passwordHash,
						},
					)
				m.passwordHasherServiceMock.EXPECT().
					CompareHashAndPassword(passwordHash, password).
					Return(true)
				m.jwtMock.EXPECT().
					GenerateToken(1).
					Return(token, nil)
			},
		},
		{
			name: "error UserNotFoundErr",
			args: args{
				username: username,
				password: password,
			},
			wantErr: true,
			setMockBehavior: func(m *MockFactory) {
				m.userServiceMock.EXPECT().
					GetByUsername(username).
					Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockFactory := createMocks(t)
			tt.setMockBehavior(mockFactory)

			s := NewService(
				mockFactory.userServiceMock,
				mockFactory.jwtMock,
				mockFactory.passwordHasherServiceMock,
			)

			gotAccessToken, err := s.Login(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAccessToken != tt.wantAccessToken {
				t.Errorf("Login() gotAccessToken = %v, want %v", gotAccessToken, tt.wantAccessToken)
			}
		})
	}
}

func TestService_ValidateTokenAndGetUserId(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name       string
		args       args
		wantUserId int
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				token: "some_good_token",
			},
			wantUserId: 1,
			wantErr:    false,
		},
		{
			name: "err",
			args: args{
				token: "some_bad_token",
			},
			wantUserId: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		setMockBehavior := func(m *MockFactory, wantErr bool) {
			if !wantErr {
				m.jwtMock.EXPECT().ParseToken(tt.args.token).Return(tt.wantUserId, nil)
				return
			}
			m.jwtMock.EXPECT().ParseToken(tt.args.token).Return(0, errors.New("some error"))
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			mockFactory := createMocks(t)
			setMockBehavior(mockFactory, tt.wantErr)

			s := NewService(nil, mockFactory.jwtMock, nil)

			gotUserId, err := s.ValidateTokenAndGetUserId(tt.args.token)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTokenAndGetUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.wantUserId, gotUserId)
		})
	}
}
