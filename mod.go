package mod1

import "github.com/martende/mod3"
import "github.com/golang/protobuf/proto"

var _ = proto.Marshal

func HelloWorld() int {
	return 10 + mod3.HelloWorld()
}
