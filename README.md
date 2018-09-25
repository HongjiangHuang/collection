# collection api

### Make by sync.Map

```$xslt
m := sync.Map{}
c := MakeCollection(m)
```

### Make by slice

```$xslt
type People struct {
	Id   int
	Name string
	Age  int
}

list := []interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
coll := MakeCollection(list)
```

### Make by map

```$xslt
type People struct {
	Id   int
	Name string
	Age  int
}
m := make(map[int]interface{})
m[0] = People{Id: 1, Name: "Albert", Age: 18}
coll := MakeCollection(m)
```

### Make by array

```$xslt
type People struct {
	Id   int
	Name string
	Age  int
}

list := [2]interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
coll := MakeCollection(list)
```

### KeyBy

```$xslt
type People struct {
	Id   int
	Name string
	Age  int
}

list := []interface{}{People{Id: 1, Name: "Albert", Age: 18}, People{Id: 2, Name: "Albert1", Age: 13}}
coll := MakeCollection(list)
coll, _ := coll.KeyBy("Id")

i := m.All()
m := i.(sync.Map)
people, _ := m.Load(1)
fmt.Print(people.(People).Name)   //Albert
```