module gitee.com/favefan/doconsole

go 1.16

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/beego/beego/v2 v2.0.1
	github.com/containerd/containerd v1.4.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.5+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20210310155132-4ce2db91004e // indirect
	google.golang.org/grpc v1.36.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.21.2
	gotest.tools/v3 v3.0.3 // indirect
)

replace (
	gitee.com/favefan/doconsole/configs => ../doconsole/conf
	gitee.com/favefan/doconsole/middlewares => ../doconsole/middleware
	gitee.com/favefan/doconsole/middlewares/jwt => ../doconsole/middleware/jwt
	gitee.com/favefan/doconsole/models => ../doconsole/models
	gitee.com/favefan/doconsole/pkg/setting => ../doconsole/pkg/setting
	gitee.com/favefan/doconsole/pkg/util => ../doconsole/pkg/util
	gitee.com/favefan/doconsole/routers => ../doconsole/routers
	gitee.com/favefan/doconsole/routers/api => ../doconsole/routers/api
	gitee.com/favefan/doconsole/services => ../doconsole/services
)
