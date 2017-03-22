package mssql

import (
	//"database/sql"
	"database/sql/driver"
	//"encoding/binary"
	//"errors"
	//"fmt"
	//"io"
	//"math"
	//"net"
	//"reflect"
	//"strings"
	//"time"
	"golang.org/x/net/context" // use the "x/net/context" for backwards compatibility.
)
func Init() *MssqlDriver {
	return &MssqlDriver{processQueryText: true}
	//sql.Register("mssql", driverInstance)
	//sql.Register("sqlserver", driverInstanceNoProcess)
}
func (d *MssqlDriver) OpenMSConn(dsn string) (*MssqlConn, error) {
	return d.open(dsn)
}
func (c *MssqlConn) PrepareQuery(query string) (*MssqlStmt, error) {
	return c.prepareContext(context.Background(), query)
}
func (s *MssqlStmt) SendQuery(args []namedValue) (err error) {
	return s.sendQuery(args)
}

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