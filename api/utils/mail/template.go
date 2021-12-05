package mail

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/cateiru/cateiru-sso/api/logging"
)

const TEMPLATE_DIR_PATH = "../../templates"

// テンプレートファイルから文字列を作成します
//
// pathはこのファイルからの相対パス
func Template(path string, elements interface{}) (string, error) {
	logging.Sugar.Debugf("Use %v template.", path)

	templ, err := template.ParseFiles(fmt.Sprintf("%s/%s", TEMPLATE_DIR_PATH, path))
	if err != nil {
		return "", err
	}

	writer := new(strings.Builder)
	if err := templ.Execute(writer, elements); err != nil {
		return "", err
	}

	return writer.String(), nil
}
