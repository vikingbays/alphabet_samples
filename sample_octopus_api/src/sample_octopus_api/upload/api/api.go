package api

type ParamUpload_Req struct {
	Name    string
	Summary string
	Type    int
	Alias   []string
	Author  []string
	Path    []string
}
