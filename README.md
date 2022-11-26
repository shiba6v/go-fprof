# FProf
## About
FProfは、シンプルな関数レベルのプロファイリングツールです。
FProf is a simple function level profiling tool.

## Install
```bash
go get github.com/shiba6v/go-fprof@main
```

## Usage
- 開始時に`fprof.InitFProf()`を追加
- 各関数の始めに`defer fprof.FProf()()`を追加
- `fprof.AnalizeFProfResult()`でプロファイリング結果を出力

```go
func A() {
	defer fprof.FProf()() // 追加。各関数の始めに付ける。
	// some process
	time.Sleep(1 * time.Second)
}

func main() {
	fprof.InitFProf() // 追加。開始時の初期化。
	A()
	fmt.Println(fprof.AnalizeFProfResult()) // 追加。プロファイリング結果を出力。
}
```

また、`fpr := fprof.FProf()`と`fpr()`で挟むと、好きな区間を計測することもできます。

基本的な使い方は、 [Example Goroutine](https://github.com/shiba6v/go-fprof/tree/main/example/example_goroutine) を参照してください。
ISUCONなど、サーバーで使う場合は [Example Echo](https://github.com/shiba6v/go-fprof/tree/main/example/example_echo) を参照してください。

## Disclaimer
破壊的変更を入れる際は新しいバージョンのディレクトリを切ろうと思いますが、ISUCONなどでの使用時に挙動が変わっても責任は取れません。
そのため、コミットハッシュを指定しての使用や、forkしての使用を推奨します。

## Future Plan
- `defer fprof.FProf()()`の追加を自動化したいと思っています。
  - ASTを使ってビルド時に埋め込むのも考えられるが、任意のビルドプロセスに適合するかは怪しいです。
  - fprof.bashやFPROF_IGNOREを使って自動化を試みた残骸が残っています。
