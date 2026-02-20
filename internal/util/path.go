package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetProjectRoot(skip int, targetFileName string) string {
	// 1. 取得目前這行程式碼所在的檔案路徑
	// runtime.Caller 為動態路徑定位
	// skip = 0：代表"當前這行代碼"所在的檔案。不論誰呼叫它，它回傳的永遠是 path.go 的路徑。
	// skip = 1：代表"呼叫者"的檔案。也就是如果你在 config.go 呼叫了這個函式，傳入 1 就會拿到 config.go 的路徑。
	// skip = 2：代表"呼叫者的呼叫者"的檔案，以此類推。
	_, b, _, _ := runtime.Caller(skip)
	dir := filepath.Dir(b)

	// 2. 向上層尋找，直到找到目標檔案
	for {
		// os.Stat 檢查目標路徑下是否有目標檔案，如果沒有會報錯誤並且繼續向上層尋找
		if _, err := os.Stat(filepath.Join(dir, targetFileName)); err == nil {
			// 沒有報錯代表找到目標檔案，回傳檔案路徑
			return dir
		}

		// 向上一層
		parent := filepath.Dir(dir)
		// 邊界條件，當此層與上一層相同時，代表已經到最上層
		if parent == dir {
			// 	已經爬到系統頂層(例如 \ 或 C:\)，仍然找不到目標檔案
			break
		}

		// 為找到目標檔案並且未到最上層，繼續向上尋找目標檔案
		dir = parent
	}

	// 3. 如果找不到目標檔案，回傳當前執行目錄
	cwd, _ := os.Getwd()
	return cwd
}
