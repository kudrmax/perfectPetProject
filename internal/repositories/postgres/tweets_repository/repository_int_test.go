package tweets_repository

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

func (s *testSuite) Test_GetAll() {
	addNewFakeUserAndTweet := func(a *require.Assertions) (*models.Tweet, *models.User) {
		user := fake.User()
		testdb.MustAddUser(a, s.db, user)

		tweet := fake.Twit().SetUser(user)
		testdb.MustAddTweet(a, s.db, tweet)

		return tweet, user
	}

	s.Run("success_one_element", func() {
		a := s.Require()

		testdb.MustDeleteAllTweets(a, s.db)

		tweet, _ := addNewFakeUserAndTweet(a)

		twits, err := s.self.GetAll()

		a.NoError(err)
		a.Len(twits, 1)
		a.Equal(tweet.UserId, twits[0].UserId)
		a.Equal(tweet.Text, twits[0].Text)
	})

	s.Run("success_multiple_element", func() {
		a := s.Require()

		testdb.MustDeleteAllTweets(a, s.db)

		addNewFakeUserAndTweet(a)
		addNewFakeUserAndTweet(a)
		addNewFakeUserAndTweet(a)

		twits, err := s.self.GetAll()

		a.NoError(err)
		a.Len(twits, 3)
	})

	s.Run("success_empty", func() {
		a := s.Require()

		testdb.MustDeleteAllTweets(a, s.db)

		twits, err := s.self.GetAll()

		a.NoError(err)
		a.Empty(twits)
	})
}

func (s *testSuite) Test_Create() {
	addNewUserAndGetTweet := func(a *require.Assertions) (*models.Tweet, *models.User) {
		user := fake.User()
		testdb.MustAddUser(a, s.db, user)

		tweet := fake.Twit()
		tweet.UserId = user.Id
		tweet.User = nil

		return tweet, user
	}

	s.Run("success", func() {
		a := s.Require()

		tweet, _ := addNewUserAndGetTweet(a)

		got, err := s.self.Create(tweet)

		a.NoError(err)
		a.Equal(got, tweet)
		a.Nil(got.User)
	})

	s.Run("error_there_is_no_user", func() {
		a := s.Require()

		user := fake.User()
		tweet := fake.Twit()
		tweet.UserId = user.Id

		_, err := s.self.Create(tweet)

		a.Error(err)
		// TODO добавить ErrorIs

		// TODO добавить проверку что в БД не добавилось
	})
}
