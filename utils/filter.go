package utils

type Filter struct {
	Limit  uint `query:"limit" header:"limit" json:"limit" xml:"limit" form:"limit"`
	Offset uint `query:"offset" header:"offset" json:"offset" xml:"offset" form:"offset"`
	// Order?
}

func (f *Filter) Validate() {
	if f.Limit != 0 {
		f.Limit = max(1, min(FilterLimitMAX, f.Limit))
	} else {
		f.Limit = FilterLimitMAX
	}
}
