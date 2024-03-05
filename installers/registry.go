package installers

import (
	"github.com/LinuxUserGD/hush"
	"github.com/LinuxUserGD/hush/archive"
	"github.com/LinuxUserGD/hush/naked"
)

func GetManager(typ hush.InstallerType) hush.Manager {
	switch typ {
	case hush.InstallerTypeArchive:
		return &archive.Manager{}
	case hush.InstallerTypeNaked:
		return &naked.Manager{}
	default:
		return nil
	}
}
