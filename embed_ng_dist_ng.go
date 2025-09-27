package gongreqif

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongreqif/dist/ng-github.com-fullstack-lang-gongreqif
var NgDistNg embed.FS
