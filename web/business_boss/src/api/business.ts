import webapi from "./gocliRequest";
import type * as components from "./businessComponents";
export * from "./businessComponents";

/**
 * @description
 * @param req
 */
export function addStore(req: components.AddStoreReq) {
  return webapi.post<components.StoreInfo>(`/api/v1/business/store/add`, req);
}

/**
 * @description
 * @param params
 */
export function deleteStore(params: components.DeleteStoreReqParams) {
  return webapi.post<null>(`/api/v1/business/store/delete`, params);
}

/**
 * @description
 * @param req
 */
export function storeList(req: components.StoreListReq) {
  return webapi.post<components.StoreListResp>(
    `/api/v1/business/store/list`,
    req
  );
}

/**
 * @description
 * @param req
 */
export function updateStore(req: components.StoreInfo) {
  return webapi.post<components.StoreInfo>(
    `/api/v1/business/store/update`,
    req
  );
}

/**
 * @description
 * @param req
 */
export function addMenu(req: components.AddMenuReq) {
  return webapi.post<components.AddMenuResp>(`/api/v1/system/menu/add`, req);
}

/**
 * @description
 * @param req
 */
export function deleteMenu(req: components.DeleteMenuReq) {
  return webapi.post<null>(`/api/v1/system/menu/delete`, req);
}

/**
 * @description
 */
export function menuInfo() {
  return webapi.post<Array<components.UserMenuInfo>>(
    `/api/v1/system/menu/info`,
    {}
  );
}

/**
 * @description
 * @param req
 */
export function menuList(req: components.MenuListReq) {
  return webapi.post<components.MenuListResp>(`/api/v1/system/menu/list`, req);
}

/**
 * @description
 * @param req
 */
export function updateMenu(req: components.MenuInfo) {
  return webapi.post<null>(`/api/v1/system/menu/update`, req);
}

/**
 * @description
 * @param req
 */
export function addUser(req: components.AddUserReq) {
  return webapi.post<components.UserData>(`/api/v1/system/user/add`, req);
}

/**
 * @description
 * @param req
 */
export function deleteUser(req: components.DeleteUserReq) {
  return webapi.post<null>(`/api/v1/system/user/delete`, req);
}

/**
 * @description
 * @param req
 */
export function queryUserList(req: components.QueryUserListReq) {
  return webapi.post<components.QueryUserListResp>(
    `/api/v1/system/user/list`,
    req
  );
}

/**
 * @description
 * @param req
 */
export function updateUser(req: components.UpdateUserReq) {
  return webapi.post<null>(`/api/v1/system/user/update`, req);
}

/**
 * @description
 * @param req
 */
export function userLogin(req: components.LoginReq) {
  return webapi.post<components.LoginResp>(`/api/v1/system/user/login`, req);
}
