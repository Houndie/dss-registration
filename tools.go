// +build tools

// This file is generated and managed by toolbox.  Manually edit at your own peril.
package toolbox

import (
	_ "github.com/gohugoio/hugo"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate" //{"build_flags":"-tags 'postgres'"}
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/thechriswalker/protoc-gen-twirp_js"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
)
