// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: agent.proto

package agentpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Agent_Ping_FullMethodName                  = "/agentpb.Agent/Ping"
	Agent_SubmitJob_FullMethodName             = "/agentpb.Agent/SubmitJob"
	Agent_StoreCredential_FullMethodName       = "/agentpb.Agent/StoreCredential"
	Agent_Sync_FullMethodName                  = "/agentpb.Agent/Sync"
	Agent_ClimonAppInstall_FullMethodName      = "/agentpb.Agent/ClimonAppInstall"
	Agent_ClimonAppDelete_FullMethodName       = "/agentpb.Agent/ClimonAppDelete"
	Agent_DeployerAppInstall_FullMethodName    = "/agentpb.Agent/DeployerAppInstall"
	Agent_DeployerAppDelete_FullMethodName     = "/agentpb.Agent/DeployerAppDelete"
	Agent_ClusterAdd_FullMethodName            = "/agentpb.Agent/ClusterAdd"
	Agent_ClusterDelete_FullMethodName         = "/agentpb.Agent/ClusterDelete"
	Agent_RepositoryAdd_FullMethodName         = "/agentpb.Agent/RepositoryAdd"
	Agent_RepositoryDelete_FullMethodName      = "/agentpb.Agent/RepositoryDelete"
	Agent_ProjectAdd_FullMethodName            = "/agentpb.Agent/ProjectAdd"
	Agent_ProjectDelete_FullMethodName         = "/agentpb.Agent/ProjectDelete"
	Agent_SyncApp_FullMethodName               = "/agentpb.Agent/SyncApp"
	Agent_GetClusterApps_FullMethodName        = "/agentpb.Agent/GetClusterApps"
	Agent_GetClusterAppLaunches_FullMethodName = "/agentpb.Agent/GetClusterAppLaunches"
	Agent_GetClusterAppConfig_FullMethodName   = "/agentpb.Agent/GetClusterAppConfig"
	Agent_GetClusterAppValues_FullMethodName   = "/agentpb.Agent/GetClusterAppValues"
	Agent_InstallApp_FullMethodName            = "/agentpb.Agent/InstallApp"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error)
	GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error)
	GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error)
	InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*JobResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Agent_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_SubmitJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error) {
	out := new(StoreCredentialResponse)
	err := c.cc.Invoke(ctx, Agent_StoreCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, Agent_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClimonAppInstall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClimonAppDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_DeployerAppInstall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_DeployerAppDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClusterAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClusterDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_RepositoryAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_RepositoryDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ProjectAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ProjectDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error) {
	out := new(SyncAppResponse)
	err := c.cc.Invoke(ctx, Agent_SyncApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error) {
	out := new(GetClusterAppsResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error) {
	out := new(GetClusterAppLaunchesResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppLaunches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error) {
	out := new(GetClusterAppConfigResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error) {
	out := new(GetClusterAppValuesResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_InstallApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SubmitJob(context.Context, *JobRequest) (*JobResponse, error)
	StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error)
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error)
	ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error)
	DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error)
	DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error)
	ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error)
	ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error)
	RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error)
	RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error)
	ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error)
	ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error)
	SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error)
	GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error)
	GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error)
	InstallApp(context.Context, *InstallAppRequest) (*JobResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServer) SubmitJob(context.Context, *JobRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedAgentServer) StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCredential not implemented")
}
func (UnimplementedAgentServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedAgentServer) ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppInstall not implemented")
}
func (UnimplementedAgentServer) ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppDelete not implemented")
}
func (UnimplementedAgentServer) DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppInstall not implemented")
}
func (UnimplementedAgentServer) DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppDelete not implemented")
}
func (UnimplementedAgentServer) ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterAdd not implemented")
}
func (UnimplementedAgentServer) ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterDelete not implemented")
}
func (UnimplementedAgentServer) RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryAdd not implemented")
}
func (UnimplementedAgentServer) RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryDelete not implemented")
}
func (UnimplementedAgentServer) ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectAdd not implemented")
}
func (UnimplementedAgentServer) ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectDelete not implemented")
}
func (UnimplementedAgentServer) SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncApp not implemented")
}
func (UnimplementedAgentServer) GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterApps not implemented")
}
func (UnimplementedAgentServer) GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppLaunches not implemented")
}
func (UnimplementedAgentServer) GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppConfig not implemented")
}
func (UnimplementedAgentServer) GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppValues not implemented")
}
func (UnimplementedAgentServer) InstallApp(context.Context, *InstallAppRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallApp not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SubmitJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SubmitJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StoreCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StoreCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_StoreCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StoreCredential(ctx, req.(*StoreCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClimonAppInstall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppInstall(ctx, req.(*ClimonInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClimonAppDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppDelete(ctx, req.(*ClimonDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeployerAppInstall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppInstall(ctx, req.(*ApplicationInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeployerAppDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppDelete(ctx, req.(*ApplicationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClusterAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterAdd(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClusterDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterDelete(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_RepositoryAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryAdd(ctx, req.(*RepositoryAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_RepositoryDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryDelete(ctx, req.(*RepositoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ProjectAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectAdd(ctx, req.(*ProjectAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ProjectDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectDelete(ctx, req.(*ProjectDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SyncApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SyncApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SyncApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SyncApp(ctx, req.(*SyncAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterApps(ctx, req.(*GetClusterAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppLaunches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppLaunchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppLaunches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, req.(*GetClusterAppLaunchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppConfig(ctx, req.(*GetClusterAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppValues(ctx, req.(*GetClusterAppValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_InstallApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).InstallApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_InstallApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).InstallApp(ctx, req.(*InstallAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
		{
			MethodName: "SubmitJob",
			Handler:    _Agent_SubmitJob_Handler,
		},
		{
			MethodName: "StoreCredential",
			Handler:    _Agent_StoreCredential_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Agent_Sync_Handler,
		},
		{
			MethodName: "ClimonAppInstall",
			Handler:    _Agent_ClimonAppInstall_Handler,
		},
		{
			MethodName: "ClimonAppDelete",
			Handler:    _Agent_ClimonAppDelete_Handler,
		},
		{
			MethodName: "DeployerAppInstall",
			Handler:    _Agent_DeployerAppInstall_Handler,
		},
		{
			MethodName: "DeployerAppDelete",
			Handler:    _Agent_DeployerAppDelete_Handler,
		},
		{
			MethodName: "ClusterAdd",
			Handler:    _Agent_ClusterAdd_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Agent_ClusterDelete_Handler,
		},
		{
			MethodName: "RepositoryAdd",
			Handler:    _Agent_RepositoryAdd_Handler,
		},
		{
			MethodName: "RepositoryDelete",
			Handler:    _Agent_RepositoryDelete_Handler,
		},
		{
			MethodName: "ProjectAdd",
			Handler:    _Agent_ProjectAdd_Handler,
		},
		{
			MethodName: "ProjectDelete",
			Handler:    _Agent_ProjectDelete_Handler,
		},
		{
			MethodName: "SyncApp",
			Handler:    _Agent_SyncApp_Handler,
		},
		{
			MethodName: "GetClusterApps",
			Handler:    _Agent_GetClusterApps_Handler,
		},
		{
			MethodName: "GetClusterAppLaunches",
			Handler:    _Agent_GetClusterAppLaunches_Handler,
		},
		{
			MethodName: "GetClusterAppConfig",
			Handler:    _Agent_GetClusterAppConfig_Handler,
		},
		{
			MethodName: "GetClusterAppValues",
			Handler:    _Agent_GetClusterAppValues_Handler,
		},
		{
			MethodName: "InstallApp",
			Handler:    _Agent_InstallApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
