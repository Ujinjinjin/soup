from user_service_pb2 import CreateUserRequest
from user_service_pb2_grpc import *
import grpc


def run():
    channel = grpc.insecure_channel('localhost:10000')
    stub = UserServiceStub(channel)
    response = stub.CreateUser(
        CreateUserRequest(
            email="ujinjinjin@outlook.com",
            first_name="Camille",
            last_name="Galladjov",
            middle_name="Amrullaevich"
        )
    )
    print(f"Reply received. User Id: {response.id}")


if __name__ == '__main__':
    run()
