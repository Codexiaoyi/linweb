//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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
	suite.sweeper = newSweeper(100*time.Millisecond, func(key string) {
		println(key)
	})
}

func (suite *SweeperTestSuite) TestSweep() {
	suite.sweeper.addExpireKey("key1", 10*time.Millisecond)
	suite.sweeper.addExpireKey("key2", 700*time.Millisecond)
	time.Sleep(600 * time.Millisecond)
	assert.Equal(suite.T(), suite.sweeper.isSweeping, true)
	assert.Equal(suite.T(), 1, suite.sweeper.mLength)
	time.Sleep(300 * time.Millisecond)
	assert.Equal(suite.T(), suite.sweeper.isSweeping, false)
	assert.Equal(suite.T(), 0, suite.sweeper.mLength)
}

func TestSweeperTestSuite(t *testing.T) {
	suite.Run(t, new(SweeperTestSuite))
}
