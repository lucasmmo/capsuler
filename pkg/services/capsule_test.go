package services_test

import (
	"capsuler/pkg/capsule"
	"capsuler/pkg/services"
	"capsuler/pkg/testify"
	"testing"
	"time"
)

type mockCapsuleRepository struct {
	DB map[string]*capsule.Capsule
}

func NewMockCapsuleRepository() *mockCapsuleRepository {
	return &mockCapsuleRepository{
		DB: make(map[string]*capsule.Capsule),
	}
}

func (m *mockCapsuleRepository) Save(capsule *capsule.Capsule) error {
	m.DB[capsule.GetId()] = capsule
	return nil
}

func (m *mockCapsuleRepository) GetById(id string) (*capsule.Capsule, error) {
	return m.DB[id], nil
}

func TestCapsuleService(t *testing.T) {
	mockCapsuleRepository := NewMockCapsuleRepository()

	t.Run("should be able to create a capsule", func(t *testing.T) {
		// Arrange
		name := "capsule_test"
		description := "capsule's test"
		dateToOpen := time.Now()

		// Act
		_, err := services.CreateCapsule(name, description, dateToOpen, mockCapsuleRepository)

		// Assert
		testify.AssertNil(t, err)
	})

	t.Run("should be able to open a capsule", func(t *testing.T) {
		// Arrange
		capsule := capsule.Builder().Build()
		mockCapsuleRepository.Save(capsule)

		// Act
		err := services.OpenCapsule(capsule.GetId(), mockCapsuleRepository)

		// Assert
		testify.AssertNil(t, err)
		testify.AssertTrue(t, capsule.WasOpened())
	})

	t.Run("should be able to add a new message a capsule", func(t *testing.T) {
		// Arrange
		capsule := capsule.Builder().Build()
		mockCapsuleRepository.Save(capsule)
		message := "message test"

		// Act
		capsule, err := services.AddMessageToCapsule(capsule.GetId(), message, mockCapsuleRepository)

		// Assert
		testify.AssertNil(t, err)
		testify.AssertNotEmptyStrSlice(t, capsule.GetMessages())
	})
}
