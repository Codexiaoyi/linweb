package interfaces

// IModel is used to represent a struct that needs to be processed,
// it uses chain call, can call continuously and finally output ModelError.
type IModel interface {
	// New create a new model
	New(model interface{}) IModel
	// Validate follow the rules to validate this model.
	Validate() error
	// MapToByFieldName map this model to target model by field name.
	MapToByFieldName(dest interface{}) error
	// MapToByFieldTag map this model to target model by field tag.
	MapToByFieldTag(dest interface{}) error
	// MapFromByFieldName this model is map from source model by field name.
	MapFromByFieldName(src interface{}) error
	// MapFromByFieldTag this model is map from source model by field tag.
	MapFromByFieldTag(src interface{}) error
}
