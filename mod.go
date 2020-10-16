package mod1

import "github.com/martende/mod3"

func HelloWorld() int {
	return 10 + mod3.HelloWorld()
}
