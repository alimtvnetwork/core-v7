package pathinternal

var (
	replacers = []fromToReplacer{
		{
			from: "////",
			to:   "/",
		},
		{
			from: "///",
			to:   "/",
		},
		{
			from: "//",
			to:   "/",
		},
		{
			from: "\\\\",
			to:   "/",
		},
		{
			from: "\\",
			to:   "/",
		},
	}
)
