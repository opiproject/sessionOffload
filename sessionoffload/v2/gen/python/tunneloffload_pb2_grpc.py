# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import tunneloffload_pb2 as tunneloffload__pb2


class IpTunnelServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Capabilities = channel.unary_unary(
                '/tunneloffload.v1alpha1.IpTunnelService/Capabilities',
                request_serializer=tunneloffload__pb2.CapabilityRequest.SerializeToString,
                response_deserializer=tunneloffload__pb2.CapabilityResponse.FromString,
                )
        self.CreateIpTunnel = channel.stream_unary(
                '/tunneloffload.v1alpha1.IpTunnelService/CreateIpTunnel',
                request_serializer=tunneloffload__pb2.IpTunnelRequest.SerializeToString,
                response_deserializer=tunneloffload__pb2.CreateIpTunnelResponses.FromString,
                )
        self.GetIpTunnel = channel.unary_unary(
                '/tunneloffload.v1alpha1.IpTunnelService/GetIpTunnel',
                request_serializer=tunneloffload__pb2.TunnelId.SerializeToString,
                response_deserializer=tunneloffload__pb2.IpTunnelResponse.FromString,
                )
        self.GetIpTunnelStats = channel.unary_unary(
                '/tunneloffload.v1alpha1.IpTunnelService/GetIpTunnelStats',
                request_serializer=tunneloffload__pb2.TunnelId.SerializeToString,
                response_deserializer=tunneloffload__pb2.IpTunnelStatsResponse.FromString,
                )
        self.GetAllIpTunnels = channel.unary_stream(
                '/tunneloffload.v1alpha1.IpTunnelService/GetAllIpTunnels',
                request_serializer=tunneloffload__pb2.TunnelRequestArgs.SerializeToString,
                response_deserializer=tunneloffload__pb2.IpTunnelResponses.FromString,
                )
        self.GetAllIpTunnelsStats = channel.unary_stream(
                '/tunneloffload.v1alpha1.IpTunnelService/GetAllIpTunnelsStats',
                request_serializer=tunneloffload__pb2.TunnelRequestArgs.SerializeToString,
                response_deserializer=tunneloffload__pb2.IpTunnelStatsResponses.FromString,
                )


class IpTunnelServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Capabilities(self, request, context):
        """Get which capabilities are available while using the
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateIpTunnel(self, request_iterator, context):
        """Creation of IP Tunnel
        This API should be generic and allow creations of many IP tunnels
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetIpTunnel(self, request, context):
        """Getting a tunnel by it's ID
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetIpTunnelStats(self, request, context):
        """Getting a tunnel by it's ID
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllIpTunnels(self, request, context):
        """Getting all the ipTunnels currently configured
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllIpTunnelsStats(self, request, context):
        """Get all the iptunnels stats responses
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_IpTunnelServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Capabilities': grpc.unary_unary_rpc_method_handler(
                    servicer.Capabilities,
                    request_deserializer=tunneloffload__pb2.CapabilityRequest.FromString,
                    response_serializer=tunneloffload__pb2.CapabilityResponse.SerializeToString,
            ),
            'CreateIpTunnel': grpc.stream_unary_rpc_method_handler(
                    servicer.CreateIpTunnel,
                    request_deserializer=tunneloffload__pb2.IpTunnelRequest.FromString,
                    response_serializer=tunneloffload__pb2.CreateIpTunnelResponses.SerializeToString,
            ),
            'GetIpTunnel': grpc.unary_unary_rpc_method_handler(
                    servicer.GetIpTunnel,
                    request_deserializer=tunneloffload__pb2.TunnelId.FromString,
                    response_serializer=tunneloffload__pb2.IpTunnelResponse.SerializeToString,
            ),
            'GetIpTunnelStats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetIpTunnelStats,
                    request_deserializer=tunneloffload__pb2.TunnelId.FromString,
                    response_serializer=tunneloffload__pb2.IpTunnelStatsResponse.SerializeToString,
            ),
            'GetAllIpTunnels': grpc.unary_stream_rpc_method_handler(
                    servicer.GetAllIpTunnels,
                    request_deserializer=tunneloffload__pb2.TunnelRequestArgs.FromString,
                    response_serializer=tunneloffload__pb2.IpTunnelResponses.SerializeToString,
            ),
            'GetAllIpTunnelsStats': grpc.unary_stream_rpc_method_handler(
                    servicer.GetAllIpTunnelsStats,
                    request_deserializer=tunneloffload__pb2.TunnelRequestArgs.FromString,
                    response_serializer=tunneloffload__pb2.IpTunnelStatsResponses.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'tunneloffload.v1alpha1.IpTunnelService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class IpTunnelService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Capabilities(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tunneloffload.v1alpha1.IpTunnelService/Capabilities',
            tunneloffload__pb2.CapabilityRequest.SerializeToString,
            tunneloffload__pb2.CapabilityResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateIpTunnel(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/tunneloffload.v1alpha1.IpTunnelService/CreateIpTunnel',
            tunneloffload__pb2.IpTunnelRequest.SerializeToString,
            tunneloffload__pb2.CreateIpTunnelResponses.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetIpTunnel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tunneloffload.v1alpha1.IpTunnelService/GetIpTunnel',
            tunneloffload__pb2.TunnelId.SerializeToString,
            tunneloffload__pb2.IpTunnelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetIpTunnelStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tunneloffload.v1alpha1.IpTunnelService/GetIpTunnelStats',
            tunneloffload__pb2.TunnelId.SerializeToString,
            tunneloffload__pb2.IpTunnelStatsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllIpTunnels(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/tunneloffload.v1alpha1.IpTunnelService/GetAllIpTunnels',
            tunneloffload__pb2.TunnelRequestArgs.SerializeToString,
            tunneloffload__pb2.IpTunnelResponses.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllIpTunnelsStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/tunneloffload.v1alpha1.IpTunnelService/GetAllIpTunnelsStats',
            tunneloffload__pb2.TunnelRequestArgs.SerializeToString,
            tunneloffload__pb2.IpTunnelStatsResponses.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
