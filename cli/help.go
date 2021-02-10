package cli

import (
	"fmt"
	"strings"
)

type Help struct {
	name       string
	Desc       string
	Parameters []Parameter
}

func (h Help) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%5s%-20s%s", "", h.name, h.Desc))
	for _, param := range h.Parameters {
		sb.WriteString(fmt.Sprintf("\n%26s-%-12s%s", "", param.Name, param.Desc))
	}
	return sb.String()
}

type Parameter struct {
	Name string
	Desc string
}
