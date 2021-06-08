# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import parse_pb2 as parse__pb2


class ParserStub(object):
    """The Parser service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Parse = channel.unary_unary(
                '/parser.Parser/Parse',
                request_serializer=parse__pb2.ParseRequest.SerializeToString,
                response_deserializer=parse__pb2.ParseReply.FromString,
                )


class ParserServicer(object):
    """The Parser service definition.
    """

    def Parse(self, request, context):
        """Sends a Parse Request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ParserServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Parse': grpc.unary_unary_rpc_method_handler(
                    servicer.Parse,
                    request_deserializer=parse__pb2.ParseRequest.FromString,
                    response_serializer=parse__pb2.ParseReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'parser.Parser', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Parser(object):
    """The Parser service definition.
    """

    @staticmethod
    def Parse(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/parser.Parser/Parse',
            parse__pb2.ParseRequest.SerializeToString,
            parse__pb2.ParseReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
