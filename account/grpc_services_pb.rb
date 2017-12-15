# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: account/grpc.proto for package 'account'

require 'grpc'
require 'account/grpc_pb'

module Account
  module Account
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'account.Account'

      rpc :validateUserPassword, ValidateUserRequest, ValidateUserReply
      rpc :getUserInfo, GetUserInfoRequest, GetUserInfoReply
      rpc :getTokenInfo, GetTokenInfoRequest, GetTokenInfoReply
    end

    Stub = Service.rpc_stub_class
  end
end
