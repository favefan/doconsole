module gitee.com/favefan/doconsole

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/beego/beego/v2 v2.0.1
	github.com/containerd/containerd v1.4.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.5+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/kevinburke/ssh_config v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20210310155132-4ce2db91004e // indirect
	google.golang.org/grpc v1.36.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.21.2
	gotest.tools/v3 v3.0.3 // indirect
)

replace (
	gitee.com/favefan/doconsole/configs => ../doconsole/conf
	gitee.com/favefan/doconsole/global => ../doconsole/global
	gitee.com/favefan/doconsole/initialize => ../doconsole/initialize
	gitee.com/favefan/doconsole/middlewares => ../doconsole/middleware
	gitee.com/favefan/doconsole/middlewares/jwt => ../doconsole/middleware/jwt
	gitee.com/favefan/doconsole/models => ../doconsole/models
	// gitee.com/favefan/doconsole/pkg/setting => ../doconsole/pkg/setting
	gitee.com/favefan/doconsole/pkg/util => ../doconsole/pkg/util
	gitee.com/favefan/doconsole/pkg/e => ../doconsole/pkg/e
	gitee.com/favefan/doconsole/pkg/app => ../doconsole/pkg/app
	gitee.com/favefan/doconsole/routers => ../doconsole/routers
	gitee.com/favefan/doconsole/routers/api => ../doconsole/routers/api
	gitee.com/favefan/doconsole/service => ../doconsole/service
	gitee.com/favefan/doconsole/service/websocket => ../doconsole/service/websocket
)
