# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account/models/types.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "account.User" do
    optional :username, :string, 1
    optional :display_name, :string, 2
    optional :email, :string, 3
    optional :avatar_url, :string, 4
  end
end

module Account
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("account.User").msgclass
end
