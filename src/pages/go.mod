module github.com/glxxyz/gohanzi/pages

go 1.15

replace (
	github.com/glxxyz/gohanzi/containers v0.0.0 => ../containers
	github.com/glxxyz/gohanzi/repo v0.0.0 => ../repo
)

require (
	github.com/glxxyz/gohanzi/containers v0.0.0
	github.com/glxxyz/gohanzi/repo v0.0.0
)
