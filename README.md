# go-lazy-load-pattern

[日本語](./README.ja.md)

Implementation pattern to achieve the following:

* Lazy loading
* In-memory caching
* In-memory indexing

# Applicable Cases

* When translating a desired operation directly into a program, multiple database accesses occur, leading to performance issues
    * -> Consolidate multiple database accesses into one and retain in in-memory cache
    * -> Create an index (associative array) in-memory to speed up array searches
* When ORM's Preload or Eager Loading doesn't solve the issue
* When you do not want to access the database if it's outside the output target due to filtering functions
    * -> Utilize lazy loading to avoid database access when it's outside the output target

# Implementation Example

```go
type Record struct {
	Id   string
	Name string
	Age  int
}

func lazyLoadRecords() (getter.IGetter[string, *Record], getter.IGetter[int, *Record]) {
	l := loader.NewLoader(func() ([]*Record, error) {
		var records []*Record
		db.Find($records)
		return records, nil
	})
	h := loader.NewCacheHolder[*Record](l)
	recordsByName := getter.NewIndexer[*Record](h, func(record *Record) (string, *Record, error) {
		return record.Name, record, nil
	})
	recordsByAge := getter.NewArrayIndexer[*Record](h, func(record *Record) (int, *Record, error) {
		return record.Age, record, nil
	})

	return recordsByName, recordsByAge
}

func main() {
	recordsByName, recordsByAge := lazyLoadRecords()
	
	// do various things here

	record, ok, err := recordsByName.Get("John")
	println(record)
	// show record that name is John

	records, ok, err := recordsByAge.Get(20)
	println(records)
	// show multiple records that age is 20
	
	// No matter how many times Get is called, 
	// DB access will only occur once.
}
```