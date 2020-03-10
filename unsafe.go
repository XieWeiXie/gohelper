package gohelper

type UnsafeAction struct {
	model interface{}
}

func NewUnsafeAction(model interface{}) *UnsafeAction {
	return &UnsafeAction{
		model: model,
	}
}
