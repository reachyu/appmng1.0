package resource

import (
	"appmng/src/apiserver/storage"
)

/*
func (r *ResourceManager) ListAppManages(opts *list.Options) (
	apps []*model.AppManage, total_size int, nextPageToken string, err error) {
	return r.appManageStore.ListAppManages(opts)
}

func (r *ResourceManager) GetAppManage(id string) (*model.AppManage, error) {
	return r.appManageStore.GetAppManage(id)
}

func (r *ResourceManager) DeleteAppManage(id string) error {
	return r.appManageStore.DeleteAppManage(id)
}

func (r *ResourceManager) CreateAppManage(appManage *model.AppManage) (*model.AppManage, error) {
	return r.appManageStore.CreateAppManage(appManage)
}

func (r *ResourceManager) UpdateAppManageStatus(id string,status model.AppManageStatus) error {
	return r.appManageStore.UpdateAppManageStatus(id,status)
}
*/
func GetUserinfo()  {
	storage.GetUserinfo()
}

func GetAppinfo()  {
	storage.GetAppinfo()
}