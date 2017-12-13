// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var account_actions_account_pb = require('../account/actions/account_pb.js');

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
};

exports.AccountClient = grpc.makeGenericClientConstructor(AccountService);
