package columnArgs

type doubleArg struct {
	ArgType   string
	Parameter string
}

//Types of double argumnets
const (
	dEFAULT = "DEFAULT"
	uNIQUE  = "UNIQUE"
	pRIMARY = "PRIMARY"
	cOMMENT = "COMMENT"
)

func createDoubleArg(argType string, parameter string) doubleArg {
	return doubleArg{ArgType: argType, Parameter: parameter}
}

func (da *doubleArg) ColumnArg() {}
