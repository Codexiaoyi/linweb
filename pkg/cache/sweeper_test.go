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
	suite.sweeper.addExpireKey("key1", 1*time.Second)
	suite.sweeper.addExpireKey("key2", 2*time.Second)
	suite.sweeper.addExpireKey("key3", 3*time.Second)
	suite.sweeper.addExpireKey("key4", 4*time.Second)
}

func (suite *SweeperTestSuite) TestSweep() {
	assert.Equal(suite.T(), 4, len(suite.sweeper.expireMap))
	time.Sleep(6 * time.Second)
	assert.Equal(suite.T(), 1, len(suite.sweeper.expireMap))
}

func TestSweeperTestSuite(t *testing.T) {
	suite.Run(t, new(SweeperTestSuite))
}
