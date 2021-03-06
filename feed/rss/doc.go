// Tideland Go Library - RSS Feed
//
// Copyright (C) 2012-2015 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// The Tideland Go Library atom package implements an RSS feed client.
//
// The RSS package provides the RSS XML schema as Go types for the
// usage with the standard marshalling/unmarshalling. The supported
// format is RSS 2.0. A client allows to retrieve RSS documents.
package rss

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/golib/version"
)

//--------------------
// VERSION
//--------------------

// PackageVersion returns the version of the version package.
func PackageVersion() version.Version {
	return version.New(3, 0, 0)
}

// EOF
