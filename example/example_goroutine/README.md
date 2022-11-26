## Example Goroutine
```
$ cd example
$ go run example_goroutine/main.go 
100
FProf Result [us]
Sum          150, Max          150, Avg          150, Min          150, Count            1, L13 main.A
Sum           11, Max            1, Avg            0, Min            0, Count          100, L23 main.B
```

`main.go.original.txt`にあるような元のコードに対して以下の変更を行って、プロファイリングの準備ができた`main.go`にします。
- 開始時に`fprof.InitFProf()`を付ける
- 各関数の始めに`defer fprof.FProf()()`を付ける
- 終了時に`r := fprof.AnalizeFProfResult(); fmt.Println(r)`でプロファイリング結果を出力させる。

差分確認
```
diff example/example_goroutine/main.go example/example_goroutine/main.go.original.txt 
```
