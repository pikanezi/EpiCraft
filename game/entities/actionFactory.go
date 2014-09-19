package entities

// NewAction returns the action specified by the typeAction provided.
func NewAction(self, target Entity, typeAction uint8) Action {
	switch typeAction {

	}
	return &SimpleAction{
		self:       self,
		target:     target,
		typeAction: typeAction,
	}
}
