# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from onehub.v1 import users_pb2 as onehub_dot_v1_dot_users__pb2


class UserServiceStub(object):
    """*
    Service for operating on users
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateUser = channel.unary_unary(
                '/onehub.v1.UserService/CreateUser',
                request_serializer=onehub_dot_v1_dot_users__pb2.CreateUserRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.CreateUserResponse.FromString,
                )
        self.ListUsers = channel.unary_unary(
                '/onehub.v1.UserService/ListUsers',
                request_serializer=onehub_dot_v1_dot_users__pb2.ListUsersRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.ListUsersResponse.FromString,
                )
        self.GetUser = channel.unary_unary(
                '/onehub.v1.UserService/GetUser',
                request_serializer=onehub_dot_v1_dot_users__pb2.GetUserRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.GetUserResponse.FromString,
                )
        self.GetUsers = channel.unary_unary(
                '/onehub.v1.UserService/GetUsers',
                request_serializer=onehub_dot_v1_dot_users__pb2.GetUsersRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.GetUsersResponse.FromString,
                )
        self.DeleteUser = channel.unary_unary(
                '/onehub.v1.UserService/DeleteUser',
                request_serializer=onehub_dot_v1_dot_users__pb2.DeleteUserRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.DeleteUserResponse.FromString,
                )
        self.UpdateUser = channel.unary_unary(
                '/onehub.v1.UserService/UpdateUser',
                request_serializer=onehub_dot_v1_dot_users__pb2.UpdateUserRequest.SerializeToString,
                response_deserializer=onehub_dot_v1_dot_users__pb2.UpdateUserResponse.FromString,
                )


class UserServiceServicer(object):
    """*
    Service for operating on users
    """

    def CreateUser(self, request, context):
        """*
        Create a new sesssion
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUsers(self, request, context):
        """*
        List all users from a user.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUser(self, request, context):
        """*
        Get a particular user
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUsers(self, request, context):
        """*
        Batch get multiple users by ID
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteUser(self, request, context):
        """*
        Delete a particular user
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUser(self, request, context):
        """*
        Updates specific fields of a user
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUser,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.CreateUserRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.CreateUserResponse.SerializeToString,
            ),
            'ListUsers': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUsers,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.ListUsersRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.ListUsersResponse.SerializeToString,
            ),
            'GetUser': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUser,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.GetUserRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.GetUserResponse.SerializeToString,
            ),
            'GetUsers': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUsers,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.GetUsersRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.GetUsersResponse.SerializeToString,
            ),
            'DeleteUser': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteUser,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.DeleteUserRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.DeleteUserResponse.SerializeToString,
            ),
            'UpdateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateUser,
                    request_deserializer=onehub_dot_v1_dot_users__pb2.UpdateUserRequest.FromString,
                    response_serializer=onehub_dot_v1_dot_users__pb2.UpdateUserResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'onehub.v1.UserService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UserService(object):
    """*
    Service for operating on users
    """

    @staticmethod
    def CreateUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/CreateUser',
            onehub_dot_v1_dot_users__pb2.CreateUserRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.CreateUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListUsers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/ListUsers',
            onehub_dot_v1_dot_users__pb2.ListUsersRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.ListUsersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/GetUser',
            onehub_dot_v1_dot_users__pb2.GetUserRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.GetUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUsers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/GetUsers',
            onehub_dot_v1_dot_users__pb2.GetUsersRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.GetUsersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/DeleteUser',
            onehub_dot_v1_dot_users__pb2.DeleteUserRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.DeleteUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/onehub.v1.UserService/UpdateUser',
            onehub_dot_v1_dot_users__pb2.UpdateUserRequest.SerializeToString,
            onehub_dot_v1_dot_users__pb2.UpdateUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
