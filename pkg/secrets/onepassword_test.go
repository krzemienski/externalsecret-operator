package secrets

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOnePassword(t *testing.T) {
	secretKey := "testkey"
	secretValue := "testvalue"
	expectedValue := secretValue

	Convey("Given an initialized OnePasswordBackend", t, func() {
		backend := NewOnePasswordBackend("Personal", StubOnePasswordClient{})

		Convey("When retrieving a secret", func() {
			actualValue, err := backend.Get(secretKey)
			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
				So(actualValue, ShouldEqual, expectedValue)
			})
		})
	})
}
