# Release Tool

## Description
CHANGELOG.mdからリリース情報を取得するコマンドラインツールです。
変更履歴の書式は、https://keepachangelog.com/ja/1.0.0/に対応します。
## Usage
```
$ release-tool -h
Usage of release-tool:
  -c string
        CHANGELOG.mdのパス (default "CHANGELOG.md")
  -latest リリース最終バージョンを表示(デフォルト動作)
  -l リリースバージョンの一覧を表示
  -r string
        指定したバージョンのリリース情報を表示
  -v ツールバージョンを表示
```

## Build
```
cd src && go build -o ../release-tool
```