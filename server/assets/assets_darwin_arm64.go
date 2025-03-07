//go:build server

package assets

import "embed"

var (
	//go:embed fs/CIPHERC2.asc fs/*.txt fs/*.zip fs/darwin/arm64/*
	assetsFs embed.FS
)
