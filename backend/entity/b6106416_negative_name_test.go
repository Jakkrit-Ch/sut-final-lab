package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name cannot be blank", func(t *testing.T) {
		e := Employee{
			Name:       "",
			Email:      "employee@gmail.com",
			EmployeeID: "6106416",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})
}
