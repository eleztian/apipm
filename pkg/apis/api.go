// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	"github.com/eleztian/apipm/pkg/apis/middlewares"
	ui "github.com/eleztian/apipm/pkg/apis/ui/descriptors"
	v1 "github.com/eleztian/apipm/pkg/apis/v1/descriptors"

	def "github.com/caicloud/nirvana/definition"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func DescriptorAPIs() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/apis",
		Middlewares: middlewares.Middlewares(),
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			v1.Descriptor(),
		},
	}
}

func DescriptorUIs() def.Descriptor {
	return def.Descriptor{
		Description: "UIs",
		Path:        "/uis",
		Middlewares: middlewares.Middlewares(),
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			ui.Descriptor(),
		},
	}
}
