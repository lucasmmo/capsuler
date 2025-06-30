package capsule

type MessageAdder struct {
	repository Repository
}

func NewMessageAdder(repository Repository) *MessageAdder {
	return &MessageAdder{
		repository: repository,
	}
}

func (m *MessageAdder) Add(id, message string) ([]string, error) {

	capsule, err := m.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := capsule.AddMessage(message); err != nil {
		return nil, err
	}

	return capsule.Messages, nil
}
