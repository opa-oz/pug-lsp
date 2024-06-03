package lsp

import (
	"github.com/opa-oz/pug-lsp/pkg/handlers"
	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/simple"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var (
	lsName  string
	version string
	handler protocol.Handler
)

func SpinUp(name, ver string, verbosity int, debug bool) {
	lsName = name
	ver = version
	// This increases logging verbosity (optional)
	commonlog.Configure(verbosity, nil)
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
		TextDocumentDidSave:    handlers.TextDocumentDidSave,
	}

	server := server.NewServer(&handler, lsName, debug)

	server.RunStdio()
}
