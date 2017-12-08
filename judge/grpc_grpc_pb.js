// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var judge_models_types_pb = require('../judge/models/types_pb.js');
var judge_actions_judge_pb = require('../judge/actions/judge_pb.js');

function serialize_judge_AccessReply(arg) {
  if (!(arg instanceof judge_actions_judge_pb.AccessReply)) {
    throw new Error('Expected argument of type judge.AccessReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_AccessReply(buffer_arg) {
  return judge_actions_judge_pb.AccessReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_AccessRequest(arg) {
  if (!(arg instanceof judge_actions_judge_pb.AccessRequest)) {
    throw new Error('Expected argument of type judge.AccessRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_AccessRequest(buffer_arg) {
  return judge_actions_judge_pb.AccessRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_Resource(arg) {
  if (!(arg instanceof judge_models_types_pb.Resource)) {
    throw new Error('Expected argument of type judge.Resource');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_Resource(buffer_arg) {
  return judge_models_types_pb.Resource.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_ResourceMutationReply(arg) {
  if (!(arg instanceof judge_models_types_pb.ResourceMutationReply)) {
    throw new Error('Expected argument of type judge.ResourceMutationReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_ResourceMutationReply(buffer_arg) {
  return judge_models_types_pb.ResourceMutationReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_ResourceSelector(arg) {
  if (!(arg instanceof judge_models_types_pb.ResourceSelector)) {
    throw new Error('Expected argument of type judge.ResourceSelector');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_ResourceSelector(buffer_arg) {
  return judge_models_types_pb.ResourceSelector.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_Subject(arg) {
  if (!(arg instanceof judge_models_types_pb.Subject)) {
    throw new Error('Expected argument of type judge.Subject');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_Subject(buffer_arg) {
  return judge_models_types_pb.Subject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_SubjectMutationReply(arg) {
  if (!(arg instanceof judge_models_types_pb.SubjectMutationReply)) {
    throw new Error('Expected argument of type judge.SubjectMutationReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_SubjectMutationReply(buffer_arg) {
  return judge_models_types_pb.SubjectMutationReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_SubjectSelector(arg) {
  if (!(arg instanceof judge_models_types_pb.SubjectSelector)) {
    throw new Error('Expected argument of type judge.SubjectSelector');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_SubjectSelector(buffer_arg) {
  return judge_models_types_pb.SubjectSelector.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_UpdateResourceRequest(arg) {
  if (!(arg instanceof judge_models_types_pb.UpdateResourceRequest)) {
    throw new Error('Expected argument of type judge.UpdateResourceRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_UpdateResourceRequest(buffer_arg) {
  return judge_models_types_pb.UpdateResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_judge_UpdateSubjectRequest(arg) {
  if (!(arg instanceof judge_models_types_pb.UpdateSubjectRequest)) {
    throw new Error('Expected argument of type judge.UpdateSubjectRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_judge_UpdateSubjectRequest(buffer_arg) {
  return judge_models_types_pb.UpdateSubjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var JudgeService = exports.JudgeService = {
  // Subject CURD
  createSubject: {
    path: '/judge.Judge/createSubject',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.Subject,
    responseType: judge_models_types_pb.SubjectMutationReply,
    requestSerialize: serialize_judge_Subject,
    requestDeserialize: deserialize_judge_Subject,
    responseSerialize: serialize_judge_SubjectMutationReply,
    responseDeserialize: deserialize_judge_SubjectMutationReply,
  },
  deleteSubject: {
    path: '/judge.Judge/deleteSubject',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.SubjectSelector,
    responseType: judge_models_types_pb.SubjectMutationReply,
    requestSerialize: serialize_judge_SubjectSelector,
    requestDeserialize: deserialize_judge_SubjectSelector,
    responseSerialize: serialize_judge_SubjectMutationReply,
    responseDeserialize: deserialize_judge_SubjectMutationReply,
  },
  updateSubject: {
    path: '/judge.Judge/updateSubject',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.UpdateSubjectRequest,
    responseType: judge_models_types_pb.SubjectMutationReply,
    requestSerialize: serialize_judge_UpdateSubjectRequest,
    requestDeserialize: deserialize_judge_UpdateSubjectRequest,
    responseSerialize: serialize_judge_SubjectMutationReply,
    responseDeserialize: deserialize_judge_SubjectMutationReply,
  },
  getSubject: {
    path: '/judge.Judge/getSubject',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.SubjectSelector,
    responseType: judge_models_types_pb.Subject,
    requestSerialize: serialize_judge_SubjectSelector,
    requestDeserialize: deserialize_judge_SubjectSelector,
    responseSerialize: serialize_judge_Subject,
    responseDeserialize: deserialize_judge_Subject,
  },
  // Resource CURD
  getResource: {
    path: '/judge.Judge/getResource',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.ResourceSelector,
    responseType: judge_models_types_pb.Resource,
    requestSerialize: serialize_judge_ResourceSelector,
    requestDeserialize: deserialize_judge_ResourceSelector,
    responseSerialize: serialize_judge_Resource,
    responseDeserialize: deserialize_judge_Resource,
  },
  deleteResource: {
    path: '/judge.Judge/deleteResource',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.ResourceSelector,
    responseType: judge_models_types_pb.ResourceMutationReply,
    requestSerialize: serialize_judge_ResourceSelector,
    requestDeserialize: deserialize_judge_ResourceSelector,
    responseSerialize: serialize_judge_ResourceMutationReply,
    responseDeserialize: deserialize_judge_ResourceMutationReply,
  },
  updateResource: {
    path: '/judge.Judge/updateResource',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.UpdateResourceRequest,
    responseType: judge_models_types_pb.ResourceMutationReply,
    requestSerialize: serialize_judge_UpdateResourceRequest,
    requestDeserialize: deserialize_judge_UpdateResourceRequest,
    responseSerialize: serialize_judge_ResourceMutationReply,
    responseDeserialize: deserialize_judge_ResourceMutationReply,
  },
  createResource: {
    path: '/judge.Judge/createResource',
    requestStream: false,
    responseStream: false,
    requestType: judge_models_types_pb.Resource,
    responseType: judge_models_types_pb.ResourceMutationReply,
    requestSerialize: serialize_judge_Resource,
    requestDeserialize: deserialize_judge_Resource,
    responseSerialize: serialize_judge_ResourceMutationReply,
    responseDeserialize: deserialize_judge_ResourceMutationReply,
  },
  // Judge Service
  validatePolicy: {
    path: '/judge.Judge/validatePolicy',
    requestStream: false,
    responseStream: false,
    requestType: judge_actions_judge_pb.AccessRequest,
    responseType: judge_actions_judge_pb.AccessReply,
    requestSerialize: serialize_judge_AccessRequest,
    requestDeserialize: deserialize_judge_AccessRequest,
    responseSerialize: serialize_judge_AccessReply,
    responseDeserialize: deserialize_judge_AccessReply,
  },
};

exports.JudgeClient = grpc.makeGenericClientConstructor(JudgeService);
