package object_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aden-q/monkey/internal/object"
)

var _ = Describe("Object", func() {
	Describe("Integer", func() {
		Context("integer object", func() {
			It("truthy integer object", func() {
				var val int64 = 5

				expectedIntegerObj := &object.Integer{
					Value: val,
				}
				obj := object.NewInteger(val)

				Expect(obj).To(Equal(expectedIntegerObj))
				Expect(obj.Inspect()).To(Equal("5"))
				Expect(obj.Type()).To(Equal(object.INTEGER_OBJ))
				Expect(obj.IsTruthy()).To(Equal(true))
			})

			It("false integer object", func() {
				var val int64 = 0

				expectedIntegerObj := &object.Integer{
					Value: val,
				}
				obj := object.NewInteger(val)

				Expect(obj).To(Equal(expectedIntegerObj))
				Expect(obj.Inspect()).To(Equal("0"))
				Expect(obj.Type()).To(Equal(object.INTEGER_OBJ))
				Expect(obj.IsTruthy()).To(Equal(false))
			})
		})
	})

	Describe("Boolean", func() {
		It("truthy boolean object", func() {
			var val bool = true

			expectedBooleanObj := object.TRUE
			obj := object.NewBoolean(val)

			Expect(obj).To(Equal(expectedBooleanObj))
			Expect(obj.Inspect()).To(Equal("true"))
			Expect(obj.Type()).To(Equal(object.BOOLEAN_OBJ))
			Expect(obj.IsTruthy()).To(Equal(true))
		})

		It("false boolean object", func() {
			var val bool = false

			expectedBooleanObj := object.FALSE
			obj := object.NewBoolean(val)

			Expect(obj).To(Equal(expectedBooleanObj))
			Expect(obj.Inspect()).To(Equal("false"))
			Expect(obj.Type()).To(Equal(object.BOOLEAN_OBJ))
			Expect(obj.IsTruthy()).To(Equal(false))
		})
	})

	Describe("Nil", func() {
		It("false nil object", func() {
			expectedNilObj := object.NIL
			obj := object.NewNil()

			Expect(obj).To(Equal(expectedNilObj))
			Expect(obj.Inspect()).To(Equal("nil"))
			Expect(obj.Type()).To(Equal(object.NIL_OBJ))
			Expect(obj.IsTruthy()).To(Equal(false))
		})
	})

	Describe("String", func() {
		It("false string object", func() {
			var val string = ""

			expectedStringObj := &object.String{
				Value: val,
			}
			obj := object.NewString(val)

			Expect(obj).To(Equal(expectedStringObj))
			Expect(obj.Inspect()).To(Equal(""))
			Expect(obj.Type()).To(Equal(object.STRING_OBJ))
			Expect(obj.IsTruthy()).To(Equal(false))
		})

		It("truthy string object", func() {
			var val string = "hello"

			expectedStringObj := &object.String{
				Value: val,
			}
			obj := object.NewString(val)

			Expect(obj).To(Equal(expectedStringObj))
			Expect(obj.Inspect()).To(Equal("hello"))
			Expect(obj.Type()).To(Equal(object.STRING_OBJ))
			Expect(obj.IsTruthy()).To(Equal(true))
		})
	})

	Describe("Hash", func() {
		It("false hash object", func() {
			items := map[object.HashKey]object.Object{}

			expectedHashObj := &object.Hash{
				Items: items,
			}
			obj := object.NewHash(items)

			Expect(obj).To(Equal(expectedHashObj))
			Expect(obj.Inspect()).To(Equal("{}"))
			Expect(obj.Type()).To(Equal(object.HASH_OBJ))
			Expect(obj.IsTruthy()).To(Equal(false))
		})

		It("truthy hash object", func() {
			items := map[object.HashKey]object.Object{
				object.NewString("key").HashKey(): object.TRUE,
			}

			expectedHashObj := &object.Hash{
				Items: items,
			}
			obj := object.NewHash(items)

			Expect(obj).To(Equal(expectedHashObj))
			Expect(obj.Inspect()).To(Equal("{key: true}"))
			Expect(obj.Type()).To(Equal(object.HASH_OBJ))
			Expect(obj.IsTruthy()).To(Equal(true))
		})
	})
})
