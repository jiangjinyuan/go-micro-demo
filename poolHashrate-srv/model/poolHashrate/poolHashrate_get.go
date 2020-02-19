package poolHashrate

import (
	"github.com/jiangjinyuan/go-micro-demo/basic/db"
	"github.com/micro/go-micro/util/log"
	proto "github.com/jiangjinyuan/go-micro-demo/poolHashrate-srv/proto/poolHashrate"
)
func (s *service) GetPoolHashrate(poolName string,poolID int32) (ret *proto.PoolHashrate, err error) {
	queryString := `SELECT pool_name, pool_id, hashrate FROM hashrate WHERE pool_name = ? and pool_id = ? limit 1`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.PoolHashrate{}

	// 查询
	err = o.QueryRow(queryString, poolName,poolID).Scan(&ret.PoolName, &ret.PoolId, &ret.Hashrate)
	if err != nil {
		log.Logf("[GetPoolHashrate] 查询数据失败，err：%s", err)
		return
	}
	return
}
