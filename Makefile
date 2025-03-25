kitex_gen_user:
	kitex --thrift-plugin validator -module github.com/zhiqinkuang/easy-note idl/user.thrift # execute in the project root directory

kitex_gen_note:
	kitex --thrift-plugin validator -module github.com/zhiqinkuang/easy-note idl/note.thrift # execute in the project root directory

install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest


hertz_gen_model:
	hz model --idl=idl/api.thrift --mod=github.com/zhiqinkuang/easy-note --model_dir=hertz_gen

hertz_gen_client:
	hz client --idl=idl/api.thrift --base_domain=http://127.0.0.1:8080 --client_dir=api_request --mod=github.com/zhiqinkuang/easy-note --model_dir=hertz_gen