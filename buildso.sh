go build  -gcflags="all=-N -l" -buildmode=plugin -o ./pluginB.so ./pluginB/plugin.go
go build    -buildmode=plugin -o ./bin/pluginB.so ./pluginB/plugin.go
go build    -o ./bin/gowor