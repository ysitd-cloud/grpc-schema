// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var deployer_actions_access_pb = require('../deployer/actions/access_pb.js');

function serialize_deployer_CreateApplicationReply(arg) {
  if (!(arg instanceof deployer_actions_access_pb.CreateApplicationReply)) {
    throw new Error('Expected argument of type deployer.CreateApplicationReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_CreateApplicationReply(buffer_arg) {
  return deployer_actions_access_pb.CreateApplicationReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_CreateApplicationRequest(arg) {
  if (!(arg instanceof deployer_actions_access_pb.CreateApplicationRequest)) {
    throw new Error('Expected argument of type deployer.CreateApplicationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_CreateApplicationRequest(buffer_arg) {
  return deployer_actions_access_pb.CreateApplicationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_GetApplicationByIdRequest(arg) {
  if (!(arg instanceof deployer_actions_access_pb.GetApplicationByIdRequest)) {
    throw new Error('Expected argument of type deployer.GetApplicationByIdRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_GetApplicationByIdRequest(buffer_arg) {
  return deployer_actions_access_pb.GetApplicationByIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_GetApplicationByIdResponse(arg) {
  if (!(arg instanceof deployer_actions_access_pb.GetApplicationByIdResponse)) {
    throw new Error('Expected argument of type deployer.GetApplicationByIdResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_GetApplicationByIdResponse(buffer_arg) {
  return deployer_actions_access_pb.GetApplicationByIdResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_ListApplicationsByUsernameReply(arg) {
  if (!(arg instanceof deployer_actions_access_pb.ListApplicationsByUsernameReply)) {
    throw new Error('Expected argument of type deployer.ListApplicationsByUsernameReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_ListApplicationsByUsernameReply(buffer_arg) {
  return deployer_actions_access_pb.ListApplicationsByUsernameReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_ListApplicationsByUsernameRequest(arg) {
  if (!(arg instanceof deployer_actions_access_pb.ListApplicationsByUsernameRequest)) {
    throw new Error('Expected argument of type deployer.ListApplicationsByUsernameRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_ListApplicationsByUsernameRequest(buffer_arg) {
  return deployer_actions_access_pb.ListApplicationsByUsernameRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_UpdateDeploymentImageReply(arg) {
  if (!(arg instanceof deployer_actions_access_pb.UpdateDeploymentImageReply)) {
    throw new Error('Expected argument of type deployer.UpdateDeploymentImageReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_UpdateDeploymentImageReply(buffer_arg) {
  return deployer_actions_access_pb.UpdateDeploymentImageReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_deployer_UpdateDeploymentImageRequest(arg) {
  if (!(arg instanceof deployer_actions_access_pb.UpdateDeploymentImageRequest)) {
    throw new Error('Expected argument of type deployer.UpdateDeploymentImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_deployer_UpdateDeploymentImageRequest(buffer_arg) {
  return deployer_actions_access_pb.UpdateDeploymentImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var DeployerService = exports.DeployerService = {
  listApplicationsByUsername: {
    path: '/deployer.Deployer/listApplicationsByUsername',
    requestStream: false,
    responseStream: false,
    requestType: deployer_actions_access_pb.ListApplicationsByUsernameRequest,
    responseType: deployer_actions_access_pb.ListApplicationsByUsernameReply,
    requestSerialize: serialize_deployer_ListApplicationsByUsernameRequest,
    requestDeserialize: deserialize_deployer_ListApplicationsByUsernameRequest,
    responseSerialize: serialize_deployer_ListApplicationsByUsernameReply,
    responseDeserialize: deserialize_deployer_ListApplicationsByUsernameReply,
  },
  createApplication: {
    path: '/deployer.Deployer/createApplication',
    requestStream: false,
    responseStream: false,
    requestType: deployer_actions_access_pb.CreateApplicationRequest,
    responseType: deployer_actions_access_pb.CreateApplicationReply,
    requestSerialize: serialize_deployer_CreateApplicationRequest,
    requestDeserialize: deserialize_deployer_CreateApplicationRequest,
    responseSerialize: serialize_deployer_CreateApplicationReply,
    responseDeserialize: deserialize_deployer_CreateApplicationReply,
  },
  updateDeploymentImage: {
    path: '/deployer.Deployer/updateDeploymentImage',
    requestStream: false,
    responseStream: false,
    requestType: deployer_actions_access_pb.UpdateDeploymentImageRequest,
    responseType: deployer_actions_access_pb.UpdateDeploymentImageReply,
    requestSerialize: serialize_deployer_UpdateDeploymentImageRequest,
    requestDeserialize: deserialize_deployer_UpdateDeploymentImageRequest,
    responseSerialize: serialize_deployer_UpdateDeploymentImageReply,
    responseDeserialize: deserialize_deployer_UpdateDeploymentImageReply,
  },
  getApplicationById: {
    path: '/deployer.Deployer/getApplicationById',
    requestStream: false,
    responseStream: false,
    requestType: deployer_actions_access_pb.GetApplicationByIdRequest,
    responseType: deployer_actions_access_pb.GetApplicationByIdResponse,
    requestSerialize: serialize_deployer_GetApplicationByIdRequest,
    requestDeserialize: deserialize_deployer_GetApplicationByIdRequest,
    responseSerialize: serialize_deployer_GetApplicationByIdResponse,
    responseDeserialize: deserialize_deployer_GetApplicationByIdResponse,
  },
};

exports.DeployerClient = grpc.makeGenericClientConstructor(DeployerService);
