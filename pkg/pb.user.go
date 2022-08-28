package pkg

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type GetUserResponse struct {
	state protoimpl.MessageState
	sizeCache protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Users []*User `json:"users,omitempty" protobuf:"bytes,1,rep,name=users,proto3"`
}

type User struct {
	state protoimpl.MessageState
	sizeCache protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `json:"id,omitempty" protobuf:"varint,1,opt,name=id,proto3"`
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name,proto3"`
	Email string `json:"email,omitempty" protobuf:"bytes,3,opt,name=email,proto3"`
}

func (*GetUserResponse) Descriptor() ([]byte,[]int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

func (g *GetUserResponse) GetUsers() []*User {
	if g != nil {
		return g.Users
	}
	return nil
}

func (u *User) Reset() {
	*u = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(u))
		ms.StoreMessageInfo(mi)
	}
}

func (u *User) String() string {
	return protoimpl.X.MessageStringOf(u)
}

func (*User) ProtoMessage {}

func (u *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && u != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(u))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.Messageof(u)
}

func (*CreateUserResponse) Descriptor() ([]byte,[]int) {
	return file_api_user_proto_rawDescGZIP(), []int{3}
}

func (c *Cre)