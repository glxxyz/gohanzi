module github.com/glxxyz/gohanzi

go 1.15

replace (
	github.com/glxxyz/gohanzi/containers v0.0.0 => ./src/containers
	github.com/glxxyz/gohanzi/pages v0.0.0 => ./src/pages
	github.com/glxxyz/gohanzi/repo v0.0.0 => ./src/repo
)

require (
	github.com/glxxyz/gohanzi/containers v0.0.0
	github.com/glxxyz/gohanzi/pages v0.0.0
	github.com/glxxyz/gohanzi/repo v0.0.0
)
