package cmd

const (
	supernodeKeyPrefix         = "supernode."
	supernodeListenPortKey     = supernodeKeyPrefix + "listenPort"
	supernodeManagementPortKey = supernodeKeyPrefix + "managementPort"

	edgeKeyPrefix               = "edge."
	edgeAllowP2PKey             = edgeKeyPrefix + "allowP2P"
	edgeAllowRoutingKey         = edgeKeyPrefix + "allowRouting"
	edgeCommunityNameKey        = edgeKeyPrefix + "communityName"
	edgeDisablePMTUDiscoveryKey = edgeKeyPrefix + "disablePMTUDiscovery"
	edgeDisableMulticastKey     = edgeKeyPrefix + "disableMulticast"
	edgeDynamicIPModeKey        = edgeKeyPrefix + "dynamicIPMode"
	edgeEncryptionKeyKey        = edgeKeyPrefix + "encryptionKey"
	edgeLocalPortKey            = edgeKeyPrefix + "localPort"
	edgeManagementPortKey       = edgeKeyPrefix + "managementPort"
	edgeRegisterIntervalKey     = edgeKeyPrefix + "registerInterval"
	edgeRegisterTTLKey          = edgeKeyPrefix + "registerTTL"
	edgeSupernodeHostPortKey    = edgeKeyPrefix + "supernodeHostPort"
	edgeTypeOfServiceKey        = edgeKeyPrefix + "typeOfService"
	edgeEncryptionMethodKey     = edgeKeyPrefix + "encryptionMethod"
	edgeDeviceNameKey           = edgeKeyPrefix + "deviceName"
	edgeAddressModeKey          = edgeKeyPrefix + "addressMode"
	edgeDeviceIPKey             = edgeKeyPrefix + "deviceIP"
	edgeDeviceNetmaskKey        = edgeKeyPrefix + "deviceNetmask"
	edgeDeviceMACAddressKey     = edgeKeyPrefix + "deviceMACAddress"
	edgeMTUKey                  = edgeKeyPrefix + "MTU"
)

const (
	couldNotBindFlagsErrorMessage = "Could not bind flags"
)
