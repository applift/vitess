// Code generated by protoc-gen-go.
// source: topodata.proto
// DO NOT EDIT!

/*
Package topodata is a generated protocol buffer package.

It is generated from these files:
	topodata.proto

It has these top-level messages:
	KeyRange
	TabletAlias
	Tablet
	Shard
	Keyspace
	ShardReplication
	ShardReference
	SrvKeyspace
	CellInfo
*/
package topodata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KeyspaceIdType describes the type of the sharding key for a
// range-based sharded keyspace.
type KeyspaceIdType int32

const (
	// UNSET is the default value, when range-based sharding is not used.
	KeyspaceIdType_UNSET KeyspaceIdType = 0
	// UINT64 is when uint64 value is used.
	// This is represented as 'unsigned bigint' in mysql
	KeyspaceIdType_UINT64 KeyspaceIdType = 1
	// BYTES is when an array of bytes is used.
	// This is represented as 'varbinary' in mysql
	KeyspaceIdType_BYTES KeyspaceIdType = 2
)

var KeyspaceIdType_name = map[int32]string{
	0: "UNSET",
	1: "UINT64",
	2: "BYTES",
}
var KeyspaceIdType_value = map[string]int32{
	"UNSET":  0,
	"UINT64": 1,
	"BYTES":  2,
}

func (x KeyspaceIdType) String() string {
	return proto.EnumName(KeyspaceIdType_name, int32(x))
}
func (KeyspaceIdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TabletType represents the type of a given tablet.
type TabletType int32

const (
	// UNKNOWN is not a valid value.
	TabletType_UNKNOWN TabletType = 0
	// MASTER is the master server for the shard. Only MASTER allows DMLs.
	TabletType_MASTER TabletType = 1
	// REPLICA is a slave type. It is used to serve live traffic.
	// A REPLICA can be promoted to MASTER. A demoted MASTER will go to REPLICA.
	TabletType_REPLICA TabletType = 2
	// RDONLY (old name) / BATCH (new name) is used to serve traffic for
	// long-running jobs. It is a separate type from REPLICA so
	// long-running queries don't affect web-like traffic.
	TabletType_RDONLY TabletType = 3
	TabletType_BATCH  TabletType = 3
	// SPARE is a type of servers that cannot serve queries, but is available
	// in case an extra server is needed.
	TabletType_SPARE TabletType = 4
	// EXPERIMENTAL is like SPARE, except it can serve queries. This
	// type can be used for usages not planned by Vitess, like online
	// export to another storage engine.
	TabletType_EXPERIMENTAL TabletType = 5
	// BACKUP is the type a server goes to when taking a backup. No queries
	// can be served in BACKUP mode.
	TabletType_BACKUP TabletType = 6
	// RESTORE is the type a server uses when restoring a backup, at
	// startup time.  No queries can be served in RESTORE mode.
	TabletType_RESTORE TabletType = 7
	// DRAINED is the type a server goes into when used by Vitess tools
	// to perform an offline action. It is a serving type (as
	// the tools processes may need to run queries), but it's not used
	// to route queries from Vitess users. In this state,
	// this tablet is dedicated to the process that uses it.
	TabletType_DRAINED TabletType = 8
)

var TabletType_name = map[int32]string{
	0: "UNKNOWN",
	1: "MASTER",
	2: "REPLICA",
	3: "RDONLY",
	// Duplicate value: 3: "BATCH",
	4: "SPARE",
	5: "EXPERIMENTAL",
	6: "BACKUP",
	7: "RESTORE",
	8: "DRAINED",
}
var TabletType_value = map[string]int32{
	"UNKNOWN":      0,
	"MASTER":       1,
	"REPLICA":      2,
	"RDONLY":       3,
	"BATCH":        3,
	"SPARE":        4,
	"EXPERIMENTAL": 5,
	"BACKUP":       6,
	"RESTORE":      7,
	"DRAINED":      8,
}

func (x TabletType) String() string {
	return proto.EnumName(TabletType_name, int32(x))
}
func (TabletType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// KeyRange describes a range of sharding keys, when range-based
// sharding is used.
type KeyRange struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyRange) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *KeyRange) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

// TabletAlias is a globally unique tablet identifier.
type TabletAlias struct {
	// cell is the cell (or datacenter) the tablet is in
	Cell string `protobuf:"bytes,1,opt,name=cell" json:"cell,omitempty"`
	// uid is a unique id for this tablet within the shard
	// (this is the MySQL server id as well).
	Uid uint32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *TabletAlias) Reset()                    { *m = TabletAlias{} }
