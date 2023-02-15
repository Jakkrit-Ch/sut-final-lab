package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeValidateID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check employee_id is invalid", func(t *testing.T) {
		e := Employee{
			Name:       "Jakkrit Chaiwan",
			Email:      "employee@gmail.com",
			EmployeeID: "M12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeID is invalid"))
	})
}
