package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"xs.bbs/pkg/conf"
)

var node *sf.Node

func Init(c *conf.Config) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", c.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(c.MachineID)
	if err != nil {
		return
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