func (m *TabletAlias) String() string            { return proto.CompactTextString(m) }
func (*TabletAlias) ProtoMessage()               {}
func (*TabletAlias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TabletAlias) GetCell() string {
	if m != nil {
		return m.Cell
	}
	return ""
}

func (m *TabletAlias) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// Tablet represents information about a running instance of vttablet.
type Tablet struct {
	// alias is the unique name of the tablet.
	Alias *TabletAlias `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	// Fully qualified domain name of the host.
	Hostname string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	// Map of named ports. Normally this should include vt and grpc.
	// Going forward, the mysql port will be stored in mysql_port
	// instead of here.
	// For accessing mysql port, use topoproto.MysqlPort to fetch, and
	// topoproto.SetMysqlPort to set. These wrappers will ensure
	// legacy behavior is supported.
	PortMap map[string]int32 `protobuf:"bytes,4,rep,name=port_map,json=portMap" json:"port_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Keyspace name.
	Keyspace string `protobuf:"bytes,5,opt,name=keyspace" json:"keyspace,omitempty"`
	// Shard name. If range based sharding is used, it should match
	// key_range.
	Shard string `protobuf:"bytes,6,opt,name=shard" json:"shard,omitempty"`
	// If range based sharding is used, range for the tablet's shard.
	KeyRange *KeyRange `protobuf:"bytes,7,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// type is the current type of the tablet.
	Type TabletType `protobuf:"varint,8,opt,name=type,enum=topodata.TabletType" json:"type,omitempty"`
	// It this is set, it is used as the database name instead of the
	// normal "vt_" + keyspace.
	DbNameOverride string `protobuf:"bytes,9,opt,name=db_name_override,json=dbNameOverride" json:"db_name_override,omitempty"`
	// tablet tags
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// MySQL hostname.
	MysqlHostname string `protobuf:"bytes,12,opt,name=mysql_hostname,json=mysqlHostname" json:"mysql_hostname,omitempty"`
	// MySQL port. Use topoproto.MysqlPort and topoproto.SetMysqlPort
	// to access this variable. The functions provide support
	// for legacy behavior.
	MysqlPort int32 `protobuf:"varint,13,opt,name=mysql_port,json=mysqlPort" json:"mysql_port,omitempty"`
}

func (m *Tablet) Reset()                    { *m = Tablet{} }
func (m *Tablet) String() string            { return proto.CompactTextString(m) }
func (*Tablet) ProtoMessage()               {}
func (*Tablet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tablet) GetAlias() *TabletAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Tablet) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Tablet) GetPortMap() map[string]int32 {
	if m != nil {
		return m.PortMap
	}
	return nil
}

func (m *Tablet) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

func (m *Tablet) GetShard() string {
	if m != nil {
		return m.Shard
	}
	return ""
}

func (m *Tablet) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Tablet) GetType() TabletType {
	if m != nil {
		return m.Type
	}
	return TabletType_UNKNOWN
}

func (m *Tablet) GetDbNameOverride() string {
	if m != nil {
		return m.DbNameOverride
	}
	return ""
}

func (m *Tablet) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Tablet) GetMysqlHostname() string {
	if m != nil {
		return m.MysqlHostname
	}
	return ""
}

func (m *Tablet) GetMysqlPort() int32 {
	if m != nil {
		return m.MysqlPort
	}
	return 0
}

// A Shard contains data about a subset of the data whithin a keyspace.
type Shard struct {
	// No lock is necessary to update this field, when for instance
	// TabletExternallyReparented updates this. However, we lock the
	// shard for reparenting operations (InitShardMaster,
	// PlannedReparentShard,EmergencyReparentShard), to guarantee
	// exclusive operation.
	MasterAlias *TabletAlias `protobuf:"bytes,1,opt,name=master_alias,json=masterAlias" json:"master_alias,omitempty"`
	// key_range is the KeyRange for this shard. It can be unset if:
	// - we are not using range-based sharding in this shard.
	// - the shard covers the entire keyrange.
	// This must match the shard name based on our other conventions, but
	// helpful to have it decomposed here.
	// Once set at creation time, it is never changed.
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// served_types has at most one entry per TabletType
	// The keyspace lock is always taken when changing this.
	ServedTypes []*Shard_ServedType `protobuf:"bytes,3,rep,name=served_types,json=servedTypes" json:"served_types,omitempty"`
	// SourceShards is the list of shards we're replicating from,
	// using filtered replication.
	// The keyspace lock is always taken when changing this.
	SourceShards []*Shard_SourceShard `protobuf:"bytes,4,rep,name=source_shards,json=sourceShards" json:"source_shards,omitempty"`
	// Cells is the list of cells that contain tablets for this shard.
	// No lock is necessary to update this field.
	Cells []string `protobuf:"bytes,5,rep,name=cells" json:"cells,omitempty"`
	// tablet_controls has at most one entry per TabletType.
	// The keyspace lock is always taken when changing this.
	TabletControls []*Shard_TabletControl `protobuf:"bytes,6,rep,name=tablet_controls,json=tabletControls" json:"tablet_controls,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Shard) GetMasterAlias() *TabletAlias {
	if m != nil {
		return m.MasterAlias
	}
	return nil
}

