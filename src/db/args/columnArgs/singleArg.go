package columnArgs

type singleArg struct {
	ArgType string
}

//General arguments that just add itself to a column
const (
	nULL           = "NULL"
	nOT_NULL       = "NOT NULL"
	vISIBLE        = "VISIBLE"
	iNVISIBLE      = "INVISIBLE"
	aUTO_INCREMENT = "AUTO_INCREMENT"
)

func createSingleArg(argType string) singleArg {
	return singleArg{ArgType: argType}
}

func (sa singleArg) ColumnArg() {}
