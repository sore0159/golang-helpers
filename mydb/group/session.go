package group

import (
	"mule/mydb/db"
	sq "mule/mydb/sql"
)

// Session is a convienience struct to automate the main
// usage pattern of SQL groups
type Session struct {
	Group SQLGrouper
	D     db.DBer
}

func NewSession(group SQLGrouper, d db.DBer) *Session {
	return &Session{
		Group: group,
		D:     d,
	}
}

func (s *Session) Select(conditions ...interface{}) error {
	ps, err := sq.Convert2P(conditions)
	if my, bad := Check(err, "Session select failed on argument conversion", "args", conditions); bad {
		return my
	}
	return Select(s.D, s.Group, ps...)
}
func (s *Session) SelectWhere(where sq.Condition) error {
	return SelectWhere(s.D, s.Group, where)
}

func (s *Session) Delete(where sq.Condition) error {
	return DeleteWhere(s.D, s.Group.SQLTable(), where)
}

func (s *Session) Close() (err error) {
	err = Insert(s.D, s.Group)
	if my, bad := Check(err, "session close failure on insert"); bad {
		return my
	}
	err = Update(s.D, s.Group)
	if my, bad := Check(err, "session close failure on update"); bad {
		return my
	}
	err = Delete(s.D, s.Group)
	if my, bad := Check(err, "session close failure on delete"); bad {
		return my
	}
	return nil
}
