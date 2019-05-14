// Copyright 2013, Chandra Sekar S.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the README.md file.

package main

import "html/template"

var (
	indexTmpl = template.Must(template.New("index").Parse(indexView))
	pkgTmpl   = template.Must(template.New("pkg").Parse(pkgView))
)

const indexView = `
<html>
	<head>
		<title>gorepos - Packages</title>
	</head>
	<body>
		<h1>Available Packages</h1>
		<ul>
			{{$host := .host}}
			{{range .pkgs}}
				<li><a href="{{.Path}}">{{$host}}{{.Path}}</a></li>
			{{end}}
		</ul>
	</body>
</html>
`

const pkgView = `
<html>
	<head>
		<meta name="go-import" content="{{.host}}{{.pkg.Path}} {{.pkg.Vcs}} {{.pkg.Repo}}">
		<title>gorepos - {{.host}}{{.pkg.Path}}</title>
	</head>
	<body>
		<h1>{{.host}}{{.pkg.Path}}</h1>
		<span style="font-weight: bold">VCS:</span> {{.pkg.Vcs}}<br>
		<span style="font-weight: bold">Repo-Root:</span> {{.pkg.Repo}}
	</body>
</html>
`
