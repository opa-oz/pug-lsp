package server

import (
	"fmt"
	"io"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/pug-lsp/pkg/documents"
	"github.com/opa-oz/pug-lsp/pkg/utils"
	"github.com/pkg/errors"
	protocol "github.com/tliron/glsp/protocol_3_16"
	glsps "github.com/tliron/glsp/server"
)

type Server struct {
	documentStore      *documents.DocumentStore
	logger             utils.Logger
	server             *glsps.Server
	clientCapabilities protocol.ClientCapabilities
	handler            *protocol.Handler
	opts               *ServerOpts
	name               string
}

type ServerOpts struct {
	Name    string
	Version string
	LogFile io.Writer
	Debug   bool
}

func NewServer(opts ServerOpts) *Server {
	if opts.LogFile == nil {
		opts.LogFile = io.Discard
	}
	logger := utils.NewFileLogger(opts.LogFile, fmt.Sprintf("[%s]\t", opts.Name), 10)
	fs, _ := utils.NewFileStore("", logger)

	handler := protocol.Handler{}
	glspServer := glsps.NewServer(&handler, opts.Name, opts.Debug)

	todo.Opts.AttachLogger(logger.Logger)

	server := Server{
		documentStore:      documents.NewDocumentStore(logger, *fs),
		logger:             logger,
		server:             glspServer,
		clientCapabilities: protocol.ClientCapabilities{},
		handler:            &handler,
		opts:               &opts,
		name:               opts.Name,
	}

	handler.Initialize = server.initialize
	handler.Initialized = server.initialized
	handler.SetTrace = server.setTrace
	handler.Shutdown = server.shutdown
	handler.TextDocumentDidOpen = server.TextDocumentDidOpen
	handler.TextDocumentDidChange = server.TextDocumentDidChange
	handler.TextDocumentCompletion = server.TextDocumentCompletion
	handler.TextDocumentHover = server.TextDocumentHover
	handler.CompletionItemResolve = server.CompletionItemResolve
	handler.TextDocumentDefinition = server.TextDocumentDefinition

	return &server
}

func (s *Server) Run() error {
	return errors.Wrap(s.server.RunStdio(), s.name)
}
