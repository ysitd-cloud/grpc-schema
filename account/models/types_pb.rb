# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/models/types.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "account.User" do
    optional :username, :string, 1
    optional :display_name, :string, 2
    optional :email, :string, 3
    optional :avatar_url, :string, 4
  end
  add_message "account.Service" do
    optional :id, :string, 1
  end
  add_message "account.Token" do
    optional :issuer, :message, 1, "account.User"
    optional :audience, :message, 2, "account.Service"
    optional :expire, :message, 3, "google.protobuf.Timestamp"
    repeated :scopes, :string, 4
  end
end

module Account
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("account.User").msgclass
  Service = Google::Protobuf::DescriptorPool.generated_pool.lookup("account.Service").msgclass
  Token = Google::Protobuf::DescriptorPool.generated_pool.lookup("account.Token").msgclass
end
