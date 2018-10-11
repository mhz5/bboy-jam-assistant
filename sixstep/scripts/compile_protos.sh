mkdir $BBOY_APP_ROOT/sixstep/protos
mkdir $BBOY_APP_ROOT/sixstep/protos/sixstep

protoc --go_out=$BBOY_APP_ROOT/sixstep/protos/sixstep --proto_path=$BBOY_APP_ROOT/protos $BBOY_APP_ROOT/protos/*.proto

