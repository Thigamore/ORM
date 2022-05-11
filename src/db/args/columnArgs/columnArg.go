package columnArgs

import "errors"

type ColumnArg interface {
	ColumnArg()
}

//Creates a column argument
func CreateColumnArg(argType string, parameters ...string) (ColumnArg, error) {
	//goes through the arg types
	switch argType {
	case "NULL":
		return createSingleArg("NULL"), nil
	case "NOT NULL":
		return createSingleArg("NOT NULL"), nil
	default:
		return nil, errors.New("Could not find the parameter type for " + argType)
	}
}
