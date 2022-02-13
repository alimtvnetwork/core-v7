package internalpathextender

type IdGetter interface {
	Id() string
}

type IdIntGetter interface {
	IdInt() int
}

// NameGetter
//
//  Meaningful name to the path or the meaningful identifier
type NameGetter interface {
	// Name
	//
	//  Meaningful name to the path or the meaningful identifier
	Name() string
}

type Identifier interface {
	IdGetter
	IdIntGetter
	NameGetter
}
