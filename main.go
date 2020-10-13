package main

import (
	"github.com/jpeach/go-message-fmt/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.MessageFormat)
}
