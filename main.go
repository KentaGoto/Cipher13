package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ROT13変換用の関数
func rot13(s string) string {
	rot13 := func(r rune) rune {
		switch {
		case 'a' <= r && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case 'A' <= r && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		default:
			return r
		}
	}
	return strings.Map(rot13, s)
}

func main() {
	// ルーターの作成
	router := gin.Default()

	// APIのルーティング
	router.GET("/api/rot13", func(c *gin.Context) {
		// クエリパラメータから元の文字列を取得
		original := c.Query("s")
		// 元の文字列をROT13変換
		encrypted := rot13(original)
		// 結果をJSONとして返す
		c.JSON(200, gin.H{
			"original": original,
			"rot13":    encrypted,
		})
	})

	// 静的ファイルのサーブ
	router.Static("/static", "./my-app/build/static")
	router.GET("/", func(c *gin.Context) {
		c.File("./my-app/build/index.html")
	})

	// サーバーの起動
	router.Run(":8080")
}
