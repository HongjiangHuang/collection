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

// Range Execute a callback over each item.
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

// Filter Run a filter over each of the items.
func (c *Collection) Filter(f func(key, value interface{}) bool) *Collection {
	m := sync.Map{}

	c.Range(func(index, item interface{}) bool {

		if f(index, item) {
			m.Store(index, item)
		}

		return true
	})

	return MakeCollection(m)
}

// Where Filter items by the given key value pair.
func (c *Collection) Where(key, operator string, value interface{}) *Collection {
	return c.Filter(c.operatorForWhere(key, operator, value))
}

// operatorForWhere Get an operator checker callback.
func (c *Collection) operatorForWhere(key, operator string, value interface{}) func(key, item interface{}) bool {
	return func(index, item interface{}) bool {

		returnVal := true

		ref := reflect.ValueOf(item)
		k := ref.FieldByName(key)

		if !k.IsValid() {
			return false
		}

		switch v := value.(type) {
		case int:
			switch operator {
			case "=", "==":
				return int64(v) == k.Int()
			case "!=", "<>":
				return int64(v) != k.Int()
			case "<":
				return int64(v) < k.Int()
			case ">":
				returnVal = int64(v) > k.Int()
				return returnVal
			case ">=":
				return int64(v) >= k.Int()
			case "<=":
				return int64(v) <= k.Int()
			default:
				return false
			}
		case string:
			switch operator {
			case "=", "==":
				return v == k.String()
			case "!=", "<>":
				return v != k.String()
			default:
				return false
			}
		default:
			panic("Unsupported types")
		}

		return true
	}
}
