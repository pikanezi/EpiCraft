package entities


// MoveAction is an action
type MoveAction struct {
	SimpleAction
}

// Play moves the actor to the target, if possible.
func (action *MoveAction) Play() error {
	switch action.target.GetType() {
	case EntityStoneWall:
		return ErrCannotMove
	}
	if action.GetSelf().GetRange(action.GetTarget()) <= 1.0 {
		return ErrTooFarAway
	}
	action.GetSelf().SetPosX(action.GetTarget().GetPosX())
	action.GetSelf().SetPosY(action.GetTarget().GetPosY())
	return nil
}
