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

package interfaces

import "reflect"

// The DI container, regist object in the linweb init.
type IInjector interface {
	// Add singleton objects to the container.
	AddSingleton(objects ...interface{})
	// Add Transient objects to the container, create a new object in every reuqest.
	AddTransient(objects ...interface{})
	// Inject object in the request is coming.
	Inject(value reflect.Value)
}
