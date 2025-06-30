package capsule

import "time"

type Creator struct {
	repository Repository
}

func NewCreator(repository Repository) *Creator {
	return &Creator{
		repository: repository,
	}
}

func (c *Creator) Create(name, description string, dateToOpen time.Time) (string, error) {
	capsule := NewEntity(name, description, dateToOpen)

	if err := c.repository.Save(capsule); err != nil {
		return "", err
	}

	return capsule.Id, nil
}
