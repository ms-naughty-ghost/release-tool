package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	toolVersion := "1.0.0"
	// コマンドラインオプションを定義
	var (
		// -hはヘルプオプション
		// -v オプション このツールのバージョン
		version = flag.Bool("v", false, "このツールのバージョンを表示する")
		// -c オプション CHANGELOG.md のパス
		filePath = flag.String("c", "CHANGELOG.md", "CHANGELOG.md のパス")
		// -latest オプション latestバージョンを表示する
		latestVersion = flag.Bool("latest", true, "Latestバージョンを表示する")
		// -l オプション リリースバージョンのリスト
		versionList = flag.Bool("l", false, "リリースバージョンのリストを表示する")
		// -r オプション 指定したリリースバージョンの変更履歴を表示する
		releaseVersion = flag.String("r", "", "指定したリリースバージョンの変更履歴を表示する")
	)
	// コマンドライン引数を取得
	flag.Parse()
	// バージョンを表示(コマンドラインオプション -v)
	if *version {
		println(toolVersion)
		os.Exit(0)
	}
	// filePathを読み込む
	buf, e := os.OpenFile(*filePath, os.O_RDONLY, 0)
	if e != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", e)
		os.Exit(1)
	}
	defer buf.Close()
	
	// リリースバージョンのリストを表示する(コマンドラインオプション -l)
	if *versionList {
		rex := regexp.MustCompile(`## \[.+\]`)
		replace := regexp.MustCompile(`## \[|\]`)
		// リリースバージョンのリストを表示する
		fr := bufio.NewReader(buf)
		for {
			line, _, err := fr.ReadLine()
			if err != nil {
				break
			}
			if rex.Match(line) {
				verLine := rex.FindString(string(line))
				version := replace.ReplaceAllString(string(verLine), "")
				if version != "Unreleased" {
					fmt.Println(version)
				}
			}
		}
		os.Exit(0)
	}
	// 指定したリリースバージョンの変更履歴を表示する(コマンドラインオプション -r)
	if *releaseVersion != "" {
		rex := regexp.MustCompile(`## \[.+\]`)
		target := regexp.MustCompile(`## \[` + *releaseVersion + `\]`)
		// 指定したリリースバージョンの変更履歴を表示する
		fr := bufio.NewReader(buf)
		for {
			line, _, err := fr.ReadLine()
			if err != nil {
				break
			}
			if target.Match(line) {
				fmt.Println(string(line))
				for {
					line, _, err := fr.ReadLine()
					if err != nil {
						break
					}
					if rex.Match(line) {
						break
					}
					fmt.Println(string(line))
				}
				// 指定したリリースバージョンの変更履歴を表示したので、終了する
				break
			}
		}
		os.Exit(0)
	}
	// latestバージョンを表示する(コマンドラインオプション -latest)
	if *latestVersion {
		rex := regexp.MustCompile(`## \[.+\]`)
		replace := regexp.MustCompile(`## \[|\]`)
		// latestバージョンを表示する
		fr := bufio.NewReader(buf)
		for {
			line, _, err := fr.ReadLine()
			if err != nil {
				break
			}
			if rex.Match(line) {
				verLine := rex.FindString(string(line))
				version := replace.ReplaceAllString(string(verLine), "")
				if version != "Unreleased" {
					fmt.Println(version)
					// latestバージョンのみ表示するので、最初の1件で終了する
					break
				}
			}
		}
		os.Exit(0)
	}
	
}