module github.com/ilovelili/dongfeng-core-proxy

go 1.12

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emicklei/go-restful v2.9.3+incompatible
	github.com/go-log/log v0.1.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/gocarina/gocsv v0.0.0-20190313153828-c075544dca88
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.1.1
	github.com/hashicorp/consul v1.4.4
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/hashicorp/go-rootcerts v1.0.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/hashicorp/serf v0.8.2
	github.com/ilovelili/aliyun-client v0.0.0-20190407060233-e71f697c8303
	github.com/ilovelili/dongfeng-error-code v0.0.0-20190404061658-b59a7f3fe1a3
	github.com/ilovelili/dongfeng-logger v0.0.0-20190403091018-f20598e7c461
	github.com/ilovelili/dongfeng-protobuf v0.0.0-20190404052200-ec920597149a
	github.com/ilovelili/dongfeng-shared-lib v0.0.0-20190108085915-4093ff764c36
	github.com/json-iterator/go v1.1.6
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/lestrrat-go/jwx v0.0.0-20190405030015-23dc777739fc
	github.com/lestrrat-go/pdebug v0.0.0-20180220043849-39f9a71bcabe
	github.com/micro/cli v0.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.0.0
	github.com/micro/go-plugins v0.16.2
	github.com/micro/go-rcache v0.3.0
	github.com/micro/go-web v0.4.0
	github.com/micro/h2c v1.0.0
	github.com/micro/mdns v0.1.0
	github.com/micro/util v0.2.0
	github.com/miekg/dns v1.1.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/pkg/errors v0.8.1
	github.com/segmentio/ksuid v1.0.2
	github.com/sirupsen/logrus v1.4.1
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/sys v0.0.0-20190405154228-4b34438f7a67
	golang.org/x/text v0.3.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
)

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.1

replace github.com/ilovelili/dongfeng-protobuf => github.com/ilovelili/dongfeng-protobuf v0.0.0-20190618065646-870ed0f2e9aa

replace github.com/ilovelili/dongfeng-error-code => github.com/ilovelili/dongfeng-error-code v0.0.0-20190618065903-9bcc1dd6022c

replace github.com/ilovelili/aliyun-client => github.com/ilovelili/aliyun-client v0.0.0-20190605074008-4fbaa377d984

replace github.com/micro/go-micro v1.0.0 => github.com/micro/go-micro v0.14.1
