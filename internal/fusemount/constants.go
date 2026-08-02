package fusemount

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultMountRoot     = "/app/mounts"
	KeyEnabled           = "fuse_enabled"
	DefaultEntryTimeoutS = 30
	DefaultAttrTimeoutS  = 3
)


var MountRoot = resolveMountRoot()

func resolveMountRoot() string {
	if v := strings.TrimSpace(os.Getenv("LITEPAN_MOUNT_ROOT")); v != "" {
		return filepath.Clean(v)
	}
	return defaultMountRoot
}
