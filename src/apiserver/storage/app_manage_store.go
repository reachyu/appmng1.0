package storage

import (
	"appmng/src/common"
	"fmt"
	"encoding/json"
)

/*
func ListAppManages(opts *list.Options) ([]*model.AppManage, int, string, error) {
	errorF := func(err error) ([]*model.AppManage, int, string, error) {
		return nil, 0, "", util.NewInternalServerError(err, "Failed to list app_manages: %v", err)
	}

	buildQuery := func(sqlBuilder sq.SelectBuilder) sq.SelectBuilder {
		return sqlBuilder.From("app_manages") // .Where(sq.Eq{"dr": model.AppReady})
	}

	sqlBuilder := buildQuery(sq.Select("*"))

	// SQL for row list
	rowsSql, rowsArgs, err := opts.AddPaginationToSelect(sqlBuilder).ToSql()
	if err != nil {
		return errorF(err)
	}

	// SQL for getting total size. This matches the query to get all the rows above, in order
	// to do the same filter, but counts instead of scanning the rows.
	sizeSql, sizeArgs, err := buildQuery(sq.Select("count(*)")).ToSql()
	if err != nil {
		return errorF(err)
	}

	// Use a transaction to make sure we're returning the total_size of the same rows queried
	tx, err := s.db.Begin()
	if err != nil {
		glog.Errorf("Failed to start transaction to list app_manages")
		return errorF(err)
	}

	rows, err := tx.Query(rowsSql, rowsArgs...)
	if err != nil {
		tx.Rollback()
		return errorF(err)
	}
	apps, err := s.scanRows(rows)
	if err != nil {
		tx.Rollback()
		return errorF(err)
	}
	rows.Close()

	sizeRow, err := tx.Query(sizeSql, sizeArgs...)
	if err != nil {
		tx.Rollback()
		return errorF(err)
	}
	total_size, err := list.ScanRowToTotalSize(sizeRow)
	if err != nil {
		tx.Rollback()
		return errorF(err)
	}
	sizeRow.Close()

	err = tx.Commit()
	if err != nil {
		glog.Errorf("Failed to commit transaction to list app_manages")
		return errorF(err)
	}

	if len(apps) <= opts.PageSize {
		return apps, total_size, "", nil
	}

	npt, err := opts.NextPageToken(apps[opts.PageSize])
	return apps[:opts.PageSize], total_size, npt, err
}

func scanRows(rows *sql.Rows) ([]*model.AppManage, error) {
	var apps []*model.AppManage
	for rows.Next() {
		var ApplyName, ApplyType, ApplyFrame, ApplyEnvironment string
		var ApplyBrief, CreatedBy, UpdatedBy sql.NullString
		var id int64
		var CreatedAt, UpdateAt sql.NullInt64
		//var status model.PipelineStatus
		if err := rows.Scan(&id, &ApplyName, &ApplyType, &ApplyFrame, &ApplyEnvironment, &ApplyBrief, &CreatedAt, &UpdateAt, &CreatedBy, &UpdatedBy); err != nil {
			return nil, err
		}
		apps = append(apps, &model.AppManage{
			ID:               id,
			ApplyName:        ApplyName,
			ApplyType:        ApplyType,
			ApplyFrame:       ApplyFrame,
			ApplyEnvironment: ApplyEnvironment,
			ApplyBrief:       ApplyBrief,
			CreatedAt:        CreatedAt,
			UpdateAt:         UpdateAt,
			CreatedBy:        CreatedBy,
			UpdatedBy:        UpdatedBy})
	}
	return apps, nil
}

func GetAppManage(id string) (*model.AppManage, error) {
	sql, args, err := sq.
		Select("*").
		From("app_manages").
		Where(sq.Eq{"id": id}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get app_manages: %v", err.Error())
	}
	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get app_manages: %v", err.Error())
	}
	defer r.Close()
	apps, err := s.scanRows(r)

	if err != nil || len(apps) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get app_manages: %v", err.Error())
	}
	if len(apps) == 0 {
		return nil, util.NewResourceNotFoundError("AppManage", fmt.Sprint(id))
	}
	return apps[0], nil
}

func DeleteAppManage(id string) error {
	sql, args, err := sq.Delete("app_manages").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete app_manages: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete app_manages: %v", err.Error())
	}
	return nil
}

func CreateAppManage(p *model.AppManage) (*model.AppManage, error) {
	newAppManage := *p
	now := s.time.Now().Unix()
	sql, args, err := sq.
		Insert("app_manages").
		SetMap(
			sq.Eq{
				"apply_name":        newAppManage.ApplyName,
				"apply_type":        newAppManage.ApplyType,
				"apply_frame":       newAppManage.ApplyFrame,
				"apply_environment": newAppManage.ApplyEnvironment,
				"apply_brief":       newAppManage.ApplyBrief,
				"created_at":        now}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert app to app_manages table: %v",
			err.Error())
	}
	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewInvalidInputError(
				"Failed to create a new app. The name %v already exist. Please specify a new name.", p.ApplyName)
		}
		return nil, util.NewInternalServerError(err, "Failed to add app to app_manages table: %v",
			err.Error())
	}
	return &newAppManage, nil
}

func UpdateAppManageStatus(id string, status model.AppManageStatus) error {
	sql, args, err := sq.
		Update("app_manages").
		SetMap(sq.Eq{"dr": status}).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to update the app_manages metadata: %s", err.Error())
	}
	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to update the app_manages metadata: %s", err.Error())
	}
	return nil
}
*/

func GetUserinfo() {
	db := common.GetInstance().GetMysqlDB()
	rows, err := db.Raw("select * from user").Rows()
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}

	for rows.Next() {
		//定义变量接收查询数据
		var id, name, age string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("no data")
		}
		fmt.Println("查询结果======%s %s %s \n", id, name, age)
	}

	//关闭结果集（释放连接）
	if rows != nil {
		rows.Close()
	}
}

func GetAppinfo() string {
	db := common.GetInstance().GetMysqlDB()
	rows, err := db.Raw("select id,apply_name,apply_type,apply_frame,apply_environment from app_manages").Rows()
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return ""
	}
	var rlt map[string]string
	rlt = make(map[string]string)
	for rows.Next() {
		//定义变量接收查询数据
		var ApplyName, ApplyType, ApplyFrame, ApplyEnvironment string
		var id int64
		err := rows.Scan(&id, &ApplyName, &ApplyType, &ApplyFrame, &ApplyEnvironment)
		if err != nil {
			fmt.Println("no data,error==== %s",err)
		}
		fmt.Println("查询结果======%s %s %s %s %s \n", id, ApplyName, ApplyType, ApplyFrame, ApplyEnvironment)
		rlt["id"] = string(id)
		rlt["ApplyName"] = ApplyName
		rlt["ApplyType"] = ApplyType
		rlt["ApplyFrame"] = ApplyFrame
		rlt["ApplyEnvironment"] = ApplyEnvironment
	}

	payloadRlt, _ := json.Marshal(rlt)

	//关闭结果集（释放连接）
	if rows != nil {
		rows.Close()
	}

	return string(payloadRlt)

}
