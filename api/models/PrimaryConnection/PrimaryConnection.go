package PrimaryConnection

type Model struct {
	IPv4            string
	IPv6            string
	MAC             string
	VLANID          int64
	InterfaceTypeID int64
	IsReservedIP    bool
	IsVirtualIP     bool
	IsVirtualIFace  bool
}
