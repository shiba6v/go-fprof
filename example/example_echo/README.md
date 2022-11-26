## Example Echo

```bash
$ cd example
$ go run example_echo/main.go
# In another tab,
$ curl localhost:1323/fuga
{"fuga":1000}
$ curl localhost:1323/hoge
{"hoge":1000}
$ curl localhost:1323/piyo
{"piyo":1000}
$ curl localhost:1323/fprof_result
FProf Result [us]
Sum           60, Max           60, Avg           60, Min           60, Count            1, L11 main.Hoge
Sum           63, Max           63, Avg           63, Min           63, Count            1, L17 main.Fuga
Sum          135, Max            6, Avg            0, Min            0, Count         1000, L28 main.Piyo
```

`main.go.original.txt`にあるような元のコードに対して以下の変更を行って、プロファイリングの準備ができた`main.go`にします。
- 開始時に`fprof.InitFProf()`を付ける
- 各関数の始めに`defer fprof.FProf()()`を付ける
- 結果を吐き出すエンドポイント`fprof.AnalizeFProfResult()`を作る。

差分確認
```
diff example/example_echo/main.go example/example_echo/main.go.original.txt 
```

Piyoでは、`fpr := fprof.FProf()`と`fpr()`で挟むことで、好きな区間を計測します。ただし、シンプルな作り故にオーバーヘッドがそれなりにあります。
