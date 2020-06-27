package common

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine
var PubSub *gochannel.GoChannel
