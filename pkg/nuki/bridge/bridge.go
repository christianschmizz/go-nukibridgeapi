package bridge

// BridgeType describes the type of a bridge
type BridgeType int

const (
	// TypeHardware represents a hardware based bridging device
	TypeHardware BridgeType = 1

	// TypeSoftware represents a software based bridge on Android
	TypeSoftware BridgeType = 2
)
