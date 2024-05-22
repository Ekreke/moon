package repoimpl

import (
	"context"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/query"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/repo"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/data"
	"github.com/aide-cloud/moon/pkg/types"
	"github.com/aide-cloud/moon/pkg/vobj"
	"gorm.io/gen"
)

func NewUserRepo(data *data.Data) repo.UserRepo {
	return &userRepoImpl{
		data: data,
	}
}

type userRepoImpl struct {
	data *data.Data
}

func (l *userRepoImpl) UpdateByID(ctx context.Context, user *bo.UpdateUserParams) error {
	userModel := updateUserParamsToModel(user)
	return userModel.UpdateByID(ctx, l.data.GetMainDB(ctx))
}

func (l *userRepoImpl) DeleteByID(ctx context.Context, id uint32) error {
	userModel := &model.SysUser{
		ID: id,
	}
	return userModel.DeleteByID(ctx, l.data.GetMainDB(ctx))
}

func (l *userRepoImpl) UpdateStatusByIds(ctx context.Context, status vobj.Status, ids ...uint32) error {
	_, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysUser.Where(query.SysUser.ID.In(ids...)).Update(query.SysUser.Status, status)
	return err
}

func (l *userRepoImpl) UpdatePassword(ctx context.Context, id uint32, password types.Password) error {
	userModel := &model.SysUser{
		ID:       id,
		Password: password.String(),
		Salt:     password.GetSalt(),
	}
	return userModel.UpdateByID(ctx, l.data.GetMainDB(ctx))
}

func (l *userRepoImpl) Create(ctx context.Context, user *bo.CreateUserParams) (*model.SysUser, error) {
	userModel := createUserParamsToModel(user)
	if err := userModel.Create(ctx, l.data.GetMainDB(ctx)); err != nil {
		return nil, err
	}
	return userModel, nil
}

func (l *userRepoImpl) BatchCreate(ctx context.Context, users []*bo.CreateUserParams) error {
	userModels := types.SliceToWithFilter(users, func(item *bo.CreateUserParams) (*model.SysUser, bool) {
		if types.IsNil(item) || types.TextIsNull(item.Name) {
			return nil, false
		}
		return createUserParamsToModel(item), true
	})
	return query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysUser.CreateInBatches(userModels, 10)
}

func (l *userRepoImpl) GetByID(ctx context.Context, id uint32) (*model.SysUser, error) {
	return query.Use(l.data.GetMainDB(ctx)).SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(id)).First()
}

func (l *userRepoImpl) GetByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	return query.Use(l.data.GetMainDB(ctx)).SysUser.WithContext(ctx).Where(query.SysUser.Username.Eq(username)).First()
}

func (l *userRepoImpl) FindByPage(ctx context.Context, params *bo.QueryUserListParams) ([]*model.SysUser, error) {
	q := query.Use(l.data.GetMainDB(ctx)).SysUser.WithContext(ctx)

	var wheres []gen.Condition
	if !params.Status.IsUnknown() {
		wheres = append(wheres, query.SysUser.Status.Eq(params.Status))
	}
	if !params.Gender.IsUnknown() {
		wheres = append(wheres, query.SysUser.Gender.Eq(params.Gender))
	}
	if !params.Role.IsAll() {
		wheres = append(wheres, query.SysUser.Role.Eq(params.Role))
	}
	if !types.TextIsNull(params.Keyword) {
		q = q.Or(
			query.SysUser.Username.Like(params.Keyword),
			query.SysUser.Nickname.Like(params.Keyword),
			query.SysUser.Email.Like(params.Keyword),
			query.SysUser.Phone.Like(params.Keyword),
			query.SysUser.Remark.Like(params.Keyword),
		)
	}

	q = q.Where(wheres...)
	if !types.IsNil(params) {
		page := params.Page
		total, err := q.Count()
		if err != nil {
			return nil, err
		}
		params.Page.SetTotal(int(total))
		pageNum, pageSize := page.GetPageNum(), page.GetPageSize()
		if pageNum <= 1 {
			q = q.Limit(pageSize)
		} else {
			q = q.Offset((pageNum - 1) * pageSize).Limit(pageSize)
		}
	}
	return q.Order(query.SysUser.ID.Desc()).Find()
}

func (l *userRepoImpl) UpdateUser(ctx context.Context, user *model.SysUser) error {
	return user.UpdateByID(ctx, l.data.GetMainDB(ctx))
}

func (l *userRepoImpl) FindByIds(ctx context.Context, ids ...uint32) ([]*model.SysUser, error) {
	return query.Use(l.data.GetMainDB(ctx)).SysUser.WithContext(ctx).Where(query.SysUser.ID.In(ids...)).Find()
}

// createUserParamsToModel create user params to model
func createUserParamsToModel(user *bo.CreateUserParams) *model.SysUser {
	if types.IsNil(user) {
		return nil
	}
	return &model.SysUser{
		Username: user.Name,
		Nickname: user.Nickname,
		Password: user.Password.String(),
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   user.Status,
		Remark:   user.Remark,
		Avatar:   user.Avatar,
		Salt:     user.Password.GetSalt(),
		Gender:   user.Gender,
	}
}

// updateUserParamsToModel update user params to model
func updateUserParamsToModel(user *bo.UpdateUserParams) *model.SysUser {
	if types.IsNil(user) {
		return nil
	}
	userModel := createUserParamsToModel(&user.CreateUserParams)
	userModel.ID = user.ID
	return userModel
}
