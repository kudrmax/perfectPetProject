package users_repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/storage"
	"github.com/kudrmax/perfectPetProject/internal/services/test/fake"
	"github.com/kudrmax/perfectPetProject/internal/services/test/testdb"
)

func TestRepository_GetByUsername(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	storage    *storage.Storage // TODO заменить storage на db
	repository *Repository
}

func (s *testSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithCancel(s.T().Context())

	s.storage = testdb.MustInit(s.Assert())
	s.repository = New(s.storage)
}

func (s *testSuite) TearDownSuite() {
	s.cancel()
	if err := s.storage.Close(); err != nil {
		s.FailNow("failed to close database")
	}
}

func (s *testSuite) Test_GetByUsername() {
	s.Run("success", func() {
		a := s.Require()

		user1 := fake.User()
		user2 := fake.User()
		user3 := fake.User()
		testdb.MustAddUser(a, s.storage.DB, user1)
		testdb.MustAddUser(a, s.storage.DB, user2)
		testdb.MustAddUser(a, s.storage.DB, user3)

		got, err := s.repository.GetByUsername(user2.Username)

		a.NoError(err)
		a.Equal(user2, got)
	})

	s.Run("success_not_found", func() {
		a := s.Require()

		got, err := s.repository.GetByUsername(fake.RandString())

		a.NoError(err)
		a.Nil(got)
	})

	s.Run("success_empty_username", func() {
		a := s.Require()

		got, err := s.repository.GetByUsername("")

		a.NoError(err)
		a.Nil(got)
	})
}
