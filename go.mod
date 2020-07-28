module github.com/floj/jqlay

go 1.14

replace github.com/jingweno/jqplay => ./

require (
	github.com/boltdb/bolt v1.3.1
	github.com/jingweno/jqplay v0.0.0-00010101000000-000000000000
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/unrolled/secure v1.0.8
	gopkg.in/gin-gonic/gin.v1 v1.1.5-0.20170702092826-d459835d2b07
)
