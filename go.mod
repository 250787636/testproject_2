module ky/ssp

go 1.14

replace (
	kmesh2/ass => ../kmesh2/ass
	kmesh2/common => ../kmesh2/common
	kmesh2/core => ../kmesh2/core
	//kmesh2/dps => ../ssp_kmesh2/dps
	kmesh2/haproxy => ../kmesh2/haproxy
	kmesh2/utils => ../kmesh2/utils

//kmesh2/ass => ../myKmesh/ass
//kmesh2/common => ../myKmesh/common
//kmesh2/core => ../myKmesh/core
////kmesh2/dps => ../ssp_kmesh2/dps
//kmesh2/haproxy => ../myKmesh/haproxy
//kmesh2/utils => ../myKmesh/utils
)

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/dchest/captcha v0.0.0-20200903113550-03f5f0333e1f
	github.com/fsnotify/fsnotify v1.5.4
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gocarina/structs v0.0.0-20140918155756-eba5a0f1cc3d // indirect
	github.com/google/uuid v1.3.0
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jlaffaye/ftp v0.0.0-20220630165035-11536801d1ff
	github.com/json-iterator/go v1.1.10
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pikanezi/mapslice v0.0.0-20160614093333-dd6dd9af94b7
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tidwall/gjson v1.6.7
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/typa01/go-utils v0.0.0-20181126045345-a86b05b01c1e
	github.com/wumansgy/goEncrypt v0.0.0-20201114063050-efa0a0601707
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/text v0.3.5
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/olivere/elastic.v5 v5.0.86
	gorm.io/gorm v1.23.8
	kmesh2/utils v0.0.0
)
