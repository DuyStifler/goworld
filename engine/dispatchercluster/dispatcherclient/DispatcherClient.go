package dispatcherclient

import (
	"net"

	"../../../../goworld/engine/common"
	"../../../../goworld/engine/consts"
	"../../../../goworld/engine/gwlog"
	"../../../../goworld/engine/netutil"
	"../../../../goworld/engine/proto"
)

type DispatcherClientType int

const (
	GameDispatcherClientType DispatcherClientType = 1 + iota
	GateDispatcherClientType
)

// DispatcherClient is a client connection to the dispatcher
type DispatcherClient struct {
	*proto.GoWorldConnection
	dctype        DispatcherClientType
	isReconnect   bool
	isRestoreGame bool
}

func newDispatcherClient(dctype DispatcherClientType, conn net.Conn, isReconnect bool, isRestoreGame bool) *DispatcherClient {
	gwc := proto.NewGoWorldConnection(netutil.NewBufferedConnection(netutil.NetConnection{conn}), false, "")
	if dctype != GameDispatcherClientType && dctype != GateDispatcherClientType {
		gwlog.Fatalf("invalid dispatcher client type: %v", dctype)
	}

	dc := &DispatcherClient{
		GoWorldConnection: gwc,
		dctype:            dctype,
		isReconnect:       isReconnect,
		isRestoreGame:     isRestoreGame,
	}
	dc.SetAutoFlush(consts.DISPATCHER_CLIENT_FLUSH_INTERVAL)
	return dc
}

// Close the dispatcher client
func (dc *DispatcherClient) Close() error {
	return dc.GoWorldConnection.Close()
}

func (dc *DispatcherClient) SendNotifyListAttrPopOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}) {
	_ = dc.GoWorldConnection.SendNotifyListAttrPopOnClient(u, id, id2, path)
}

func (dc *DispatcherClient) SendLoadEntitySomewhere(s string, id common.EntityID, u uint16) error {
	return dc.GoWorldConnection.SendLoadEntitySomewhere(s, id, u)
}

func (dc *DispatcherClient) SendCreateEntitySomewhere(u uint16, id common.EntityID, s string, data map[string]interface{}) error {
	return dc.SendCreateEntitySomewhere(u, id, s, data)
}

func (dc *DispatcherClient) SendSrvdisRegister(s string, s2 string, b bool) {
	_ = dc.GoWorldConnection.SendSrvdisRegister(s, s2, b)
}

func (dc *DispatcherClient) SendNotifyCreateEntity(id common.EntityID) error {
	return dc.GoWorldConnection.SendNotifyCreateEntity(id)
}

func (dc *DispatcherClient) SendPacket(packet *netutil.Packet) {
	_ = dc.GoWorldConnection.SendPacket(packet)
}

func (dc *DispatcherClient) SendRealMigrate(id common.EntityID, u uint16, bytes []byte) error {
	return dc.GoWorldConnection.SendRealMigrate(id, u, bytes)
}

func (dc *DispatcherClient) SendMigrateRequest(id common.EntityID, id2 common.EntityID, u uint16) error {
	return dc.GoWorldConnection.SendMigrateRequest(id, id2, u)
}

func (dc *DispatcherClient) SendNotifyDestroyEntity(id common.EntityID) error {
	return dc.GoWorldConnection.SendNotifyDestroyEntity(id)
}

func (dc *DispatcherClient) SendNotifyListAttrChangeOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}, u2 uint32, val interface{}) {
	_ = dc.GoWorldConnection.SendNotifyListAttrChangeOnClient(u, id, id2, path, u2, val)
}

func (dc *DispatcherClient) SendSetClientFilterProp(u uint16, id common.ClientID, s string, s2 string) {
	_ = dc.GoWorldConnection.SendSetClientFilterProp(u, id, s, s2)
}

func (dc *DispatcherClient) SendNotifyListAttrAppendOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}, val interface{}) {
	_ = dc.GoWorldConnection.SendNotifyListAttrAppendOnClient(u, id, id2, path, val)
}

func (dc *DispatcherClient) SendNotifyMapAttrClearOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}) {
	_ = dc.GoWorldConnection.SendNotifyMapAttrClearOnClient(u, id, id2, path)
}

func (dc *DispatcherClient) SendNotifyMapAttrDelOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}, s string) {
	_ = dc.GoWorldConnection.SendNotifyMapAttrDelOnClient(u, id, id2, path, s)
}

func (dc *DispatcherClient) SendNotifyMapAttrChangeOnClient(u uint16, id common.ClientID, id2 common.EntityID, path []interface{}, s string, val interface{}) {
	_ = dc.GoWorldConnection.SendNotifyMapAttrChangeOnClient(u, id, id2, path, s, val)
}

func (dc *DispatcherClient) SendCallEntityMethodOnClient(u uint16, id common.ClientID, id2 common.EntityID, s string, args []interface{}) {
	_ = dc.GoWorldConnection.SendCallEntityMethodOnClient(u, id, id2, s, args)
}

func (dc *DispatcherClient) SendDestroyEntityOnClient(u uint16, id common.ClientID, s string, id2 common.EntityID) {
	_ = dc.GoWorldConnection.SendDestroyEntityOnClient(u, id, s, id2)
}

func (dc *DispatcherClient) SendCreateEntityOnClient(u uint16, id common.ClientID, s string, id2 common.EntityID, b bool, clientData map[string]interface{}, f float32, f2 float32, f3 float32, f4 float32) {
	_ = dc.GoWorldConnection.SendCreateEntityOnClient(u, id, s, id2, b, clientData, f, f2, f3, f4)
}
