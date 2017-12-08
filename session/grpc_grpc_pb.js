// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var session_actions_access_pb = require('../session/actions/access_pb.js');

function serialize_session_CreateSessionReply(arg) {
  if (!(arg instanceof session_actions_access_pb.CreateSessionReply)) {
    throw new Error('Expected argument of type session.CreateSessionReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_CreateSessionReply(buffer_arg) {
  return session_actions_access_pb.CreateSessionReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_CreateSessoinRequest(arg) {
  if (!(arg instanceof session_actions_access_pb.CreateSessoinRequest)) {
    throw new Error('Expected argument of type session.CreateSessoinRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_CreateSessoinRequest(buffer_arg) {
  return session_actions_access_pb.CreateSessoinRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_DeleteSessionReply(arg) {
  if (!(arg instanceof session_actions_access_pb.DeleteSessionReply)) {
    throw new Error('Expected argument of type session.DeleteSessionReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_DeleteSessionReply(buffer_arg) {
  return session_actions_access_pb.DeleteSessionReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_DeleteSessionRequest(arg) {
  if (!(arg instanceof session_actions_access_pb.DeleteSessionRequest)) {
    throw new Error('Expected argument of type session.DeleteSessionRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_DeleteSessionRequest(buffer_arg) {
  return session_actions_access_pb.DeleteSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_GetSessionReply(arg) {
  if (!(arg instanceof session_actions_access_pb.GetSessionReply)) {
    throw new Error('Expected argument of type session.GetSessionReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_GetSessionReply(buffer_arg) {
  return session_actions_access_pb.GetSessionReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_GetSessionRequest(arg) {
  if (!(arg instanceof session_actions_access_pb.GetSessionRequest)) {
    throw new Error('Expected argument of type session.GetSessionRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_GetSessionRequest(buffer_arg) {
  return session_actions_access_pb.GetSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_UpdateSessionReply(arg) {
  if (!(arg instanceof session_actions_access_pb.UpdateSessionReply)) {
    throw new Error('Expected argument of type session.UpdateSessionReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_UpdateSessionReply(buffer_arg) {
  return session_actions_access_pb.UpdateSessionReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_session_UpdateSessionRequest(arg) {
  if (!(arg instanceof session_actions_access_pb.UpdateSessionRequest)) {
    throw new Error('Expected argument of type session.UpdateSessionRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_session_UpdateSessionRequest(buffer_arg) {
  return session_actions_access_pb.UpdateSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var SessionService = exports.SessionService = {
  createSession: {
    path: '/session.Session/createSession',
    requestStream: false,
    responseStream: false,
    requestType: session_actions_access_pb.CreateSessoinRequest,
    responseType: session_actions_access_pb.CreateSessionReply,
    requestSerialize: serialize_session_CreateSessoinRequest,
    requestDeserialize: deserialize_session_CreateSessoinRequest,
    responseSerialize: serialize_session_CreateSessionReply,
    responseDeserialize: deserialize_session_CreateSessionReply,
  },
  getSession: {
    path: '/session.Session/getSession',
    requestStream: false,
    responseStream: false,
    requestType: session_actions_access_pb.GetSessionRequest,
    responseType: session_actions_access_pb.GetSessionReply,
    requestSerialize: serialize_session_GetSessionRequest,
    requestDeserialize: deserialize_session_GetSessionRequest,
    responseSerialize: serialize_session_GetSessionReply,
    responseDeserialize: deserialize_session_GetSessionReply,
  },
  updateSession: {
    path: '/session.Session/updateSession',
    requestStream: false,
    responseStream: false,
    requestType: session_actions_access_pb.UpdateSessionRequest,
    responseType: session_actions_access_pb.UpdateSessionReply,
    requestSerialize: serialize_session_UpdateSessionRequest,
    requestDeserialize: deserialize_session_UpdateSessionRequest,
    responseSerialize: serialize_session_UpdateSessionReply,
    responseDeserialize: deserialize_session_UpdateSessionReply,
  },
  deleteSession: {
    path: '/session.Session/deleteSession',
    requestStream: false,
    responseStream: false,
    requestType: session_actions_access_pb.DeleteSessionRequest,
    responseType: session_actions_access_pb.DeleteSessionReply,
    requestSerialize: serialize_session_DeleteSessionRequest,
    requestDeserialize: deserialize_session_DeleteSessionRequest,
    responseSerialize: serialize_session_DeleteSessionReply,
    responseDeserialize: deserialize_session_DeleteSessionReply,
  },
};

exports.SessionClient = grpc.makeGenericClientConstructor(SessionService);
