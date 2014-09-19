package entities

const (
	// ActionMove is the action of moving.
	ActionMove = iota + 1

	// ActionAttack is the action of attacking.
	ActionAttack

	// ActionCastSpell is the action of casting a spell.
	ActionCastSpell

	// ActionGather is the action of gathering a resource or chest.
	ActionGather

	// ActionLoot is the action of looting a corpse.
	ActionLoot

	// ActionConsume is the action of consuming something (food, drink, potion...)
	ActionConsume
)

// Action represents an action that an entity can do.
// List of actions:
//
//		move
//		attack
//		cast a spell
//		gather
//		consume
//		loot
//
type Action interface {
	SetTarget(Entity)
	SetSelf(Entity)

	GetTypeAction() uint8
	GetSelf() Entity
	GetTarget() Entity

	Play() error
}

// SimpleAction is the simplest action possible: it does nothing.
type SimpleAction struct {
	self       Entity
	target     Entity
	typeAction uint8
}

// GetTypeAction returns the type of action.
func (action *SimpleAction) GetTypeAction() uint8 {
	return action.GetTypeAction()
}

// SetTarget sets the target of the action.
func (action *SimpleAction) SetTarget(target Entity) {
	action.target = target
}

// GetTarget returns the target of the action.
func (action *SimpleAction) GetTarget() Entity {
	return action.target
}

// SetSelf set the actor of the action.
func (action *SimpleAction) SetSelf(self Entity) {
	action.self = self
}

// GetSelf return the actor of the action.
func (action *SimpleAction) GetSelf() Entity {
	return action.self
}

// Play the action.
func (action *SimpleAction) Play() error {
	return nil
}
