import grpc

from netograph.dsetapi import dset_pb2_grpc
from netograph.userapi import user_pb2_grpc


def connect_dset(token):
    channel = grpc.secure_channel(
        'grpc.netograph.io:443',
        grpc.composite_channel_credentials(
            grpc.ssl_channel_credentials(),
            grpc.access_token_call_credentials(token),
        ),
        options=[
            ('grpc.ssl_target_name_override', "grpc.netograph.io"),
        ]
    )
    return dset_pb2_grpc.DsetStub(channel)


def connect_user(token):
    channel = grpc.secure_channel(
        'grpc.netograph.io:443',
        grpc.composite_channel_credentials(
            grpc.ssl_channel_credentials(),
            grpc.access_token_call_credentials(token),
        ),
        options=[
            ('grpc.ssl_target_name_override', "grpc.netograph.io"),
        ]
    )
    return user_pb2_grpc.UserStub(channel)