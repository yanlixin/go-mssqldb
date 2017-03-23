package mssql

import (
	"database/sql/driver"
	"golang.org/x/net/context" // use the "x/net/context" for backwards compatibility.
)

func (s *MssqlStmt) ExecOutPutParams(args []driver.NamedValue) (driver.Result, error) {
	list := make([]namedValue, len(args))
	for i, nv := range args {
		list[i] = namedValue(nv)
	}
	return s.exec(context.Background(), list)
}
func (r *MssqlStmt) OutPutParams() ([]driver.NamedValue, error) {
	return r.params, nil
}