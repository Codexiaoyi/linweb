package linweb

import (
	"linweb/interfaces"
	"linweb/pkg/context"
	"linweb/pkg/model"
	"linweb/pkg/router"
)

func NewLinweb() *Linweb {
	return &Linweb{router: NewRouter(), context: NewContext()}
}

func NewRouter() interfaces.IRouter {
	return router.New()
}

func NewContext() interfaces.IContext {
	return &context.Context{}
}

func NewModel(modelStruct interface{}) interfaces.IModel {
	return model.New(modelStruct)
}
