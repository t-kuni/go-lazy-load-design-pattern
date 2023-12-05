# go-lazy-load-pattern

以下を実現する実装パターン

* 遅延ロード
* インメモリキャッシュ
* インメモリインデキシング

# 活用できるケース

* やりたい事を素直にプログラムに落とし込むと複数回のDBアクセスが発生しパフォーマンスの問題につながる場合
  * -> 複数回のDBアクセスを1度にまとめてインメモリキャッシュに保持する
  * -> インメモリでインデックス（連想配列）を作成し、配列の検索を高速化する
* ORMのPreloadやEager Loadingで解決できない場合
* フィルタリングの機能により、出力対象外の場合はDBアクセスを行いたくない場合
  * -> 遅延ロードにより、出力対象外の場合はDBアクセスを行わないようにする

# 実装例

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