func (m *Shard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard) GetServedTypes() []*Shard_ServedType {
	if m != nil {
		return m.ServedTypes
	}
	return nil
}

func (m *Shard) GetSourceShards() []*Shard_SourceShard {
	if m != nil {
		return m.SourceShards
	}
	return nil
}

func (m *Shard) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Shard) GetTabletControls() []*Shard_TabletControl {
	if m != nil {
		return m.TabletControls
	}
	return nil
}

// ServedType is an entry in the served_types
type Shard_ServedType struct {
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *Shard_ServedType) Reset()                    { *m = Shard_ServedType{} }
func (m *Shard_ServedType) String() string            { return proto.CompactTextString(m) }
func (*Shard_ServedType) ProtoMessage()               {}
func (*Shard_ServedType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Shard_ServedType) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Shard_ServedType) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

// SourceShard represents a data source for filtered replication
// accross shards. When this is used in a destination shard, the master
// of that shard will run filtered replication.
type Shard_SourceShard struct {
	// Uid is the unique ID for this SourceShard object.
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// the source keyspace
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	// the source shard
	Shard string `protobuf:"bytes,3,opt,name=shard" json:"shard,omitempty"`
	// the source shard keyrange
	KeyRange *KeyRange `protobuf:"bytes,4,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// the source table list to replicate
	Tables []string `protobuf:"bytes,5,rep,name=tables" json:"tables,omitempty"`
}

func (m *Shard_SourceShard) Reset()                    { *m = Shard_SourceShard{} }
func (m *Shard_SourceShard) String() string            { return proto.CompactTextString(m) }
func (*Shard_SourceShard) ProtoMessage()               {}
func (*Shard_SourceShard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *Shard_SourceShard) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Shard_SourceShard) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

func (m *Shard_SourceShard) GetShard() string {
	if m != nil {
		return m.Shard
	}
	return ""
}

func (m *Shard_SourceShard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard_SourceShard) GetTables() []string {
	if m != nil {
		return m.Tables
	}
	return nil
}

// TabletControl controls tablet's behavior
type Shard_TabletControl struct {
	// which tablet type is affected
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// what to do
	DisableQueryService bool     `protobuf:"varint,3,opt,name=disable_query_service,json=disableQueryService" json:"disable_query_service,omitempty"`
	BlacklistedTables   []string `protobuf:"bytes,4,rep,name=blacklisted_tables,json=blacklistedTables" json:"blacklisted_tables,omitempty"`
}

func (m *Shard_TabletControl) Reset()                    { *m = Shard_TabletControl{} }
func (m *Shard_TabletControl) String() string            { return proto.CompactTextString(m) }
func (*Shard_TabletControl) ProtoMessage()               {}
func (*Shard_TabletControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 2} }

func (m *Shard_TabletControl) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Shard_TabletControl) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Shard_TabletControl) GetDisableQueryService() bool {
	if m != nil {
		return m.DisableQueryService
	}
	return false
}

func (m *Shard_TabletControl) GetBlacklistedTables() []string {
	if m != nil {
		return m.BlacklistedTables
	}
	return nil
}

// A Keyspace contains data about a keyspace.
type Keyspace struct {
	// name of the column used for sharding
	// empty if the keyspace is not sharded
	ShardingColumnName string `protobuf:"bytes,1,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	// type of the column used for sharding
	// UNSET if the keyspace is not sharded
	ShardingColumnType KeyspaceIdType `protobuf:"varint,2,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	// ServedFrom will redirect the appropriate traffic to
	// another keyspace.
	ServedFroms []*Keyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_froms,json=servedFroms" json:"served_froms,omitempty"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Keyspace) GetShardingColumnName() string {
	if m != nil {
		return m.ShardingColumnName
	}
	return ""
}

func (m *Keyspace) GetShardingColumnType() KeyspaceIdType {
	if m != nil {
		return m.ShardingColumnType
	}
	return KeyspaceIdType_UNSET
}

func (m *Keyspace) GetServedFroms() []*Keyspace_ServedFrom {
	if m != nil {
		return m.ServedFroms
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type Keyspace_ServedFrom struct {
	// the tablet type (key for the map)
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the cells to limit this to
	Cells []string `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *Keyspace_ServedFrom) Reset()                    { *m = Keyspace_ServedFrom{} }
