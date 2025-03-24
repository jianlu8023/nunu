package tpl

import "embed"

//go:embed create/*.tpl
var CreateTemplateFS embed.FS

//go:embed mycreate/*.tpl
var MyCreateTemplateFS embed.FS
