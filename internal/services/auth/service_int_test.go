package auth

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/testdb"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/jwt_token_generator"
	"github.com/kudrmax/perfectPetProject/internal/services/password_hasher"
	"github.com/kudrmax/perfectPetProject/internal/services/test/fake"
	"github.com/kudrmax/perfectPetProject/internal/services/users"
)

const (
	password        = "some_password"
	anotherPassword = "some_password_another"
)

func TestServiceAuth(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	db   *sql.DB
	self *Service
}

func (s *testSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithCancel(s.T().Context())

	s.db = testdb.MustInit(s.T())
	s.self = NewService(
		users.NewService(users_repository.New(s.db)),
		jwt_token_generator.NewService("some_secret", time.Minute*15),
		password_hasher.NewService(),
	)
}

func (s *testSuite) TearDownSuite() {
	s.cancel()
	if err := s.db.Close(); err != nil {
		s.FailNow("failed to close database")
	}
}

func (s *testSuite) Test_Register() {
	s.Run("success", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())

		accessToken, err := s.self.Register(
			user.Name,
			user.Username,
			password,
		)
		a.NoError(err)
		a.NotEmpty(accessToken)

		userFromDB := testdb.MustGetUserByUsername(a, s.db, user.Username)
		a.Equal(user.Username, userFromDB.Username)
		a.Equal(user.Name, userFromDB.Name)
		a.NotEqual(user.PasswordHash, userFromDB.PasswordHash)
		a.NotEmpty(userFromDB.Id)
	})

	s.Run("error_user_already_exists", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())

		testdb.MustAddUser(a, s.db, user)

		accessToken, err := s.self.Register(
			user.Name,
			user.Username,
			password,
		)

		a.Empty(accessToken)
		a.ErrorIs(err, UserAlreadyExistsErr)
	})
}

func (s *testSuite) Test_Login() {
	s.Run("success", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())

		_, err := s.self.Register(user.Name, user.Username, password)
		a.NoError(err)

		accessTokenFromLogin, err := s.self.Login(user.Username, password)
		a.NoError(err)
		a.NotEmpty(accessTokenFromLogin)
	})

	s.Run("error_not_found", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())

		accessTokenFromLogin, err := s.self.Login(user.Username, password)
		a.ErrorIs(err, UserNotFoundErr)
		a.Empty(accessTokenFromLogin)
	})

	s.Run("error_wrong_password", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())
		_, err := s.self.Register(user.Name, user.Username, password)
		a.NoError(err)

		accessTokenFromLogin, err := s.self.Login(user.Username, anotherPassword)
		a.ErrorIs(err, WrongPasswordErr)
		a.Empty(accessTokenFromLogin)
	})
}

func (s *testSuite) Test_ValidateTokenAndGetUserId() {
	s.Run("success", func() {
		a := s.Require()

		user := fake.User(fake.WithoutPasswordHash())

		accessToken, err := s.self.Register(user.Name, user.Username, password)
		a.NoError(err)
		a.NotEmpty(accessToken)

		user = testdb.MustGetUserByUsername(a, s.db, user.Username)

		userID, err := s.self.ValidateTokenAndGetUserId(accessToken)
		a.NoError(err)
		a.Equal(user.Id, userID)
	})

	s.Run("error_invalid_token", func() {
		a := s.Require()

		userID, err := s.self.ValidateTokenAndGetUserId("some_other_token")
		a.Error(err)
		a.Empty(userID)
	})
}
