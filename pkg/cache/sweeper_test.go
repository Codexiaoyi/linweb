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
	suite.sweeper = newSweeper(100*time.Microsecond, func(key string) {
		println(key)
	})
}

func (suite *SweeperTestSuite) TestSweep() {
	suite.sweeper.addExpireKey("key1", 10*time.Microsecond)
	assert.Equal(suite.T(), suite.sweeper.isSweeping, true)
	assert.Equal(suite.T(), 1, len(suite.sweeper.expireMap))
	time.Sleep(1200 * time.Microsecond)
	assert.Equal(suite.T(), 0, len(suite.sweeper.expireMap))
	assert.Equal(suite.T(), suite.sweeper.isSweeping, false)
}

func TestSweeperTestSuite(t *testing.T) {
	suite.Run(t, new(SweeperTestSuite))
}
