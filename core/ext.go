package core

var (
	globalExtMap  map[string]*Ext     = make(map[string]*Ext)
	globalFuncMap map[string]Callable = make(map[string]Callable)
)

type Ext struct {
	Name      string
	Version   string
	Functions map[string]*ExtFunction
	Constants map[ZString]*ZVal
	Classes   []*ZClass
}

type ExtFunction struct {
	Func func(ctx Context, args []*ZVal) (*ZVal, error)
	Args []*ExtFunctionArg
}

func (e *ExtFunction) Call(ctx Context, args []*ZVal) (*ZVal, error) {
	return e.Func(ctx, args)
}

type ExtFunctionArg struct {
	ArgName  string // without the $ sign
	Optional bool   // is this argument optional?
}

func RegisterExt(e *Ext) {
	globalExtMap[e.Name] = e
}

func HasExt(name string) bool {
	_, res := globalExtMap[name]
	return res
}

func GetExt(name string) *Ext {
	v, ok := globalExtMap[name]
	if ok {
		return v
	}
	return nil
}
