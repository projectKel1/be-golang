package service

import (
	"errors"
	"group-project-3/features/company"
	"group-project-3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	mockCompanyDataLayer := new(mocks.CompanyData)
	returnData := []company.Core{{
		ID:          1,
		Name:        "PT. Abadi Selamanya",
		Address:     "Surabaya",
		Description: "Perusahaan Abadi",
		Email:       "support@abadi.com",
		Type:        "Retail",
		StartedHour: "09:00",
		EndedHour:   "17:00",
		Visi:        "Menjalankan Bisnis",
		Misi:        "Menjalankan Usaha",
	}}

	t.Run("test case success read all data", func(t *testing.T) {
		pageNumber := 1
		pageSize := 0
		mockCompanyDataLayer.On("SelectAll", pageNumber, pageSize).Return(returnData, nil).Once()
		srv := New(mockCompanyDataLayer)
		result, err := srv.GetAll(1, 0)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, result[0].Name)
		mockCompanyDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed read all data", func(t *testing.T) {
		pageNumber := 0
		pageSize := 0
		mockCompanyDataLayer.On("SelectAll", pageNumber, pageSize).Return(nil, errors.New("error read data")).Once()
		srv := New(mockCompanyDataLayer)
		result, err := srv.GetAll(0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		mockCompanyDataLayer.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	id := 1
	mockCompanyDataLayer := new(mocks.CompanyData)
	t.Run("test case success delete company", func(t *testing.T) {
		mockCompanyDataLayer.On("Delete", uint(id)).Return(nil).Once()
		srv := New(mockCompanyDataLayer)
		err := srv.DeleteById(1)
		assert.Nil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed delete company", func(t *testing.T) {
		mockCompanyDataLayer.On("Delete", uint(id)).Return(errors.New("error id not valid")).Once()
		srv := New(mockCompanyDataLayer)
		err := srv.DeleteById(1)
		assert.NotNil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	mockCompanyDataLayer := new(mocks.CompanyData)
	t.Run("test case success update company data", func(t *testing.T) {
		id := 1
		inputData := company.Core{
			Name:        "PT. Aman Sejahtera",
			Address:     "Malang",
			Description: "Perusahaan di bidang keamanan data",
			Email:       "help@aman.com",
			Type:        "Tech",
			StartedHour: "10:00",
			EndedHour:   "18:00",
			Visi:        "Menjadi perusahaan keamanan nomor 1 di Indonesia",
			Misi:        "Indonesia bebas cybercrime",
		}
		mockCompanyDataLayer.On("Update", uint(id), inputData).Return(nil).Once()
		srv := New(mockCompanyDataLayer)
		err := srv.EditById(1, inputData)
		assert.Nil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed update company data", func(t *testing.T) {
		id := 10000
		inputData := company.Core{
			Name:        "PT. Aman Sejahtera",
			Address:     "Malang",
			Description: "Perusahaan di bidang keamanan data",
			Email:       "help@aman.com",
			Type:        "Tech",
			StartedHour: "10:00",
			EndedHour:   "18:00",
			Visi:        "Menjadi perusahaan keamanan nomor 1 di Indonesia",
			Misi:        "Indonesia bebas cybercrime",
		}
		mockCompanyDataLayer.On("Update", uint(id), inputData).Return(errors.New("error update data company")).Once()
		srv := New(mockCompanyDataLayer)
		err := srv.EditById(uint(id), inputData)
		assert.NotNil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})

}

func TestGetById(t *testing.T) {
	mockCompanyDataLayer := new(mocks.CompanyData)
	returnData := company.Core{
		ID:          1,
		Name:        "PT. Indonesia Pintar",
		Address:     "Jakarta",
		Description: "Perusahaan di bidang kepintaran",
		Email:       "helpdesk@ipintar.com",
		Type:        "Edutech",
		StartedHour: "07:00",
		EndedHour:   "15:00",
		Visi:        "Memajukan ilmu pengetahuan bangsa",
		Misi:        "Indonesia menjadi Negara dengan IQ tinggi",
	}
	t.Run("test case success read data", func(t *testing.T) {
		id := 1
		mockCompanyDataLayer.On("Select", uint(id)).Return(returnData, nil).Once()
		srv := New(mockCompanyDataLayer)
		result, err := srv.GetById(uint(id))
		assert.Nil(t, err)
		assert.Equal(t, returnData, result)
		mockCompanyDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed read data", func(t *testing.T) {
		id := 10000
		mockCompanyDataLayer.On("Select", uint(id)).Return(company.Core{}, errors.New("error read data")).Once()
		srv := New(mockCompanyDataLayer)
		result, err := srv.GetById(uint(id))
		assert.NotNil(t, err)
		assert.NotNil(t, result)
		mockCompanyDataLayer.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockCompanyDataLayer := new(mocks.CompanyData)
	t.Run("test case success create company", func(t *testing.T) {
		inputData := company.Core{
			Name:        "PT. Adil Makmur",
			Address:     "Solo",
			Description: "Perusahaan yang menegakkan keadilan manusia",
			Email:       "HR@adil.com",
			Type:        "Law",
			StartedHour: "09:00",
			EndedHour:   "17:00",
			Visi:        "Menegakkan keadilan tanpa pandang bulu",
			Misi:        "Meningkatkan tingkat keadilan kalangan bawah di Indonesia",
		}
		mockCompanyDataLayer.On("Insert", inputData).Return(nil).Once()
		srv := New(mockCompanyDataLayer)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed create company, email required", func(t *testing.T) {
		inputData := company.Core{
			Name:        "PT. Adil Makmur",
			Address:     "Solo",
			Description: "Perusahaan yang menegakkan keadilan manusia",
			Type:        "Law",
			StartedHour: "09:00",
			EndedHour:   "17:00",
			Visi:        "Menegakkan keadilan tanpa pandang bulu",
			Misi:        "Meningkatkan tingkat keadilan kalangan bawah di Indonesia",
		}
		srv := New(mockCompanyDataLayer)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		mockCompanyDataLayer.AssertExpectations(t)
	})

}
