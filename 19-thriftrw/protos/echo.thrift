namespace go echo

struct EchoReq {
    1: required string msg;
}

struct EchoRes {
    1: optional string msg;
}

service Echo {
    EchoRes echo(1: EchoReq req);
}