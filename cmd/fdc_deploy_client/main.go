package main

import (
	"fmt"
	proto "github.com/rirush/fdc_deploy/fdc_protocol"
)

func main() {
	pkg := proto.Package{
		Command:   proto.COMMAND_PING,
		Arguments: []string{"Welp", "It's an argument", "What else should I say?"},
	}
	data := proto.Marshal(pkg)
	unmarshalledPkg := proto.Package{}
	proto.Unmarshal(data, &unmarshalledPkg)
	fmt.Println(data)
	fmt.Println(unmarshalledPkg)
}
