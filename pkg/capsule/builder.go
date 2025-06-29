package capsule

import "time"

type capsuleBuilder struct {
	capsule *Capsule
}

func Builder() *capsuleBuilder {
	return &capsuleBuilder{
		capsule: newCapsule(),
	}
}

func (b *capsuleBuilder) WithId(value string) *capsuleBuilder {
	b.capsule.id = value
	return b
}

func (b *capsuleBuilder) WithCreatedAt(value time.Time) *capsuleBuilder {
	b.capsule.createdAt = value
	return b
}

func (b *capsuleBuilder) WithUpdatedAt(value time.Time) *capsuleBuilder {
	b.capsule.updatedAt = value
	return b
}

func (b *capsuleBuilder) WithMessages(value []string) *capsuleBuilder {
	b.capsule.messages = value
	return b
}

func (b *capsuleBuilder) WithName(value string) *capsuleBuilder {
	b.capsule.name = value
	return b
}

func (b *capsuleBuilder) WithDescription(value string) *capsuleBuilder {
	b.capsule.description = value
	return b
}

func (b *capsuleBuilder) WithDateToOpen(value time.Time) *capsuleBuilder {
	b.capsule.dateToOpen = value
	return b
}

func (b *capsuleBuilder) WithIsOpen(value bool) *capsuleBuilder {
	b.capsule.isOpen = value
	return b
}

func (b *capsuleBuilder) Build() *Capsule {
	return b.capsule
}
