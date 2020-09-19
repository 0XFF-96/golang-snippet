package error_better_example2

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/apex/log"
	"github.com/pkg/errors"
)

func TestServiceMethod(t *testing.T) {
	err := ServiceMethod()
	if err != nil {
		// t.Log(err)
		log.WithError(err).Errorf("is this a err message")
	}

	if err == sql.ErrNoRows {
		t.Log("sql Err No rows")
	}

	if errors.Is(err, sql.ErrNoRows) {
		t.Log("sql err")
	}

	// Internal Error -> sql: no rows in result set !
	// why
}

func ServiceMethod() error {

	err := MiddleMethod1()
	if err != nil {
		// duplicate !
		// log.WithError(err).Errorf("MiddleMethod1")
		return err
	}

	err = MiddleMethod2()
	if err != nil {
		// log.WithError(err).Errorf("MiddleMethod2")
		return err
	}

	err = MiddleMethod3()
	if err != nil {
		// log.WithError(err).Errorf("MiddleMethod3")
		return err
	}

	return nil
}



func MiddleMethod1() error {
	param := rand.Intn(100)

	// NOTE:
	err := basicMethod()
	if err != nil {
		return errors.WithMessagef(err, "param:%d", param)
	}
	//
	//err = basicMethod2()
	//if err != nil {
	//	return err
	//}
	return nil
}

func MiddleMethod2()( err error ){
	param2 := rand.Intn(100)

	defer func()  {
		if err != nil {
			err =  errors.WithMessagef(err, "param2:%d", param2)
			// err = errors.Wrapf(err, "param2:%d", param2)
		}
	}()

	err = basicMethod()
	if err != nil {
		return err
	}

	err = basicMethod2()
	if err != nil {
		return err
	}

	err = basicMethod3()
	if err != nil {
		return err
	}
	return nil
}

func MiddleMethod3() error {
	err := basicMethod()
	if err != nil {
		return err
	}

	err = basicMethod2()
	if err != nil {
		return err
	}

	err = basicMethod3()
	if err != nil {
		return err
	}
	return nil
}

func basicMethod() error {
	return errors.WithStack(throwError())
}

func basicMethod2() error {
	return errors.WithStack(throwError())
}

func basicMethod3() error {
	return errors.WithStack(throwError())
}

// external error causation
func throwError() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i > 8 {
		return sql.ErrNoRows
	} else {
		return nil
	}
}

