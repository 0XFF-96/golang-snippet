package example2

import (
	"testing"
)

type TestRepo interface {
	TestStringFunc() (string,error)
}

type baseDao struct {
	status string
}

func NewBaseDao() TestRepo {
	return &baseDao{
		status: "1",
	}
}

func (b baseDao) TestStringFunc() (string,error ){return b.status, nil}


type Repository interface {
	// 需要创建了数据表并且 make xo 之后手动在这里添加 BaseRepo，否则编译不过
	TestRepo
	RepoFunc2() error
}


type dao struct {
	TestRepo
	baseStatus string
}

func (d *dao) RepoFunc2() error {
	return nil
}

func NewDao() Repository {
	base := NewBaseDao()
	return &dao{
		baseStatus: "dao",
		TestRepo: base,
	}
}


// 不行，空指针
func TestDao(t *testing.T){
	repo := NewDao()

	t.Log(repo.RepoFunc2())
	t.Log(repo.TestStringFunc())
}


func TestDao2(t *testing.T) {

}