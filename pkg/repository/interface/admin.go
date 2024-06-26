package interfaces

import (
	"admin/service/pkg/domain"
	md "admin/service/pkg/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails md.AdminSignUp) (md.AdminDetailsResponse, error)
	FindAdminByEmail(admin md.AdminLogin) (md.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
}
