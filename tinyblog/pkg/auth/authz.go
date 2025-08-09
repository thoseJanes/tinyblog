package auth

import (
	"time"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)


const(
	aclModel = `[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

type Authz struct{
	Enforcer *casbin.SyncedEnforcer
}

func NewAuthz(db *gorm.DB) (*Authz, error) {
	adpt, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	md, err := model.NewModelFromString(aclModel)
	if err != nil {
		return nil, err // 这里是不是直接abort比较好？（模型创建失败？模型内容有误？）如果模型string正确，该错误是否根本不可能发生？因此不需要处理？
	}

	enforcer, err := casbin.NewSyncedEnforcer(md, adpt)
	if err != nil {
		return nil, err
	}

	if err = enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	enforcer.StartAutoLoadPolicy(time.Second * 5)

	return &Authz{enforcer}, nil
}


func (authz *Authz) Authorize(sub, obj, act string) (bool, error) {
	return authz.Enforcer.Enforce(sub, obj, act)
}