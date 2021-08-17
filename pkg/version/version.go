package version

type ToolBox struct {
	Name    string
	Version string
}

func (tb ToolBox) Info() (pkgName, pkgVersion string) {
	// fmt.Println(tb.name)
	// fmt.Println(tb.version)
	return tb.Name, tb.Version
}
