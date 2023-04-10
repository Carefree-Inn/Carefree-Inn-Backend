package errno

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func Throw() error {
	return ThrowNot()
}

func ThrowNot() error {
	return errors.WithStack(errors.New("eee"))
}

func TestWrapError(t *testing.T) {
	fmt.Printf("%+v", Throw())
}
