// Copyright 2023 The goldmark-figure authors
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package figure

import (
	"github.com/yuin/goldmark"
	aparser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/mangoumbrella/goldmark-figure/ast"
	fparser "github.com/mangoumbrella/goldmark-figure/parser"
)

type extension struct {
	renderImageLink        bool
	skipNoCaption          bool
	altTextFallbackCaption bool
}

// Figure is an extension to render <figure> elements.
var Figure = &extension{}

func (f *extension) WithImageLink() *extension {
	opts := *f
	opts.renderImageLink = true
	return &opts
}

func (f *extension) WithSkipNoCaption() *extension {
	opts := *f
	opts.skipNoCaption = true
	return &opts
}

func (f *extension) WithAltTextFallbackCaption() *extension {
	opts := *f
	opts.altTextFallbackCaption = true
	return &opts
}

func (f *extension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		aparser.WithParagraphTransformers(
			util.Prioritized(
				fparser.NewFigureParagraphTransformer(f.skipNoCaption, f.altTextFallbackCaption),
				120,
			),
		),
		aparser.WithASTTransformers(
			util.Prioritized(fparser.NewFigureASTTransformer(), 0),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(ast.NewFigureHTMLRenderer(f.renderImageLink), 0),
	))
}
