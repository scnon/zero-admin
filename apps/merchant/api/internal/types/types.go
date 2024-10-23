// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type AddStoreReq struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Status    int    `json:"status"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Address   string `json:"address"`
}

type Base struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type DeleteStoreReq struct {
	Ids []int64
}

type DeptCreateReq struct {
	ParentID uint64 `json:"parent_id"`
	Name     string `json:"name"`
	Status   int32  `json:"status"`
	Sort     int32  `json:"sort"`
}

type DeptCreateResp struct {
	Base
	Data uint64 `json:"data"`
}

type DeptData struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	ParentID   uint64 `json:"parent_id"`
	Sort       int32  `json:"sort"`
	Status     int32  `json:"status"`
	Creator    string `json:"creator"`
	Updater    string `json:"updater"`
	CreateTime uint64 `json:"create_time"`
	UpdateTime uint64 `json:"update_time"`
}

type DeptDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type DeptDeleteResp struct {
	Base
}

type DeptInfo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type DeptListReq struct {
}

type DeptListResp struct {
	Base
	Data []*DeptData `json:"data"`
}

type DeptUpdateReq struct {
	ID       uint64 `json:"id"`
	ParentID uint64 `json:"parent_id"`
	Name     string `json:"name"`
	Sort     int32  `json:"sort"`
	Status   int32  `json:"status"`
}

type DeptUpdateResp struct {
	Base
}

type LoginData struct {
	UserId       uint64   `json:"userId"`
	Username     string   `json:"username"`
	NickName     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	ExpireTime   uint64   `json:"expireTime"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Base
	Data LoginData `json:"data,omitempty"`
}

type MenuCreateReq struct {
	ParentID uint64 `json:"parent_id"`
	Title    string `json:"title"`
	Path     string `json:"path"`
	Sort     int32  `json:"sort"`
	Name     string `json:"name"`
}

type MenuCreateResp struct {
	Base
}

type MenuData struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	ParentID   uint64 `json:"parent_id"`
	Title      string `json:"title"`
	Path       string `json:"path"`
	Sort       int32  `json:"sort"`
	Creator    string `json:"creator"`
	Updater    string `json:"updater"`
	CreateTime uint64 `json:"create_time"`
	UpdateTime uint64 `json:"update_time"`
}

type MenuDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type MenuDeleteResp struct {
	Base
}

type MenuInfoResp struct {
	Base
	Data []*UserMenuInfo `json:"data"`
}

type MenuListReq struct {
}

type MenuListResp struct {
	Base
	Data []MenuData `json:"data"`
}

type MenuMetaInfo struct {
	Title string `json:"title"`
}

type MenuUpdateReq struct {
	ID       uint64 `json:"id"`
	ParentID uint64 `json:"parent_id"`
	Title    string `json:"title"`
	Path     string `json:"path"`
	Sort     int32  `json:"sort"`
	Name     string `json:"name"`
}

type MenuUpdateResp struct {
	Base
}

type PageData struct {
	Page     uint32 `json:"page,default=1"`
	PageSize uint32 `json:"page_size,default=10"`
}

type RoleCreateReq struct {
	Name   string `json:"name"`
	Sort   int32  `json:"sort"`
	Status int32  `json:"status"`
	Remark string `json:"remark"`
}

type RoleCreateResp struct {
	Base
	Data uint64 `json:"data"`
}

type RoleData struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
	Sort       int32  `json:"sort"`
	Status     int32  `json:"status"`
	Cretor     string `json:"creator"`
	Updater    string `json:"updator"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type RoleDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type RoleDeleteResp struct {
	Base
}

type RoleList struct {
	Total uint64      `json:"total"`
	List  []*RoleData `json:"list"`
}

type RoleListReq struct {
	PageData
	Name   string `json:"name,optional"`
	Status int32  `json:"status,default=1"`
}

type RoleListResp struct {
	Base
	Data RoleList `json:"data"`
}

type RoleMenuIdsReq struct {
	RoleId uint64 `json:"role_id"`
}

type RoleMenuIdsResp struct {
	Base
	Data []uint64 `json:"data"`
}

type RoleUpdateReq struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   int32  `json:"sort"`
	Status int32  `json:"status"`
}

type RoleUpdateResp struct {
	Base
}

type StoreInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Status    int    `json:"status"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Address   string `json:"address"`
}

type StoreListReq struct {
	Page     int32 `json:"page,default=1"`
	PageSize int32 `json:"page_size,default=10"`
}

type StoreListResp struct {
	Total int64       `json:"total"`
	List  []StoreInfo `json:"list"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Status   int32  `json:"status,default=1"`
	Sort     int32  `json:"sort,default=0"`
	Remark   string `json:"remark,default=''"`
}

type UserCreateResp struct {
	Base
}

type UserData struct {
	Id         uint64   `json:"id"`
	Username   string   `json:"username"`
	NickName   string   `json:"nickname"`
	Avatar     string   `json:"avatar"`
	Status     int32    `json:"status"`
	Sort       int32    `json:"sort"`
	Remark     string   `json:"remark"`
	Creator    string   `json:"creator"`
	Updator    string   `json:"updator"`
	Dept       DeptInfo `json:"dept"`
	CreateTime uint64   `json:"createTime"`
	UpdateTime uint64   `json:"updateTime"`
}

type UserDeleteReq struct {
	Ids []uint64 `json:"ids"`
}

type UserDeleteResp struct {
	Base
}

type UserListData struct {
	Total uint64     `json:"total"`
	List  []UserData `json:"list"`
}

type UserListReq struct {
	PageData
	Username *string `json:"username,optional"`
	NickName *string `json:"nickname,optional"`
	Status   *int32  `json:"status,default=1"`
}

type UserListResp struct {
	Base
	Data UserListData `json:"data,omitempty"`
}

type UserMenuInfo struct {
	ID       uint64         `json:"id"`
	ParentId uint64         `json:"parent_id"`
	Path     string         `json:"path"`
	Sort     int32          `json:"sort"`
	Name     string         `json:"name"`
	Meta     MenuMetaInfo   `json:"meta"`
	Children []UserMenuInfo `json:"children,omitempty"`
}

type UserRestPwdReq struct {
	UserId   uint64 `json:"userId"`
	Password string `json:"password"`
}

type UserRestPwdResp struct {
	Base
}

type UserRoleIdsReq struct {
	UserId uint64 `json:"userId"`
}

type UserRoleIdsResp struct {
	Base
	Data []uint64 `json:"data,omitempty"`
}

type UserUpdateReq struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Status   int32  `json:"status"`
	Sort     int32  `json:"sort"`
	Remark   string `json:"remark"`
}

type UserUpdateResp struct {
	Base
}
