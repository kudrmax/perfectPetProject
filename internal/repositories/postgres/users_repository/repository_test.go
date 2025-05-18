package users_repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/testdb"
	"github.com/kudrmax/perfectPetProject/internal/services/test/fake"
)

func TestRepository_GetByUsername(t *testing.T) {
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
