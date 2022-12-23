// THIS FILE IS GENERATED. EDITING IS FUTILE.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     RawKindJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package svg

import (
	"github.com/grafana/grafana/pkg/kindsys"
)

// TODO standard generated docs
type Kind struct {
	decl kindsys.Decl[kindsys.RawProperties]
}

// type guard
var _ kindsys.Raw = &Kind{}

// TODO standard generated docs
func NewKind() (*Kind, error) {
	decl, err := kindsys.LoadCoreKind[kindsys.RawProperties]("kinds/raw/svg", nil, nil)
	if err != nil {
		return nil, err
	}

	return &Kind{
		decl: decl,
	}, nil
}

// TODO standard generated docs
func (k *Kind) Name() string {
	return "SVG"
}

// TODO standard generated docs
func (k *Kind) MachineName() string {
	return "svg"
}

// TODO standard generated docs
func (k *Kind) Maturity() kindsys.Maturity {
	return k.decl.Properties.Maturity
}

// Decl returns the [kindsys.Decl] containing both CUE and Go representations of the
// svg declaration in .cue files.
func (k *Kind) Decl() kindsys.Decl[kindsys.RawProperties] {
	return k.decl
}

// Props returns a [kindsys.SomeKindProps], with underlying type [kindsys.RawProperties],
// representing the static properties declared in the svg kind.
//
// This method is identical to calling Decl().Props. It is provided to satisfy [kindsys.Interface].
func (k *Kind) Props() kindsys.SomeKindProperties {
	return k.decl.Properties
}
