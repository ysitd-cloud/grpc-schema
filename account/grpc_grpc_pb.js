// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var account_actions_account_pb = require('../account/actions/account_pb.js');

function serialize_account_GetTokenInfoReply(arg) {
  if (!(arg instanceof account_actions_account_pb.GetTokenInfoReply)) {
    throw new Error('Expected argument of type account.GetTokenInfoReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_GetTokenInfoReply(buffer_arg) {
  return account_actions_account_pb.GetTokenInfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_account_GetTokenInfoRequest(arg) {
  if (!(arg instanceof account_actions_account_pb.GetTokenInfoRequest)) {
    throw new Error('Expected argument of type account.GetTokenInfoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_GetTokenInfoRequest(buffer_arg) {
  return account_actions_account_pb.GetTokenInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_account_GetUserInfoReply(arg) {
  if (!(arg instanceof account_actions_account_pb.GetUserInfoReply)) {
    throw new Error('Expected argument of type account.GetUserInfoReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_GetUserInfoReply(buffer_arg) {
  return account_actions_account_pb.GetUserInfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_account_GetUserInfoRequest(arg) {
  if (!(arg instanceof account_actions_account_pb.GetUserInfoRequest)) {
    throw new Error('Expected argument of type account.GetUserInfoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_GetUserInfoRequest(buffer_arg) {
  return account_actions_account_pb.GetUserInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_account_ValidateUserReply(arg) {
  if (!(arg instanceof account_actions_account_pb.ValidateUserReply)) {
    throw new Error('Expected argument of type account.ValidateUserReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_ValidateUserReply(buffer_arg) {
  return account_actions_account_pb.ValidateUserReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_account_ValidateUserRequest(arg) {
  if (!(arg instanceof account_actions_account_pb.ValidateUserRequest)) {
    throw new Error('Expected argument of type account.ValidateUserRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_account_ValidateUserRequest(buffer_arg) {
  return account_actions_account_pb.ValidateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var AccountService = exports.AccountService = {
  validateUserPassword: {
    path: '/account.Account/validateUserPassword',
    requestStream: false,
    responseStream: false,
    requestType: account_actions_account_pb.ValidateUserRequest,
    responseType: account_actions_account_pb.ValidateUserReply,
    requestSerialize: serialize_account_ValidateUserRequest,
    requestDeserialize: deserialize_account_ValidateUserRequest,
    responseSerialize: serialize_account_ValidateUserReply,
    responseDeserialize: deserialize_account_ValidateUserReply,
  },
  getUserInfo: {
    path: '/account.Account/getUserInfo',
    requestStream: false,
    responseStream: false,
    requestType: account_actions_account_pb.GetUserInfoRequest,
    responseType: account_actions_account_pb.GetUserInfoReply,
    requestSerialize: serialize_account_GetUserInfoRequest,
    requestDeserialize: deserialize_account_GetUserInfoRequest,
    responseSerialize: serialize_account_GetUserInfoReply,
    responseDeserialize: deserialize_account_GetUserInfoReply,
  },
  getTokenInfo: {
    path: '/account.Account/getTokenInfo',
    requestStream: false,
    responseStream: false,
    requestType: account_actions_account_pb.GetTokenInfoRequest,
    responseType: account_actions_account_pb.GetTokenInfoReply,
    requestSerialize: serialize_account_GetTokenInfoRequest,
    requestDeserialize: deserialize_account_GetTokenInfoRequest,
    responseSerialize: serialize_account_GetTokenInfoReply,
    responseDeserialize: deserialize_account_GetTokenInfoReply,
  },
};

exports.AccountClient = grpc.makeGenericClientConstructor(AccountService);
