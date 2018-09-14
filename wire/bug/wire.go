package bug

import "github.com/google/go-cloud/wire"

func WireFoo() (Foo, error) {
  wire.Build(ProvideInterface, NewFoo)
  return nil, nil
}
