package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"` // ต้องไม่เป็นค่าว่าง
	Email      string
	EmployeeID string 	`valid:"matches(^[JMS]{1}\\d{8}$)~EmployeeID is invalid"`	// รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
}
