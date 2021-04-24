from protocol_buffers.user_service_pb2 import GetUserRequest, GetUserReply
from protocol_buffers.user_service_pb2_grpc import *
import grpc


def run():
    channel = grpc.insecure_channel('172.30.79.172:10000')
    stub = UserServiceStub(channel)
    response = stub.GetUser(
        GetUserRequest(user_id=1)
    )
    print(f"Reply received. {response}")


if __name__ == '__main__':
    run()
