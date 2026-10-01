package builder

import "github.com/evanw/esbuild/pkg/api"

// MapFormat maps the public output format name to esbuild's format.
func MapFormat(format string) api.Format {
	switch format {
	case "esm":
		return api.FormatESModule
	case "cjs":
		return api.FormatCommonJS
	default:
		return api.FormatIIFE
	}
}

// MapSourceMap maps string to api.SourceMap
func MapSourceMap(sm string) api.SourceMap {
	switch sm {
	case "l":
		return api.SourceMapLinked
	case "in":
		return api.SourceMapInline
	default:
		return api.SourceMapNone
	}
}
