package types

import "github.com/chuckleheads/habitat/components/core/ident"

type PagedSearchResp struct {
	RangeStart int           `mapstructure:"range_start"`
	RangeEnd   int           `mapstructure:"range_end"`
	TotalCount int           `mapstructure:"total_count"`
	Data       []ident.Ident `mapstructure:"data"`
}

type RDepsResp struct {
	RDeps []string `mapstructure:"rdeps"`
}

type ChannelName struct {
	Name string `mapstructure:"name"`
}
