package collection

import (
	"sync"
	"reflect"
	"errors"
)

type Collection struct {
	list interface{}
}

func MakeCollection(list interface{}) *Collection {
	return &Collection{
		list,
	}
}

// KeyBy Key an associative array by a field or using a callback.
func (c *Collection) KeyBy(key string) (*Collection, error) {
	m := sync.Map{}

	switch list := c.list.(type) {
	case map[interface{}]interface{}:
	case []interface{}:
		for _, value := range list {
			ref := reflect.ValueOf(value)
			k := reflect.Indirect(ref).FieldByName(key)

			if !k.IsValid() {
				return MakeCollection(m), errors.New(key + "不存在")
			}

			m.Store(k.Interface(), value)
		}
	}
	return MakeCollection(m), nil
}

// All Get all of the items in the collection.
func (c *Collection) All() interface{} {
	return c.list
}

//func (c *Collection) operatorForWhere(key interface{}, operator string, value interface{}) func(item interface{}) bool {
//	return func(item interface{}) bool {
//		for key, value := range reflect.ValueOf(item).Kind() {
//
//		}
//		return true
//	}
//}
