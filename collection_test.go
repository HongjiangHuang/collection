package collection

import (
	"testing"
	"os"
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

	for key, value := range coll.All() {
		if list[key].(People).Id != value.(People).Id {
			print("The list order is changed")
		}
	}
}
