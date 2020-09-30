package handleError

import ( "database/sql"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)
func foo() error { return sql.ErrNoRows }
func bar() error { return foo() }

func TestError(t *testing.T) {
	err := bar()
	if err != nil {
		fmt.Printf("got err, %+v\n", err) }
}

//Outputs:
// got err, sql: no rows in result set
func TestError2(t *testing.T) {
	err := bar()
	if err == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		// Unknown error }
	}
	//Outputs:
	// data not found, sql: no rows in result set
}

// 有时候，我们需要添加更多信息时
// 会利用 fmt.Errorf()
// 但是这会导致两个不好的效果
// 1. lost call stack
// 2. lost type check , err == sql.ErrNoRows.
func foo2() error {
	return fmt.Errorf("foo err, %v", sql.ErrNoRows)
}
func bar2() error {
	return foo2()
}

func TestError3(t *testing.T) {
	err := bar2()
	if err == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		// Unknown error }
	}
	//Outputs:
	// data not found, sql: no rows in result set
}



func foo3() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed") }
func bar3() error {
	return errors.WithMessage(foo(), "bar failed") }

// 观点： If you are using github.com/pkg/errors , keep it. There is no solution better than it for now
//
func TestError3V(t *testing.T) {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error }
	}
}