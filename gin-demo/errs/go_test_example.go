package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// 第一种处理error的方式

type Panda struct{}
type Tiger struct{}

type ZooTour1 interface {
	Enter() error
	VisitPanda(panda *Panda) error
	VisitTiger(panda *Tiger) error
	Leave() error
}

// Tour1 在每个方法返回err
func Tour1(t ZooTour1, panda *Panda, tiger *Tiger) error {
	if err := t.Enter(); err != nil {
		return errors.WithMessage(err, "Enter failed")
	}
	if err := t.VisitPanda(panda); err != nil {
		return errors.WithMessagef(err, "Visited failed,panda is %v", panda)
	}
	if err := t.VisitTiger(tiger); err != nil {
		return errors.WithMessagef(err, "Visited failed,tiger is %v", tiger)
	}

	if err := t.Leave(); err != nil {
		return errors.WithMessage(err, "Leave failed")
	}
	return nil
}

// ZooTour2 第二种error处理方式
type ZooTour2 interface {
	Enter()
	VisitPanda(panda *Panda)
	VisitTiger(panda *Tiger)
	Leave()
	Err() error
}

// Tour2 屏蔽过程中的error处理
func Tour2(t ZooTour2, panda *Panda, tiger *Tiger) error {
	t.Enter()
	t.VisitPanda(panda)
	t.VisitTiger(tiger)
	t.Leave()
	// 集中编写业务处理代码，最后统一处理
	if err := t.Err(); err != nil {
		return errors.WithMessage(err, "ZooTour failed")
	}
	return nil
}

// 第三种方式，也是bufio.Scanner就是用这个方式实现的

type MyZooTour struct {
	err error
}

// 将error保存到对象内部，处理逻辑交给每个方法，本质上仍是顺序执行

func (t *MyZooTour) VisitPanda(panda *Panda) {
	if t.err != nil {
		return
	}
	// ...
}

// SafeFunc 错误保护
func SafeFunc(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("err:%v", err)
		}
	}()
	return fn()
}

//  常用的错误处理思想  An error should be handled only once
// controller
func controller() {
	_, err := ListTigersModel("tiger")
	if err != nil {
		// 追加信息
		// err = errors.WithMessage(err, "Controller: req %+v")
		log.Printf("Controller err: %+v", err)
		return
	}
}

func ListTigersModel(name string) (string, error) {
	_, err := ListTigersDao(name)
	if err != nil {
		return "", errors.Wrapf(err, "Model: ListTiger paramemters: %s", name)
	}
	return "model success", nil
}

// ListTigersDao dao层处理错误
func ListTigersDao(name string) (string, error) {
	err := errors.New("调用ListTiger报错")
	if err != nil {
		return "", errors.Wrapf(err, "Dao: ListTigers name %s", name)
	}
	return "success", nil
}

func main() {
	controller()
	fmt.Println("----------------------------")
	controller2()
}

// 返回错误码的形式

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return e.Message
}

func NewError(msg string, code int) error {
	return &MyError{
		Message: msg,
		Code:    code,
	}
}

// rpc调用
func controller2() {
	_, err := ListTigersModel2("tiger")
	if err != nil {
		mErr, ok := err.(*MyError)
		if ok {
			// 错误码 mErr.Code
			log.Printf("Controller err: %v. 错误码：%d", mErr.Message, mErr.Code)
		}
		return
	}
}
func ListTigersModel2(name string) (string, error) {
	return "", NewError("Dao 调用错误", 10001)
}
