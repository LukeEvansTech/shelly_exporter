package rpc

type DeviceConfig struct {
	Host      string
	Username  string
	Password  string
	Type      string
	Mac       string
	Profile   string
	SwitchIDs []int
	CoverIDs  []int
}
