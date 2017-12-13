# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from session.actions import access_pb2 as session_dot_actions_dot_access__pb2


class SessionStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.createSession = channel.unary_unary(
        '/session.Session/createSession',
        request_serializer=session_dot_actions_dot_access__pb2.CreateSessoinRequest.SerializeToString,
        response_deserializer=session_dot_actions_dot_access__pb2.CreateSessionReply.FromString,
        )
    self.getSession = channel.unary_unary(
        '/session.Session/getSession',
        request_serializer=session_dot_actions_dot_access__pb2.GetSessionRequest.SerializeToString,
        response_deserializer=session_dot_actions_dot_access__pb2.GetSessionReply.FromString,
        )
    self.updateSession = channel.unary_unary(
        '/session.Session/updateSession',
        request_serializer=session_dot_actions_dot_access__pb2.UpdateSessionRequest.SerializeToString,
        response_deserializer=session_dot_actions_dot_access__pb2.UpdateSessionReply.FromString,
        )
    self.deleteSession = channel.unary_unary(
        '/session.Session/deleteSession',
        request_serializer=session_dot_actions_dot_access__pb2.DeleteSessionRequest.SerializeToString,
        response_deserializer=session_dot_actions_dot_access__pb2.DeleteSessionReply.FromString,
        )


class SessionServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def createSession(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getSession(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def updateSession(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteSession(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SessionServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'createSession': grpc.unary_unary_rpc_method_handler(
          servicer.createSession,
          request_deserializer=session_dot_actions_dot_access__pb2.CreateSessoinRequest.FromString,
          response_serializer=session_dot_actions_dot_access__pb2.CreateSessionReply.SerializeToString,
      ),
      'getSession': grpc.unary_unary_rpc_method_handler(
          servicer.getSession,
          request_deserializer=session_dot_actions_dot_access__pb2.GetSessionRequest.FromString,
          response_serializer=session_dot_actions_dot_access__pb2.GetSessionReply.SerializeToString,
      ),
      'updateSession': grpc.unary_unary_rpc_method_handler(
          servicer.updateSession,
          request_deserializer=session_dot_actions_dot_access__pb2.UpdateSessionRequest.FromString,
          response_serializer=session_dot_actions_dot_access__pb2.UpdateSessionReply.SerializeToString,
      ),
      'deleteSession': grpc.unary_unary_rpc_method_handler(
          servicer.deleteSession,
          request_deserializer=session_dot_actions_dot_access__pb2.DeleteSessionRequest.FromString,
          response_serializer=session_dot_actions_dot_access__pb2.DeleteSessionReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'session.Session', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))