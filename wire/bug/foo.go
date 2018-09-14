package bug

type Interface interface{}

type Foo interface{}

func NewFoo(i Interface) (Foo, error) {
  return nil, nil
}

func ProvideInterface() Interface {
  return nil
}