func (m *Keyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*Keyspace_ServedFrom) ProtoMessage()               {}
func (*Keyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Keyspace_ServedFrom) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Keyspace_ServedFrom) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Keyspace_ServedFrom) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

// ShardReplication describes the MySQL replication relationships
// whithin a cell.
type ShardReplication struct {
	// Note there can be only one Node in this array
	// for a given tablet.
	Nodes []*ShardReplication_Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ShardReplication) Reset()                    { *m = ShardReplication{} }
func (m *ShardReplication) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication) ProtoMessage()               {}
func (*ShardReplication) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ShardReplication) GetNodes() []*ShardReplication_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Node describes a tablet instance within the cell
type ShardReplication_Node struct {
	TabletAlias *TabletAlias `protobuf:"bytes,1,opt,name=tablet_alias,json=tabletAlias" json:"tablet_alias,omitempty"`
}

func (m *ShardReplication_Node) Reset()                    { *m = ShardReplication_Node{} }
func (m *ShardReplication_Node) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication_Node) ProtoMessage()               {}
func (*ShardReplication_Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ShardReplication_Node) GetTabletAlias() *TabletAlias {
	if m != nil {
		return m.TabletAlias
	}
	return nil
}

// ShardReference is used as a pointer from a SrvKeyspace to a Shard
type ShardReference struct {
	// Copied from Shard.
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
}

func (m *ShardReference) Reset()                    { *m = ShardReference{} }
func (m *ShardReference) String() string            { return proto.CompactTextString(m) }
func (*ShardReference) ProtoMessage()               {}
func (*ShardReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ShardReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShardReference) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// SrvKeyspace is a rollup node for the keyspace itself.
type SrvKeyspace struct {
	// The partitions this keyspace is serving, per tablet type.
	Partitions []*SrvKeyspace_KeyspacePartition `protobuf:"bytes,1,rep,name=partitions" json:"partitions,omitempty"`
	// copied from Keyspace
	ShardingColumnName string                    `protobuf:"bytes,2,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	ShardingColumnType KeyspaceIdType            `protobuf:"varint,3,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	ServedFrom         []*SrvKeyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_from,json=servedFrom" json:"served_from,omitempty"`
}

func (m *SrvKeyspace) Reset()                    { *m = SrvKeyspace{} }
func (m *SrvKeyspace) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace) ProtoMessage()               {}
func (*SrvKeyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SrvKeyspace) GetPartitions() []*SrvKeyspace_KeyspacePartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *SrvKeyspace) GetShardingColumnName() string {
	if m != nil {
		return m.ShardingColumnName
	}
	return ""
}

func (m *SrvKeyspace) GetShardingColumnType() KeyspaceIdType {
	if m != nil {
		return m.ShardingColumnType
	}
	return KeyspaceIdType_UNSET
}

func (m *SrvKeyspace) GetServedFrom() []*SrvKeyspace_ServedFrom {
	if m != nil {
		return m.ServedFrom
	}
	return nil
}

type SrvKeyspace_KeyspacePartition struct {
	// The type this partition applies to.
	ServedType TabletType `protobuf:"varint,1,opt,name=served_type,json=servedType,enum=topodata.TabletType" json:"served_type,omitempty"`
	// List of non-overlapping continuous shards sorted by range.
	ShardReferences []*ShardReference `protobuf:"bytes,2,rep,name=shard_references,json=shardReferences" json:"shard_references,omitempty"`
}

func (m *SrvKeyspace_KeyspacePartition) Reset()         { *m = SrvKeyspace_KeyspacePartition{} }
func (m *SrvKeyspace_KeyspacePartition) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace_KeyspacePartition) ProtoMessage()    {}
func (*SrvKeyspace_KeyspacePartition) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *SrvKeyspace_KeyspacePartition) GetServedType() TabletType {
	if m != nil {
		return m.ServedType
	}
	return TabletType_UNKNOWN
}

func (m *SrvKeyspace_KeyspacePartition) GetShardReferences() []*ShardReference {
	if m != nil {
		return m.ShardReferences
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type SrvKeyspace_ServedFrom struct {
	// the tablet type
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *SrvKeyspace_ServedFrom) Reset()                    { *m = SrvKeyspace_ServedFrom{} }
func (m *SrvKeyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace_ServedFrom) ProtoMessage()               {}
func (*SrvKeyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

func (m *SrvKeyspace_ServedFrom) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *SrvKeyspace_ServedFrom) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

// CellInfo contains information about a cell. CellInfo objects are
// stored in the global topology server, and describe how to reach
// local topology servers.
type CellInfo struct {
	// ServerAddress contains the address of the server for the cell.
	// The syntax of this field is topology implementation specific.
	// For instance, for Zookeeper, it is a comma-separated list of
	// server addresses.
	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress" json:"server_address,omitempty"`
	// Root is the path to store data in. It is only used when talking
	// to server_address.
	Root string `protobuf:"bytes,2,opt,name=root" json:"root,omitempty"`
}

