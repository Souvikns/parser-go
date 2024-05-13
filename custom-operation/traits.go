package customoperation

import (
	"fmt"

	"github.com/bserdar/gojsonpath"
)

var v2TraitPaths = []string{
	// operation
	"$.channels.*.[publish,subscribe]",
	"$.components.channels.*.[publish,subscribe]",
	// messages
	"$.channels.*.[publish,subscribe].message",
	"$.channels.*.[publish,subscribe].message.oneOf.*",
	"$.components.channels.*.[publish,subscribe].message",
	"$.components.channels.*.[publish,subscribe].message.oneOf.*",
	"$.components.messages.*",
}

var v3TraitPaths = []string{
	// operations
	"$.operations.*",
	"$.operations.*.channel.*",
	"$.operation.*.channel.messages.*",
	"$.operations.*.messages.*",
	"$.components.operations.*",
	"$.components.operations.*.channel.*",
	"$.components.operations.*.channel.messages.*",
	"$.components.operations.*.messages.*",
	// Channels
	"$.channels.*.messages.*",
	"$.components.channels.*.messages.*",
	// messages
	"$.components.messages.*",
}


func applyAllTraitsV2(asyncapi any, paths []string) {
	for _, i := range(paths) {
		path, _ := gojsonpath.Parse(i)
		fmt.Println(path)
	}
}
