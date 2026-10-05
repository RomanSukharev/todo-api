// точка входа: собрать всё и запустить
// Отвечает за: поднять конфиг, БД, слои, HTTP-сервер.
// Не отвечает за: SQL, валидацию title, разбор JSON.

// config → db → repository → service → handler → ListenAndServe
package todoapi

import (
	"fmt"
)

func main() {
	fmt.Println("main.go")
}