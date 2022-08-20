package bridgeapi

import "github.com/hashicorp/go-version"

// BridgeType describes the type of a bridge
type BridgeType int

const (
	// TypeHardware represents a hardware based bridging device
	TypeHardware BridgeType = 1

	// TypeSoftware represents a software based bridge on Android
	TypeSoftware BridgeType = 2
)

var (
	version1BridgesWhichSupportEncryptedTokens, version2BridgesWhichSupportEncryptedTokens version.Constraints
)

func init() {
	version1BridgesWhichSupportEncryptedTokens, _ = version.NewConstraint(">= 1.22.1, < 2")
	version2BridgesWhichSupportEncryptedTokens, _ = version.NewConstraint(">= 2.14.0, < 3")
}

// HasSupportForEncryptedTokens reports whether the bridge supports encrypted tokens
func HasSupportForEncryptedTokens(t BridgeType, firmwareVersion string) bool {
	if t != TypeHardware {
		return false
	}
	v, err := version.NewVersion(firmwareVersion)
	if err != nil {
		panic(err)
	}
	if version1BridgesWhichSupportEncryptedTokens.Check(v) || version2BridgesWhichSupportEncryptedTokens.Check(v) {
		return true
	}
	return false
}
