package helper

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func NewTx(tx *sql.Tx, err error) error {
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			NewLoggerConfigure("tx.log", logrus.WarnLevel, errRollback.Error(), logrus.WarnLevel)
			return errRollback
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				NewLoggerConfigure("tx.log", logrus.WarnLevel, errCommit.Error(), logrus.WarnLevel)
				return errCommit
			}
		}
	}
	return nil
}
