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

	var err error = nil

	c.Range(func(index, item interface{}) bool {
		ref := reflect.ValueOf(item)
		k := reflect.Indirect(ref).FieldByName(key)

		if !k.IsValid() {
			err = errors.New(key + "不存在")
			return false
		}

		m.Store(k.Interface(), item)

		return true
	})

	return MakeCollection(m), err
}

// All Get all of the items in the collection.
func (c *Collection) All() interface{} {
	return c.list
}

func (c *Collection) Range(cb func(key, item interface{}) bool) {
	switch list := c.list.(type) {
	case sync.Map:
		list.Range(cb)
		break
	case []interface{}:
		for key, value := range list {
			if !cb(key, value) {
				break
			}
		}
	case map[interface{}]interface{}:
		for key, value := range list {
			if !cb(key, value) {
				break
			}
		}
		break
	default:
		panic("Unsupported types")
	}
}

//func (c *Collection) operatorForWhere(key interface{}, operator string, value interface{}) func(item interface{}) bool {
//	return func(item interface{}) bool {
//		for key, value := range reflect.ValueOf(item).Kind() {
//
//		}
//		return true
//	}
//}
