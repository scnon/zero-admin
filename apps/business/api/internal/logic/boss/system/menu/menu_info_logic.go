package menu

import (
	"context"
	"xlife/apps/auth/rpc/auth"

	"xlife/apps/business/api/internal/svc"
	"xlife/apps/business/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuInfoLogic {
	return &MenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuInfoLogic) MenuInfo() (resp []*types.UserMenuInfo, err error) {
	// uid := ctxdata.GetUId(l.ctx)
	menuList, err := l.svcCtx.Menu.GetMenu(l.ctx, &auth.GetMenuReq{
		TenantId: l.svcCtx.Config.Tenant,
	})
	if err != nil {
		return nil, err
	}

	resp = buildTree(menuList.Menu)
	return resp, nil
}

func buildTree(list []*auth.MenuData) []*types.UserMenuInfo {
	// 创建一个映射，将每个 ListItem 转换为 Node
	nodeMap := make(map[uint64][]*types.UserMenuInfo)
	allNodes := make(map[uint64]*types.UserMenuInfo)

	// 遍历所有 ListItem，将其转换为 Node
	for _, item := range list {
		node := &types.UserMenuInfo{
			ID: item.Id,
			Meta: types.MenuMetaInfo{
				Title: item.Title,
			},
			Name:     item.Name,
			Path:     item.Path,
			ParentId: item.ParentId,
		}
		// 将节点保存到 allNodes
		allNodes[node.ID] = node
		nodeMap[node.ParentId] = append(nodeMap[node.ParentId], node)
	}

	// 递归构建树
	var build func(parentID uint64) []*types.UserMenuInfo
	build = func(parentID uint64) []*types.UserMenuInfo {
		// 获取 parentID 对应的所有子节点
		children := nodeMap[parentID]
		// 遍历每个子节点，递归构建它的子树
		for _, child := range children {
			childrenList := build(child.ID)
			for _, children := range childrenList {
				child.Children = append(child.Children, *children)
			}
		}
		return children
	}

	// 返回根节点的子树（假设 parent_id 为 0 表示根节点）
	return build(0)
}
