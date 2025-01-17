package utils

import (
	"github.com/wrapped-owls/goremy-di/remy/internal/types"
	aTypes "github.com/wrapped-owls/goremy-di/remy/test/fixtures/a/testtypes"
	bTypes "github.com/wrapped-owls/goremy-di/remy/test/fixtures/b/testtypes"
	"testing"
)

func TestGetKey__Generify(t *testing.T) {
	type (
		super interface {
			a() bool
			b() string
			c(int) float32
			d(string) struct{ name string }
		}
		sub interface {
			super
		}
	)

	options := types.ReflectionOptions{}
	if GetKey[super](options) == GetKey[sub](options) {
		t.Error("type names was the same when should not generify")
	}

	options = types.ReflectionOptions{GenerifyInterface: true}
	if GetKey[super](options) != GetKey[sub](options) {
		t.Error("generified type name should be the same")
	}
}

func TestGetKey__SameStructWithDifferentPackage(t *testing.T) {
	options := types.ReflectionOptions{UseReflectionType: true}
	if GetKey[aTypes.Syringe](options) == GetKey[bTypes.Syringe](options) {
		t.Error("type names was the same, when it should be different, because of different packages")
	}

	options = types.ReflectionOptions{UseReflectionType: true}
	if GetElemKey(t, options) != GetKey[*testing.T](options) {
		t.Error("element type should be the same from type and object")
	}
}
