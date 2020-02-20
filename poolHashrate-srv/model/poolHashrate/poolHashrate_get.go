package poolHashrate

import (
	"github.com/jiangjinyuan/go-micro-demo/basic/db"
	proto "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
	"github.com/micro/go-micro/util/log"
	"strconv"
)

func (s *service) GetPoolHashrate(poolID string) (ret *proto.PoolHashrate, err error) {
	queryString := `SELECT pool_name, pool_id, hashrate FROM hashrate WHERE pool_id = ? limit 1`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.PoolHashrate{}

	poolId,_:=strconv.Atoi(poolID)
	// 查询
	err = o.QueryRow(queryString, poolId).Scan(&ret.PoolName, &ret.PoolId, &ret.Hashrate)
	if err != nil {
		log.Logf("[GetPoolHashrate] 查询数据失败，err：%s", err)
		return
	}
	return
}
