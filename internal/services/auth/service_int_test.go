package auth

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/testdb"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/jwt_token_generator"
	"github.com/kudrmax/perfectPetProject/internal/services/password_hasher"
	"github.com/kudrmax/perfectPetProject/internal/services/test/fake"
	"github.com/kudrmax/perfectPetProject/internal/services/users"
)

const (
	password = "some_password"
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
		a.NotEmpty(accessToken)
		a.NoError(err)

		userFromDB := testdb.MustGetUserByUsername(a, s.db, user.Username)
		a.Equal(user.Username, userFromDB.Username)
		a.Equal(user.Name, userFromDB.Name)
		a.NotEqual(user.PasswordHash, userFromDB.PasswordHash)
		a.NotEmpty(userFromDB.Id)
	})
}

func EqualWithoutId(a *require.Assertions, user1, user2 models.User) {
	user1.Id = 0
	user2.Id = 0
	a.Equal(user1, user2)
}
