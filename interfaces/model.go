package interfaces

type IModel interface {
	ModelError() string
	Validate() IModel
	MapTo(dest interface{}) IModel
	MapBy(src interface{}) IModel
}
