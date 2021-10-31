package admins_test

import (
	"altaStore/business"
	"altaStore/business/admins"
	adminMock "altaStore/business/admins/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id            = 1
	name          = "name"
	username      = "username"
	password      = "12345"
	status        = "active"
	phone_number  = "1234567890"
	email         = "email@gmail.com"
	created_by    = "created_by"
	limit         = 100
	offset        = 1
	created_by_id = 2
)

var (
	adminRepository adminMock.Repository
	adminData       admins.Admins
	adminService    admins.Service
	listAdmins      []admins.Admins
	InsertAdminData admins.AdminSpec
	adminUpdate     admins.AdminUpdatable
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLoginAdmin(t *testing.T) {
	t.Run("Expects admins login success", func(t *testing.T) {
		adminRepository.On("LoginAdmin", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		admin, err := adminService.LoginAdmin(username, password)
		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, username, admin.Username)
		assert.Equal(t, name, admin.Name)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, password, admin.Password)
		assert.Equal(t, status, admin.Status)
		assert.Equal(t, phone_number, admin.Phone_number)
		assert.Equal(t, created_by, admin.CreatedBy)
	})

	t.Run("Expects admin failed loign", func(t *testing.T) {
		adminRepository.On("LoginAdmin", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.LoginAdmin(username, password)

		assert.Nil(t, admin)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAllAdmin(t *testing.T) {
	t.Run("Expects find all admins", func(t *testing.T) {
		adminRepository.On("GetAdmin", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&listAdmins, nil).Once()
		admin, err := adminService.FindAllAdmin(offset, limit)
		assert.Nil(t, err)
		assert.NotNil(t, admin)
	})

	t.Run("Expects failed find all admins", func(t *testing.T) {
		adminRepository.On("GetAdmin", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAllAdmin(offset, limit)
		assert.NotNil(t, err)
		assert.Nil(t, admin)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAdminByUsername(t *testing.T) {
	t.Run("Expects find admins", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		admin, err := adminService.FindAdminByUsername(username)
		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, username, admin.Username)
		assert.Equal(t, name, admin.Name)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, password, admin.Password)
		assert.Equal(t, status, admin.Status)
		assert.Equal(t, phone_number, admin.Phone_number)
		assert.Equal(t, created_by, admin.CreatedBy)
	})

	t.Run("Expects admins not found", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAdminByUsername(username)
		assert.NotNil(t, err)
		assert.Nil(t, admin)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAdminById(t *testing.T) {
	t.Run("Expects find admins", func(t *testing.T) {
		adminRepository.On("GetAdminById", mock.AnythingOfType("int")).Return(&adminData, nil).Once()
		admin, err := adminService.FindAdminById(id)
		assert.Nil(t, err)
		assert.NotNil(t, admin)

		assert.Equal(t, username, admin.Username)
		assert.Equal(t, name, admin.Name)
		assert.Equal(t, email, admin.Email)
		assert.Equal(t, password, admin.Password)
		assert.Equal(t, status, admin.Status)
		assert.Equal(t, phone_number, admin.Phone_number)
		assert.Equal(t, created_by, admin.CreatedBy)
	})

	t.Run("Expects admins not found", func(t *testing.T) {
		adminRepository.On("GetAdminById", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
		admin, err := adminService.FindAdminById(id)
		assert.NotNil(t, err)
		assert.Nil(t, admin)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertAdmin(t *testing.T) {

	t.Run("Expects insert admins failed, unauthorized admin token", func(t *testing.T) {
		adminRepository.On("GetAdminById", created_by_id).Return(nil, business.ErrNotFound).Once()
		err := adminService.InsertAdmin(InsertAdminData, created_by_id)
		assert.NotNil(t, err)
	})

	t.Run("Expects insert admins failed", func(t *testing.T) {
		adminRepository.On("GetAdminById", mock.AnythingOfType("int")).Return(&adminData, nil).Once()
		adminRepository.On("CreateAdmin", mock.AnythingOfType("*admins.Admins")).Return(business.ErrInternalServerError).Once()
		err := adminService.InsertAdmin(InsertAdminData, created_by_id)
		assert.NotNil(t, err)
	})

	t.Run("Expects insert admins success", func(t *testing.T) {
		adminRepository.On("GetAdminById", mock.AnythingOfType("int")).Return(&adminData, nil).Once()
		adminRepository.On("CreateAdmin", mock.AnythingOfType("*admins.Admins")).Return(nil).Once()
		err := adminService.InsertAdmin(InsertAdminData, created_by_id)
		assert.Nil(t, err)
	})

	t.Run("Expects insert admins failed, invalid input", func(t *testing.T) {
		adminRepository.On("GetAdminById", mock.AnythingOfType("int")).Return(&adminData, nil).Once()
		adminRepository.On("CreateAdmin", mock.AnythingOfType("*admins.Admins")).Return(nil).Once()
		err := adminService.InsertAdmin(admins.AdminSpec{
			Name:         name,
			Password:     password,
			Email:        "wrong!!!",
			Phone_number: phone_number,
			Username:     username,
		}, created_by_id)
		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInvalidSpec, err)
	})

}

func TestModifyAdmin(t *testing.T) {
	t.Run("Expects update admins success", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", username).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("*admins.Admins")).Return(nil).Once()
		err := adminService.ModifyAdmin(username, adminUpdate, 2)
		assert.Nil(t, err)
	})

	t.Run("Expects update admins failed, username not found", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", username).Return(nil, business.ErrNotFound).Once()
		err := adminService.ModifyAdmin(username, adminUpdate, 2)
		assert.NotNil(t, err)
	})

	t.Run("Expects update admins failed", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", username).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("*admins.Admins")).Return(business.ErrInternalServerError).Once()
		err := adminService.ModifyAdmin(username, adminUpdate, 2)
		assert.NotNil(t, err)
	})

	t.Run("Expects update admins failed, Invalid Input", func(t *testing.T) {
		adminRepository.On("GetAdminByUsername", username).Return(&adminData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("*admins.Admins")).Return(business.ErrInternalServerError).Once()
		err := adminService.ModifyAdmin(username, admins.AdminUpdatable{
			Name:         name,
			Phone_number: phone_number,
			Status:       status + "1234567890121212",
		}, 2)
		assert.NotNil(t, err)
	})
}

func setup() {
	InsertAdminData = admins.AdminSpec{
		Name:         name,
		Password:     password,
		Email:        email,
		Username:     username,
		Phone_number: phone_number,
	}

	adminData = admins.NewAdmin(
		name,
		username,
		email,
		password,
		phone_number,
		created_by,
	)

	adminUpdate = admins.AdminUpdatable{
		Name:         name,
		Phone_number: phone_number,
		Status:       status,
	}

	adminService = admins.InitAdminService(&adminRepository)

	listAdmins = append(listAdmins, adminData)

}
