package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"xxx/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	//查询数据库,查找到所有的community并返回
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows { //查询结果为空
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//GetCommunityDetailByID 根据id查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time 
              from community 
              where community_id=?`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}