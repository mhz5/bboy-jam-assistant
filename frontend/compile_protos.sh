mkdir $BBOY_APP_ROOT/frontend/protos
mkdir $BBOY_APP_ROOT/frontend/protos/sixstep

protoc --js_out=$BBOY_APP_ROOT/frontend/protos/sixstep --proto_path=$BBOY_APP_ROOT/protos $BBOY_APP_ROOT/protos/*.proto
