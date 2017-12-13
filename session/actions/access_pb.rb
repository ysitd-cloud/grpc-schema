# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: session/actions/access.proto

require 'google/protobuf'

require 'session/models/types_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "session.CreateSessoinRequest" do
    optional :service, :message, 1, "session.Service"
    optional :session, :message, 2, "session.SessionContent"
  end
  add_message "session.CreateSessionReply" do
    optional :id, :string, 1
  end
  add_message "session.GetSessionRequest" do
    optional :service, :message, 1, "session.Service"
    optional :id, :string, 2
  end
  add_message "session.GetSessionReply" do
    optional :exists, :bool, 1
    optional :session, :message, 2, "session.SessionContent"
  end
  add_message "session.UpdateSessionRequest" do
    optional :service, :message, 1, "session.Service"
    optional :id, :string, 2
    optional :content, :message, 3, "session.SessionContent"
  end
  add_message "session.UpdateSessionReply" do
    optional :success, :bool, 1
  end
  add_message "session.DeleteSessionRequest" do
    optional :service, :message, 1, "session.Service"
    optional :id, :string, 2
  end
  add_message "session.DeleteSessionReply" do
    optional :success, :bool, 1
  end
end

module Session
  CreateSessoinRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.CreateSessoinRequest").msgclass
  CreateSessionReply = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.CreateSessionReply").msgclass
  GetSessionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.GetSessionRequest").msgclass
  GetSessionReply = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.GetSessionReply").msgclass
  UpdateSessionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.UpdateSessionRequest").msgclass
  UpdateSessionReply = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.UpdateSessionReply").msgclass
  DeleteSessionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.DeleteSessionRequest").msgclass
  DeleteSessionReply = Google::Protobuf::DescriptorPool.generated_pool.lookup("session.DeleteSessionReply").msgclass
end