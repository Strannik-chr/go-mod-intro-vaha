package main

import (
  "fmt"
  c "github.com/Strannik-chr/go-packages-demo/config"
  "github.com/Strannik-chr/go-packages-demo/store"
  "github.com/Strannik-chr/go-packages-demo/utils/text" /* vahaselmurzaev@Mac modul 1-14 octember  % go run .
main.go:7:3: no required module provides package github.com/Strannik-chr/go-packages-demo/utils/txet; to add it:
        go get github.com/Strannik-chr/go-packages-demo/utils/txet - говорит что не найден файл txet*/

)




func main() {
		Name := "Shadowed"
	fmt.Println("Модуль работает")
	fmt.Println(text.Hello)
	fmt.Println(text.Greeting())
	fmt.Println(Name)
	fmt.Println(c.AppName,c.Version)
	fmt.Println(store.Name)
	// Локальная переменная Name в функции main()
	// не имеет отношения к переменной store.Name из пакета store.
	// Это разные имена в разных пространствах видимости.

}

