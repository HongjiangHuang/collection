package collection

import (
	"testing"
	"os"
	"sync"
)

type People struct {
	Id   int
	Name string
	Age  int
}

func TestKeyBy(t *testing.T) {
	list := []interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
	coll := MakeCollection(list)
	m, e := coll.KeyBy("Id")

	if e != nil {
		panic(e)
	}

	t.Run("test key by", func(t *testing.T) {
		i := m.All()
		m := i.(sync.Map)
		people, _ := m.Load(1)
		if (people).(People).Name != "Albert" {
			print("The people name with Id 1 should be Albert")
			os.Exit(2)
		}
	})
}

func TestAll(t *testing.T) {
	list := []interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
	coll := MakeCollection(list)

	for key, value := range coll.All().([]interface{}) {
		if list[key].(People).Id != value.(People).Id {
			print("The list order is changed")
			os.Exit(2)
		}
	}
}

func TestRange(t *testing.T) {
	list := []interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
	coll, _ := MakeCollection(list).KeyBy("Id")

	coll.Range(func(key, item interface{}) bool {
		if key != item.(People).Id {
			print("The list order is changed")
			os.Exit(2)
		}
		return true
	})
}
