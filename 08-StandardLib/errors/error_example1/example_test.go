package error_example1

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/apex/log"
)

// 调用关系图
// ServiceMethod
//	- MiddleMethod1
//			- basicMethod()
//			- basicMethod2()
//	- MiddleMethod2
//			- basicMethod()
//	- MiddleMethod3
//			- basicMethod()
//			- basicMethod2()
//			- basicMethod3()

func TestServiceMethod(t *testing.T) {
	err := ServiceMethod()
	if err != nil {
		t.Log(err)
	}

	// Internal Error -> sql: no rows in result set !
	// why
}

func ServiceMethod() error {

	err := MiddleMethod1()
	if err != nil {
		// duplicate !
		log.WithError(err).Errorf("MiddleMethod1")
		return err
	}

	err = MiddleMethod2()
	if err != nil {
		log.WithError(err).Errorf("MiddleMethod2")
		return err
	}

	err = MiddleMethod3()
	if err != nil {
		log.WithError(err).Errorf("MiddleMethod3")
		return err
	}

	return nil
}



func MiddleMethod1() error {
	err := basicMethod()
	if err != nil {
		return err
	}
	return nil
}

func MiddleMethod2() error {
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
	return throwError()
}

func basicMethod2() error {
	return throwError()
}

func basicMethod3() error {
	return throwError()
}

func throwError() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i > 8 {
		return sql.ErrNoRows
	} else {
		return nil
	}
}