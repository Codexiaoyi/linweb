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

package injector

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestObject struct {
	testString string
}

func (to *TestObject) Test() string {
	return to.testString
}

type InjectorTestSuite struct {
	suite.Suite
	injector *Injector
}

func (suite *InjectorTestSuite) SetupTest() {
	suite.injector = newInjector()
}

func (suite *InjectorTestSuite) TestAddSingleton() {
	type testService1 struct {
		TI *TestObject
	}
	type testService2 struct {
		TI *TestObject
	}
	ts1 := &testService1{}
	ts2 := &testService2{}
	suite.injector.AddSingleton(&TestObject{testString: "test string"})
	suite.injector.Inject(reflect.ValueOf(ts1))
	suite.NotNil(ts1.TI)
	suite.Equal("test string", ts1.TI.Test())
	suite.injector.Inject(reflect.ValueOf(ts2))
	suite.NotNil(ts2.TI)
	suite.Equal("test string", ts2.TI.Test())
	//ts2 and ts1 use same TI.
	ts2.TI.testString = "modify string"
	suite.Equal("modify string", ts1.TI.Test())
	suite.Equal("modify string", ts2.TI.Test())
}

func (suite *InjectorTestSuite) TestAddTransient() {
	type testService1 struct {
		TI *TestObject
	}
	type testService2 struct {
		TI *TestObject
	}
	ts1 := &testService1{}
	ts2 := &testService2{}
	suite.injector.AddTransient(&TestObject{testString: "test string"})
	suite.injector.Inject(reflect.ValueOf(ts1))
	suite.NotNil(ts1.TI)
	suite.Equal("test string", ts1.TI.Test())
	suite.injector.Inject(reflect.ValueOf(ts2))
	suite.NotNil(ts2.TI)
	suite.Equal("test string", ts2.TI.Test())
	//ts2 and ts1 use different TI.
	ts2.TI.testString = "modify string"
	suite.Equal("test string", ts1.TI.Test())
	suite.Equal("modify string", ts2.TI.Test())
}

func TestInjectorTestSuite(t *testing.T) {
	suite.Run(t, new(InjectorTestSuite))
}
