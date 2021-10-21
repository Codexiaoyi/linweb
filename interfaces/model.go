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

// IModel is used to represent a struct that needs to be processed,
// it uses chain call, can call continuously and finally output ModelError.
type IModel interface {
	// New create a new model
	New(model interface{}) IModel
	// Validate follow the rules to validate this model.
	Validate() IModel
	// MapToByFieldName map this model to target model by field name.
	MapToByFieldName(dest interface{}) IModel
	// MapToByFieldTag map this model to target model by field tag.
	MapToByFieldTag(dest interface{}) IModel
	// MapFromByFieldName this model is map from source model by field name.
	MapFromByFieldName(src interface{}) IModel
	// MapFromByFieldTag this model is map from source model by field tag.
	MapFromByFieldTag(src interface{}) IModel
	// model chain error
	ModelError() error
}
