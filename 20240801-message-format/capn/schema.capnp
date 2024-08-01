using Go = import "/go.capnp";
@0xfd74a990b04176d4;
$Go.package("capn");
$Go.import("capn/schema");

struct Employee {
    name @0 :Text;
    position @1 :Text;
}

struct Result {
    employee @0 :Employee;
    scores @1 :List(UInt64);
}

struct Message {
    message @0 :Text;
    result @1 :List(Result);
}