func (m *CellInfo) Reset()                    { *m = CellInfo{} }
func (m *CellInfo) String() string            { return proto.CompactTextString(m) }
func (*CellInfo) ProtoMessage()               {}
func (*CellInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CellInfo) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *CellInfo) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "topodata.KeyRange")
	proto.RegisterType((*TabletAlias)(nil), "topodata.TabletAlias")
	proto.RegisterType((*Tablet)(nil), "topodata.Tablet")
	proto.RegisterType((*Shard)(nil), "topodata.Shard")
	proto.RegisterType((*Shard_ServedType)(nil), "topodata.Shard.ServedType")
	proto.RegisterType((*Shard_SourceShard)(nil), "topodata.Shard.SourceShard")
	proto.RegisterType((*Shard_TabletControl)(nil), "topodata.Shard.TabletControl")
	proto.RegisterType((*Keyspace)(nil), "topodata.Keyspace")
	proto.RegisterType((*Keyspace_ServedFrom)(nil), "topodata.Keyspace.ServedFrom")
	proto.RegisterType((*ShardReplication)(nil), "topodata.ShardReplication")
	proto.RegisterType((*ShardReplication_Node)(nil), "topodata.ShardReplication.Node")
	proto.RegisterType((*ShardReference)(nil), "topodata.ShardReference")
	proto.RegisterType((*SrvKeyspace)(nil), "topodata.SrvKeyspace")
	proto.RegisterType((*SrvKeyspace_KeyspacePartition)(nil), "topodata.SrvKeyspace.KeyspacePartition")
	proto.RegisterType((*SrvKeyspace_ServedFrom)(nil), "topodata.SrvKeyspace.ServedFrom")
	proto.RegisterType((*CellInfo)(nil), "topodata.CellInfo")
	proto.RegisterEnum("topodata.KeyspaceIdType", KeyspaceIdType_name, KeyspaceIdType_value)
	proto.RegisterEnum("topodata.TabletType", TabletType_name, TabletType_value)
}

