package interfaces

// IModel is used to represent a struct that needs to be processed,
// it uses chain call, can call continuously and finally output ModelError.
type IModel interface {
	// // create a new model
	// New(model interface{}) IModel
	// return model's error, should be called at the end of the call chain.
	ModelError() error
	// follow the rules to validate this model.
	Validate() IModel
	// map this model to target model by field name.
	MapToByFieldName(dest interface{}) IModel
	// map this model to target model by field tag.
	MapToByFieldTag(dest interface{}) IModel
	// this model is map from source model by field name.
	MapFromByFieldName(src interface{}) IModel
	// this model is map from source model by field tag.
	MapFromByFieldTag(src interface{}) IModel
}
