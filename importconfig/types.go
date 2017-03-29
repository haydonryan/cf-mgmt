package importconfig

import (
	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	"github.com/pivotalservices/cf-mgmt/organization"
	"github.com/pivotalservices/cf-mgmt/space"
	"github.com/pivotalservices/cf-mgmt/uaa"
)

//Manager -
type Manager interface {
	ImportConfig(excludedOrgs map[string]string) error
}

//DefaultImportManager  -
type DefaultImportManager struct {
	ConfigDir       string
	UAACMgr         uaa.Manager
	OrgMgr          organization.Manager
	SpaceMgr        space.Manager
	CloudController cloudcontroller.Manager
}
