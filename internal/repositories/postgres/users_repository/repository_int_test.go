package users_repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/testdb"
	"github.com/kudrmax/perfectPetProject/internal/services/test/fake"
)

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	db   *sql.DB
	self *Repository
}

func (s *testSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithCancel(s.T().Context())

	s.db = testdb.MustInit(s.T())
	s.self = New(s.db)
}

func (s *testSuite) TearDownSuite() {
	s.cancel()
	if err := s.db.Close(); err != nil {
		s.FailNow("failed to close database")
	}
}

func (s *testSuite) Test_GetByUsername() {
	s.Run("success", func() {
		a := s.Require()

		user1 := fake.User()
		user2 := fake.User()
		user3 := fake.User()
		testdb.MustAddUser(a, s.db, user1)
		testdb.MustAddUser(a, s.db, user2)
		testdb.MustAddUser(a, s.db, user3)

		got, err := s.self.GetByUsername(user2.Username)

		a.NoError(err)
		a.Equal(user2, got)
	})

	s.Run("success_not_found", func() {
		a := s.Require()

		got, err := s.self.GetByUsername(fake.RandString())

		a.NoError(err)
		a.Nil(got)
	})

	s.Run("success_empty_username", func() {
		a := s.Require()

		got, err := s.self.GetByUsername("")

		a.NoError(err)
		a.Nil(got)
	})
}

func (s *testSuite) Test_Create() {
	s.Run("success", func() {
		a := s.Require()

		user := fake.User()

		got, err := s.self.Create(user)
		userFromDB := testdb.MustGetUserByUsername(a, s.db, user.Username)

		a.NoError(err)
		a.NotEmpty(got.Id)

		user.Id = got.Id
		a.Equal(user, got)
		a.Equal(user, userFromDB)
	})

	s.Run("error_already_exists", func() {
		a := s.Require()

		user := fake.User()
		testdb.MustAddUser(a, s.db, user)

		userFromDbBefore := testdb.MustGetUserByUsername(a, s.db, user.Username)
		got, err := s.self.Create(user)
		userFromDbAfter := testdb.MustGetUserByUsername(a, s.db, user.Username)

		a.Error(err)
		a.Nil(got)
		a.Equal(ErrUsernameAlreadyExists, err)
		a.Equal(userFromDbBefore, userFromDbAfter)
	})

	s.Run("error_empty_user", func() {
		a := s.Require()

		users := []*models.User{
			func() *models.User {
				return nil
			}(),

			func() *models.User {
				user := fake.User()
				user.Username = ""
				return user
			}(),

			func() *models.User {
				user := fake.User()
				user.Name = ""
				return user
			}(),
		}

		for _, user := range users {
			got, err := s.self.Create(user)
			a.Error(err)
			a.Nil(got)
			a.Equal(ErrEmptyUser, err)
			if user != nil {
				a.False(testdb.UserExists(s.db, user.Username))
			}
		}
	})
}

func EqualWithoutId(a *require.Assertions, user1, user2 models.User) {
	user1.Id = 0
	user2.Id = 0
	a.Equal(user1, user2)
}
