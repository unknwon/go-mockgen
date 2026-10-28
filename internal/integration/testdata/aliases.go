package testdata

import "context"

type AnyReference interface {
	Get(key string) ([]byte, error)
	Set(key string, value any) error // any alias
}

type Foo struct {
	ID   string
	Name string
}

type Foos = []Foo

type Bar struct {
	ID   string
	Name string
}

type Bars Bar

type AliasReference interface {
	GetFoos() (Foos, error)
	GetBars() (Bars, error)
}

// AliasedReference is an alias of a local interface which must not be mocked twice.
type AliasedReference = AnyReference

// AliasedContext is an alias of a foreign interface which must not be mocked at all.
type AliasedContext = context.Context
