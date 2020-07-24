// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"github.com/go-zs/go-demo/di/wiredemo/constructor"
	"github.com/google/wire"
)

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) constructor.Event {
	wire.Build(constructor.NewEvent, constructor.NewGreeter, constructor.NewMessage)
	return constructor.Event{} //返回值没有实际意义，只需符合函数签名即可
}
