# golang-time-testing

## 時刻のテストのサンプル

必要なパッケージを取得する

```bash
godep ensure
```

ディレクトリごとにテストする。

```bash
go test -v ./origin
go test -v ./di
go test -v ./stub
go test -v ./clock
go test -v ./monkey
```

```bash
go test -v ./timepkg
```
