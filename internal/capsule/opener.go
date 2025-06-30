package capsule

type Opener struct {
	repository Repository
}

func NewOpener(repository Repository) *Opener {

	return &Opener{
		repository: repository,
	}
}

func (o *Opener) Open(id string) error {
	capsule, err := o.repository.GetById(id)
	if err != nil {
		return err
	}

	if err := capsule.Open(); err != nil {
		return err
	}

	if err := o.repository.Save(capsule); err != nil {
		return err
	}

	return nil
}
