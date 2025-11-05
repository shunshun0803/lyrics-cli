package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LyricsResponse は api.lyrics.ovh のJSONレスポンスを受け取るための構造体
type LyricsResponse struct {
	Lyrics string `json:"lyrics"`
	Error  string `json:"error"` // 404の時などに "No lyrics found" が入る
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [アーティスト名] [曲名]",
	Short: "指定されたアーティストと曲名の歌詞を取得します",
	Long: `api.lyrics.ovh から歌詞を取得するCLIツールです。
例: lyrics-cli search "Queen" "Bohemian Rhapsody"`,

	// Run: コマンドが実行されたときの処理
	// args には [アーティスト名] と [曲名] がスライスとして入ってくる
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 引数が2個（アーティスト名、曲名）であるかチェック
		if len(args) != 2 {
			fmt.Println("Error: アーティスト名と曲名を2つ指定してください")
			fmt.Println("例: go run . search \"Queen\" \"Bohemian Rhapsody\"")
			os.Exit(1)
		}

		// 2. 引数をURLエンコードする（"Bohemian Rhapsody" のようなスペースを %20 に変換するため）
		artist := url.PathEscape(args[0])
		title := url.PathEscape(args[1])

		// 3. APIのURLを組み立てる
		apiURL := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", artist, title)

		// 4. APIにリクエストを送る
		resp, err := http.Get(apiURL)
		if err != nil {
			log.Fatalf("APIへのリクエストに失敗しました: %v", err)
		}
		defer resp.Body.Close()

		// 5. レスポンスボディを読み込む
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("レスポンスボディの読み込みに失敗しました: %v", err)
		}

		// 6. JSONをパースする
		var apiResp LyricsResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			log.Fatalf("JSONのパースに失敗しました: %v", err)
		}

		// 7. レスポンスのステータスと内容をチェック
		if resp.StatusCode != 200 || (apiResp.Lyrics == "" && apiResp.Error != "") {
			// APIが404を返すか、200でも "error" フィールドにメッセージを入れてくる場合
			fmt.Println("Error: 歌詞が見つかりませんでした。")
			if apiResp.Error != "" {
				fmt.Printf("APIからのエラー: %s\n", apiResp.Error)
			}
			os.Exit(1)
		}

		// 8. 成功：歌詞を出力
		fmt.Println(apiResp.Lyrics)
	},
}

func init() {
	// rootCmd (root.goで定義) に searchCmd をサブコマンドとして追加します。
	rootCmd.AddCommand(searchCmd)
}
