package main

import (
	"ProtoPrac/src/protos"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)
var person =&protos.Person{Id:    1234, Name:  "John Doe",Email: "jdoe@example.com",Phones: []*protos.Person_PhoneNumber{{Number: "555-4321", Type: protos.Person_HOME},},}
var person1 =&protos.Person{Id:    1234, Name:  "John Doe",Email: "jdoe@example.com",Phones: []*protos.Person_PhoneNumber{{Number: "555-4321", Type: protos.Person_HOME},},}
var person2 =&protos.Person{Id:    1234, Name:  "John Doe",Email: "jdoe@example.com",Phones: []*protos.Person_PhoneNumber{{Number: "555-4321", Type: protos.Person_HOME},},}
var person3 =&protos.Person{Id:    1234, Name:  "John Doe",Email: "jdoe@example.com",Phones: []*protos.Person_PhoneNumber{{Number: "555-4321", Type: protos.Person_HOME},},}
var person4 =&protos.Person{Id:    1234, Name:  "John Doe",Email: "jdoe@example.com",Phones: []*protos.Person_PhoneNumber{{Number: "555-4321", Type: protos.Person_HOME},},}
var personCopy =&protos.PersonCopy{Id:    1234,Name:  "John Doe",Email: "jdoe@example.com",IdFirst:    1234,NameFirst:  "John Doe",EmailFirst: "jdoe@example.com",IdSecond:    1234,NameSecond:  "John Doe",EmailSecond: "jdoe@example.com",IdThird:    1234,NameThird:  "John Doe",EmailThird: "jdoe@example.com",Phones: []*protos.PersonCopy_PhoneNumber{{Number: "555-4321", Type: protos.PersonCopy_HOME},},}
var personCopy1 =&protos.PersonCopy{Id:    1234,Name:  "John Doe",Email: "jdoe@example.com",IdFirst:    1234,NameFirst:  "John Doe",EmailFirst: "jdoe@example.com",IdSecond:    1234,NameSecond:  "John Doe",EmailSecond: "jdoe@example.com",IdThird:    1234,NameThird:  "John Doe",EmailThird: "jdoe@example.com",Phones: []*protos.PersonCopy_PhoneNumber{{Number: "555-4321", Type: protos.PersonCopy_HOME},},}
var personCopy2 =&protos.PersonCopy{Id:    1234,Name:  "John Doe",Email: "jdoe@example.com",IdFirst:    1234,NameFirst:  "John Doe",EmailFirst: "jdoe@example.com",IdSecond:    1234,NameSecond:  "John Doe",EmailSecond: "jdoe@example.com",IdThird:    1234,NameThird:  "John Doe",EmailThird: "jdoe@example.com",Phones: []*protos.PersonCopy_PhoneNumber{{Number: "555-4321", Type: protos.PersonCopy_HOME},},}
var personCopy3 =&protos.PersonCopy{Id:    1234,Name:  "John Doe",Email: "jdoe@example.com",IdFirst:    1234,NameFirst:  "John Doe",EmailFirst: "jdoe@example.com",IdSecond:    1234,NameSecond:  "John Doe",EmailSecond: "jdoe@example.com",IdThird:    1234,NameThird:  "John Doe",EmailThird: "jdoe@example.com",Phones: []*protos.PersonCopy_PhoneNumber{{Number: "555-4321", Type: protos.PersonCopy_HOME},},}
var personCopy4 =&protos.PersonCopy{Id:    1234,Name:  "John Doe",Email: "jdoe@example.com",IdFirst:    1234,NameFirst:  "John Doe",EmailFirst: "jdoe@example.com",IdSecond:    1234,NameSecond:  "John Doe",EmailSecond: "jdoe@example.com",IdThird:    1234,NameThird:  "John Doe",EmailThird: "jdoe@example.com",Phones: []*protos.PersonCopy_PhoneNumber{{Number: "555-4321", Type: protos.PersonCopy_HOME},},}

func protoPracticeOneOf(){
	address:=&protos.AddressBook{Data:  []*protos.WidgetData{
		{
			WidgetData: &protos.WidgetData_PersonCopy{
				PersonCopy: personCopy,
			},
		},
		{
			WidgetData: &protos.WidgetData_Person{
				Person: person,
			},
		},{
			WidgetData: &protos.WidgetData_PersonCopy{
				PersonCopy: personCopy1,
			},
		},
		{
			WidgetData: &protos.WidgetData_Person{
				Person: person1,
			},
		},{
			WidgetData: &protos.WidgetData_PersonCopy{
				PersonCopy: personCopy2,
			},
		},
		{
			WidgetData: &protos.WidgetData_Person{
				Person: person2,
			},
		},{
			WidgetData: &protos.WidgetData_PersonCopy{
				PersonCopy: personCopy3,
			},
		},
		{
			WidgetData: &protos.WidgetData_Person{
				Person: person3,
			},
		},{
			WidgetData: &protos.WidgetData_PersonCopy{
				PersonCopy: personCopy4,
			},
		},
		{
			WidgetData: &protos.WidgetData_Person{
				Person: person4,
			},
		},
	},NextPageUrl: "google.com"}

	_, err := proto.Marshal(address)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	/*if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}*/
}

func protoPracticeAny(){
	serialized,_:= proto.Marshal(person);
	serializedCopy,_:= proto.Marshal(personCopy);
	serialized1,_:= proto.Marshal(person1);
	serializedCopy1,_:= proto.Marshal(personCopy1);
	serialized2,_:= proto.Marshal(person2);
	serializedCopy2,_:= proto.Marshal(personCopy2);
	serialized3,_:= proto.Marshal(person3);
	serializedCopy3,_:= proto.Marshal(personCopy3);
	serialized4,_:= proto.Marshal(person4);
	serializedCopy4,_:= proto.Marshal(personCopy4);

	address := &protos.AddressBookAny{
		Widgets_: []*anypb.Any{
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(person),
				Value:   serialized,
			},
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(personCopy),
				Value:   serializedCopy,
			},			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(person),
				Value:   serialized1,
			},
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(personCopy),
				Value:   serializedCopy1,
			},			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(person),
				Value:   serialized2,
			},
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(personCopy),
				Value:   serializedCopy2,
			},			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(person),
				Value:   serialized3,
			},
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(personCopy),
				Value:   serializedCopy3,
			},			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(person),
				Value:   serialized4,
			},
			{
				TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(personCopy),
				Value:   serializedCopy4,
			},
		},
	}
	_, err := proto.Marshal(address)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	/*if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}*/
}

