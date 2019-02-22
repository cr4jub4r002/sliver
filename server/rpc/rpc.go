package rpc

import (
	"sliver/server/core"
	"time"

	clientpb "sliver/protobuf/client"
	sliverpb "sliver/protobuf/sliver"
)

const (
	defaultTimeout = 30 * time.Second
)

// RPCResponse - Called with response data, mapped back to reqID
type RPCResponse func([]byte, error)

// RPCHandler - RPC handlers accept bytes and return bytes
type RPCHandler func([]byte, RPCResponse)
type TunnelHandler func(*core.Client, []byte, RPCResponse)

var (
	rpcHandlers = &map[uint32]RPCHandler{
		clientpb.MsgJobs: rpcJobs,
		clientpb.MsgMtls: rpcStartMTLSListener,
		clientpb.MsgDns:  rpcStartDNSListener,

		clientpb.MsgSessions:   rpcSessions,
		clientpb.MsgGenerate:   rpcGenerate,
		clientpb.MsgProfiles:   rpcProfiles,
		clientpb.MsgNewProfile: rpcNewProfile,

		clientpb.MsgMsf:       rpcMsf,
		clientpb.MsgMsfInject: rpcMsfInject,

		sliverpb.MsgPsReq:          rpcPs,
		sliverpb.MsgProcessDumpReq: rpcProcdump,

		sliverpb.MsgLsReq:       rpcLs,
		sliverpb.MsgRmReq:       rpcRm,
		sliverpb.MsgMkdirReq:    rpcMkdir,
		sliverpb.MsgCdReq:       rpcCd,
		sliverpb.MsgPwdReq:      rpcPwd,
		sliverpb.MsgDownloadReq: rpcDownload,
		sliverpb.MsgUploadReq:   rpcUpload,

		sliverpb.MsgShellReq: rpcShell,
	}

	tunHandlers = &map[uint32]TunnelHandler{
		clientpb.MsgTunnelCreate: tunnelCreate,
	}
)

// GetRPCHandlers - Returns a map of server-side msg handlers
func GetRPCHandlers() *map[uint32]RPCHandler {
	return rpcHandlers
}

// GetTunnelHandlers - Returns a map of tunnel handlers
func GetTunnelHandlers() *map[uint32]TunnelHandler {
	return tunHandlers
}