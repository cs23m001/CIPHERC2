//go:build server

package assets

import "embed"

var (
	//go:embed fs/CIPHERC2.asc fs/*.txt fs/*.zip fs/linux/amd64/*
	assetsFs embed.FS
)
