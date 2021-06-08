import parse_pb2_grpc
import parse_pb2
import parser
from concurrent import futures
import logging
import os

import grpc


class Parser(parse_pb2_grpc.ParserServicer):
    def Parse(self, request: parse_pb2.ParseRequest, context):
        return parse_pb2.ParseReply(result=parser.main(request.content))


def serve(environment):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    parse_pb2_grpc.add_ParserServicer_to_server(Parser(), server)
    if environment is not None:
        server.add_insecure_port('0.0.0.0:50051')
    else:
        server.add_insecure_port('127.0.0.1:50051')

    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    env = os.getenv('DOCKER_ENVIRONMENT')
    serve(env)
