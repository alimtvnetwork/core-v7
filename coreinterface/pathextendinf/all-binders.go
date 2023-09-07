package pathextendinf

type PathExtenderBinder interface {
	PathExtender
	AsPathExtenderBinder() PathExtenderBinder
}
