package ident

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Ident represents a package identifier
type Ident struct {
	Origin  string `mapstructure:"origin"`
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Release string `mapstructure:"release"`
}

// New creates a new instance of an Ident from its parts
func New(origin string, name string, version string, release string) (Ident, error) {
	newIdent := Ident{}
	if origin == "" || name == "" {
		errors.Errorf("Origin and Name are required")
	}
	newIdent.Origin = origin
	newIdent.Name = name
	if version != "" {
		newIdent.Version = version
	}
	if release != "" {
		newIdent.Release = release
	}
	return newIdent, nil
}

// FromString takes a string in the form of ORIGIN/NAME/VERSION/RELEASE
// and resturns a HabPkg struct
func FromString(rawIdent string) (Ident, error) {
	pkg := Ident{}
	return pkg.FromString(rawIdent)
}

// FromString takes a string in the form of ORIGIN/NAME/VERSION/RELEASE
// and resturns a Ident struct
func (ident Ident) FromString(rawIdent string) (Ident, error) {
	parts := strings.Split(rawIdent, "/")
	if len(parts) < 2 || len(parts) > 4 {
		return Ident{}, errors.Errorf("invalid package identifier '%s'", rawIdent)
	}

	for i, part := range parts {
		switch i {
		case 0:
			ident.Origin = part
		case 1:
			ident.Name = part
		case 2:
			ident.Version = part
		case 3:
			ident.Release = part
		}
	}

	return ident, nil
}

// String makes Ident satisfy the Stringer interface.
func (ident Ident) String() string {
	formattedIdent := filepath.Join(ident.Origin, ident.Name)
	if ident.Version != "" {
		formattedIdent = filepath.Join(formattedIdent, ident.Version)
	}
	if ident.Release != "" {
		formattedIdent = filepath.Join(formattedIdent, ident.Release)
	}
	return formattedIdent
}
