package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gopkg.in/go-playground/assert.v1"
)

type SweeperTestSuite struct {
	suite.Suite
	sweeper *sweeper
}

func (suite *SweeperTestSuite) SetupTest() {
	suite.sweeper = newSweeper(3*time.Second, func(key string) {
		println(key)
	})
	suite.sweeper.addExpireKey("key2", 2*time.Second)
	suite.sweeper.addExpireKey("key6", 6*time.Second)
	suite.sweeper.addExpireKey("key9", 9*time.Second)
}

func (suite *SweeperTestSuite) TestSweep() {
	assert.Equal(suite.T(), 3, len(suite.sweeper.expireMap))
	time.Sleep(5 * time.Second)
	assert.Equal(suite.T(), 2, len(suite.sweeper.expireMap))
	time.Sleep(2 * time.Second)
	assert.Equal(suite.T(), 1, len(suite.sweeper.expireMap))
	time.Sleep(3 * time.Second)
	assert.Equal(suite.T(), 0, len(suite.sweeper.expireMap))
}

func TestSweeperTestSuite(t *testing.T) {
	suite.Run(t, new(SweeperTestSuite))
}