func init() { proto.RegisterFile("topodata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xe2, 0x46,
	0x10, 0xaf, 0xc1, 0x10, 0x18, 0x03, 0xe7, 0x6c, 0x73, 0x95, 0xe5, 0xea, 0x54, 0x84, 0x54, 0x15,
	0x5d, 0x55, 0x5a, 0x71, 0xbd, 0x36, 0x3a, 0xa9, 0x52, 0x08, 0xf1, 0xf5, 0xc8, 0x1f, 0x42, 0x17,
	0xa2, 0x36, 0x4f, 0x96, 0x83, 0x37, 0x39, 0x2b, 0xc6, 0xf6, 0xed, 0x2e, 0x91, 0xf8, 0x0c, 0xf7,
	0xd0, 0x7b, 0xee, 0x37, 0xe9, 0x53, 0x1f, 0xfb, 0xb5, 0xaa, 0xdd, 0xb5, 0xc1, 0x90, 0x26, 0xcd,
	0x55, 0x79, 0xca, 0xcc, 0xce, 0x1f, 0xcf, 0xfc, 0xe6, 0x37, 0x13, 0xa0, 0xc1, 0xe3, 0x24, 0xf6,
	0x3d, 0xee, 0x75, 0x12, 0x1a, 0xf3, 0x18, 0x55, 0x32, 0xbd, 0xd5, 0x85, 0xca, 0x11, 0x59, 0x60,
	0x2f, 0xba, 0x22, 0x68, 0x07, 0x4a, 0x8c, 0x7b, 0x94, 0x5b, 0x5a, 0x53, 0x6b, 0xd7, 0xb0, 0x52,
	0x90, 0x09, 0x45, 0x12, 0xf9, 0x56, 0x41, 0xbe, 0x09, 0xb1, 0xf5, 0x02, 0x8c, 0x89, 0x77, 0x11,
	0x12, 0xde, 0x0b, 0x03, 0x8f, 0x21, 0x04, 0xfa, 0x94, 0x84, 0xa1, 0x8c, 0xaa, 0x62, 0x29, 0x8b,
	0xa0, 0x79, 0xa0, 0x82, 0xea, 0x58, 0x88, 0xad, 0x3f, 0x75, 0x28, 0xab, 0x28, 0xf4, 0x35, 0x94,
	0x3c, 0x11, 0x29, 0x23, 0x8c, 0xee, 0xd3, 0xce, 0xb2, 0xba, 0x5c, 0x5a, 0xac, 0x7c, 0x90, 0x0d,
	0x95, 0xb7, 0x31, 0xe3, 0x91, 0x37, 0x23, 0x32, 0x5d, 0x15, 0x2f, 0x75, 0xb4, 0x0b, 0x95, 0x24,
	0xa6, 0xdc, 0x9d, 0x79, 0x89, 0xa5, 0x37, 0x8b, 0x6d, 0xa3, 0xfb, 0x6c, 0x33, 0x57, 0x67, 0x14,
	0x53, 0x7e, 0xe2, 0x25, 0x4e, 0xc4, 0xe9, 0x02, 0x6f, 0x25, 0x4a, 0x13, 0x59, 0xaf, 0xc9, 0x82,
	0x25, 0xde, 0x94, 0x58, 0x25, 0x95, 0x35, 0xd3, 0x25, 0x0c, 0x6f, 0x3d, 0xea, 0x5b, 0x65, 0x69,
	0x50, 0x0a, 0xfa, 0x16, 0xaa, 0xd7, 0x64, 0xe1, 0x52, 0x81, 0x94, 0xb5, 0x25, 0x0b, 0x47, 0xab,
	0x8f, 0x65, 0x18, 0xca, 0x34, 0x0a, 0xcd, 0x36, 0xe8, 0x7c, 0x91, 0x10, 0xab, 0xd2, 0xd4, 0xda,
	0x8d, 0xee, 0xce, 0x66, 0x61, 0x93, 0x45, 0x42, 0xb0, 0xf4, 0x40, 0x6d, 0x30, 0xfd, 0x0b, 0x57,
	0x74, 0xe4, 0xc6, 0x37, 0x84, 0xd2, 0xc0, 0x27, 0x56, 0x55, 0x7e, 0xbb, 0xe1, 0x5f, 0x0c, 0xbd,
	0x19, 0x39, 0x4d, 0x5f, 0x51, 0x07, 0x74, 0xee, 0x5d, 0x31, 0x0b, 0x64, 0xb3, 0xf6, 0xad, 0x66,
	0x27, 0xde, 0x15, 0x53, 0x9d, 0x4a, 0x3f, 0xf4, 0x25, 0x34, 0x66, 0x0b, 0xf6, 0x2e, 0x74, 0x97,
	0x10, 0xd6, 0x64, 0xde, 0xba, 0x7c, 0x7d, 0x93, 0xe1, 0xf8, 0x0c, 0x40, 0xb9, 0x09, 0x78, 0xac,
	0x7a, 0x53, 0x6b, 0x97, 0x70, 0x55, 0xbe, 0x08, 0xf4, 0xec, 0x57, 0x50, 0xcb, 0xa3, 0x28, 0x86,
	0x7b, 0x4d, 0x16, 0xe9, 0xbc, 0x85, 0x28, 0x20, 0xbb, 0xf1, 0xc2, 0xb9, 0x9a, 0x50, 0x09, 0x2b,
	0xe5, 0x55, 0x61, 0x57, 0xb3, 0x7f, 0x84, 0xea, 0xb2, 0xa8, 0xff, 0x0a, 0xac, 0xe6, 0x02, 0x0f,
	0xf5, 0x4a, 0xd1, 0xd4, 0x0f, 0xf5, 0x8a, 0x61, 0xd6, 0x5a, 0xef, 0xcb, 0x50, 0x1a, 0xcb, 0x29,
	0xec, 0x42, 0x6d, 0xe6, 0x31, 0x4e, 0xa8, 0xfb, 0x00, 0x06, 0x19, 0xca, 0x55, 0xb1, 0x74, 0x6d,
	0x7e, 0x85, 0x07, 0xcc, 0xef, 0x27, 0xa8, 0x31, 0x42, 0x6f, 0x88, 0xef, 0x8a, 0x21, 0x31, 0xab,
	0xb8, 0x89, 0xb9, 0xac, 0xa8, 0x33, 0x96, 0x3e, 0x72, 0x9a, 0x06, 0x5b, 0xca, 0x0c, 0xed, 0x41,
	0x9d, 0xc5, 0x73, 0x3a, 0x25, 0xae, 0xe4, 0x0f, 0x4b, 0x09, 0xfa, 0xf9, 0xad, 0x78, 0xe9, 0x24,
	0x65, 0x5c, 0x63, 0x2b, 0x85, 0x09, 0x6c, 0xc4, 0x2e, 0x31, 0xab, 0xd4, 0x2c, 0x0a, 0x6c, 0xa4,
	0x82, 0x5e, 0xc3, 0x13, 0x2e, 0x7b, 0x74, 0xa7, 0x71, 0xc4, 0x69, 0x1c, 0x32, 0xab, 0xbc, 0x49,
	0x7d, 0x95, 0x59, 0x41, 0xd1, 0x57, 0x5e, 0xb8, 0xc1, 0xf3, 0x2a, 0xb3, 0xcf, 0x01, 0x56, 0xa5,
	0xa3, 0x97, 0x60, 0xa4, 0x59, 0x25, 0x67, 0xb5, 0x7b, 0x38, 0x0b, 0x7c, 0x29, 0xaf, 0x4a, 0x2c,
	0xe4, 0x4a, 0xb4, 0xff, 0xd0, 0xc0, 0xc8, 0xb5, 0x95, 0x1d, 0x03, 0x6d, 0x79, 0x0c, 0xd6, 0xd6,
	0xaf, 0x70, 0xd7, 0xfa, 0x15, 0xef, 0x5c, 0x3f, 0xfd, 0x01, 0xe3, 0xfb, 0x0c, 0xca, 0xb2, 0xd0,
	0x0c, 0xbe, 0x54, 0xb3, 0xff, 0xd2, 0xa0, 0xbe, 0x86, 0xcc, 0xa3, 0xf6, 0x8e, 0xba, 0xf0, 0xd4,
	0x0f, 0x98, 0xf0, 0x72, 0xdf, 0xcd, 0x09, 0x5d, 0xb8, 0x82, 0x13, 0xc1, 0x94, 0xc8, 0x6e, 0x2a,
	0xf8, 0xd3, 0xd4, 0xf8, 0x8b, 0xb0, 0x8d, 0x95, 0x09, 0x7d, 0x03, 0xe8, 0x22, 0xf4, 0xa6, 0xd7,
	0x61, 0xc0, 0xb8, 0xa0, 0x9b, 0x2a, 0x5b, 0x97, 0x69, 0xb7, 0x73, 0x16, 0x59, 0x08, 0x6b, 0xfd,
	0x5d, 0x90, 0x37, 0x5b, 0xa1, 0xf5, 0x1d, 0xec, 0x48, 0x80, 0x82, 0xe8, 0xca, 0x9d, 0xc6, 0xe1,
	0x7c, 0x16, 0xc9, 0x43, 0x92, 0xee, 0x18, 0xca, 0x6c, 0x7d, 0x69, 0x12, 0xb7, 0x04, 0x1d, 0xde,
	0x8e, 0x90, 0x7d, 0x17, 0x64, 0xdf, 0xd6, 0x1a, 0xa8, 0xf2, 0x1b, 0x03, 0xc5, 0xee, 0x8d, 0x5c,
	0x12, 0x83, 0xbd, 0xe5, 0x8e, 0x5c, 0xd2, 0x78, 0xc6, 0x6e, 0x1f, 0xe1, 0x2c, 0x47, 0xba, 0x26,
	0xaf, 0x69, 0x3c, 0xcb, 0xd6, 0x44, 0xc8, 0xcc, 0x9e, 0x67, 0x34, 0x14, 0xea, 0xe3, 0x8e, 0x22,
	0x4f, 0xb2, 0xe2, 0x3a, 0xc9, 0xd4, 0x75, 0x69, 0xbd, 0xd7, 0xc0, 0x54, 0x9b, 0x47, 0x92, 0x30,
	0x98, 0x7a, 0x3c, 0x88, 0x23, 0xf4, 0x12, 0x4a, 0x51, 0xec, 0x13, 0x71, 0x5b, 0x44, 0x33, 0x5f,
	0x6c, 0xac, 0x55, 0xce, 0xb5, 0x33, 0x8c, 0x7d, 0x82, 0x95, 0xb7, 0xbd, 0x07, 0xba, 0x50, 0xc5,
	0x85, 0x4a, 0x5b, 0x78, 0xc8, 0x85, 0xe2, 0x2b, 0xa5, 0x75, 0x06, 0x8d, 0xf4, 0x0b, 0x97, 0x84,
	0x92, 0x68, 0x4a, 0xc4, 0x7f, 0xd6, 0xdc, 0x30, 0xa5, 0xfc, 0xd1, 0x77, 0xac, 0xf5, 0x41, 0x07,
	0x63, 0x4c, 0x6f, 0x96, 0x8c, 0xf9, 0x19, 0x20, 0xf1, 0x28, 0x0f, 0x44, 0x07, 0x59, 0x93, 0x5f,
	0xe5, 0x9a, 0x5c, 0xb9, 0x2e, 0xa7, 0x37, 0xca, 0xfc, 0x71, 0x2e, 0xf4, 0x4e, 0xea, 0x15, 0x3e,
	0x9a, 0x7a, 0xc5, 0xff, 0x41, 0xbd, 0x1e, 0x18, 0x39, 0xea, 0xa5, 0xcc, 0x6b, 0xfe, 0x7b, 0x1f,
	0x39, 0xf2, 0xc1, 0x8a, 0x7c, 0xf6, 0xef, 0x1a, 0x6c, 0xdf, 0x6a, 0x51, 0x70, 0x30, 0x77, 0xf7,
	0xef, 0xe7, 0xe0, 0xea, 0xe0, 0xa3, 0x3e, 0x98, 0xb2, 0x4a, 0x97, 0x66, 0xe3, 0x53, 0x74, 0x34,
	0xf2, 0x7d, 0xad, 0xcf, 0x17, 0x3f, 0x61, 0x6b, 0x3a, 0xb3, 0xdd, 0xc7, 0xd8, 0x86, 0x7b, 0x8e,
	0xeb, 0xa1, 0x5e, 0x29, 0x99, 0xe5, 0x96, 0x03, 0x95, 0x3e, 0x09, 0xc3, 0x41, 0x74, 0x19, 0x8b,
	0x9f, 0x08, 0xb2, 0x0b, 0xea, 0x7a, 0xbe, 0x4f, 0x09, 0x63, 0x29, 0xdb, 0xea, 0xea, 0xb5, 0xa7,
	0x1e, 0x05, 0x15, 0x69, 0x1c, 0xf3, 0x34, 0xa1, 0x94, 0x9f, 0x77, 0xa1, 0xb1, 0x3e, 0x28, 0x54,
	0x85, 0xd2, 0xd9, 0x70, 0xec, 0x4c, 0xcc, 0x4f, 0x10, 0x40, 0xf9, 0x6c, 0x30, 0x9c, 0xfc, 0xf0,
	0xbd, 0xa9, 0x89, 0xe7, 0xfd, 0xf3, 0x89, 0x33, 0x36, 0x0b, 0xcf, 0x3f, 0x68, 0x00, 0xab, 0xba,
	0x91, 0x01, 0x5b, 0x67, 0xc3, 0xa3, 0xe1, 0xe9, 0xaf, 0x43, 0x15, 0x72, 0xd2, 0x1b, 0x4f, 0x1c,
	0x6c, 0x6a, 0xc2, 0x80, 0x9d, 0xd1, 0xf1, 0xa0, 0xdf, 0x33, 0x0b, 0xc2, 0x80, 0x0f, 0x4e, 0x87,
	0xc7, 0xe7, 0x66, 0x51, 0xe6, 0xea, 0x4d, 0xfa, 0x6f, 0x94, 0x38, 0x1e, 0xf5, 0xb0, 0x63, 0xea,
	0xc8, 0x84, 0x9a, 0xf3, 0xdb, 0xc8, 0xc1, 0x83, 0x13, 0x67, 0x38, 0xe9, 0x1d, 0x9b, 0x25, 0x11,
	0xb3, 0xdf, 0xeb, 0x1f, 0x9d, 0x8d, 0xcc, 0xb2, 0x4a, 0x36, 0x9e, 0x9c, 0x62, 0xc7, 0xdc, 0x12,
	0xca, 0x01, 0xee, 0x0d, 0x86, 0xce, 0x81, 0x59, 0xb1, 0x0b, 0xa6, 0xb6, 0xbf, 0x0d, 0x4f, 0x82,
	0xb8, 0x73, 0x13, 0x70, 0xc2, 0x98, 0xfa, 0x7d, 0x7c, 0x51, 0x96, 0x7f, 0x5e, 0xfc, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xba, 0xda, 0xa7, 0xb1, 0x38, 0x0b, 0x00, 0x00,
}
