package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func main() {
	fmt.Println(govalidator.IsURL(`http://www.google.com`))
	fmt.Println(govalidator.IsURL(`http://user@pass:domain.com/path/page`))
}
