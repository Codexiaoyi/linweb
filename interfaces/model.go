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
