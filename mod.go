package mod1

import "github.com/martende/mod3"
//import "github.com/golang/protobuf/proto"

//var _ = proto.Marshal

func HelloWorld() int {
	return 30 + mod3.HelloWorld()
}
