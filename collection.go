package collection

import (
	"sync"
	"reflect"
	"errors"
)

type Collection struct {
	list []interface{}
}

func MakeCollection(list []interface{}) *Collection {
	return &Collection{
		list,
	}
}

// KeyBy Key an associative array by a field or using a callback.
func (c *Collection) KeyBy(key string) (sync.Map, error) {
	m := sync.Map{}
	for _, value := range c.list {
		ref := reflect.ValueOf(value)
		k := reflect.Indirect(ref).FieldByName(key)

		if !k.IsValid() {
			return m, errors.New(key + "不存在")
		}

		m.Store(k.Interface(), value)
	}
	return m, nil
}

// All Get all of the items in the collection.
func (c *Collection) All() []interface{} {
	return c.list
}
