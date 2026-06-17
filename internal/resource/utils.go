package resource

import (
	"fmt"

	"github.com/BitMancers/Portigo/internal/utils"
)

func nextDev(bus DiskBusType, currentIndex int) string {
	switch bus {
	case DiskTargetBusVirtIO:
		return fmt.Sprintf("vd%c", 'a'+currentIndex)
	case DiskTargetBusSCSI, DiskTargetBusSATA, "usb":
		return fmt.Sprintf("sd%c", 'a'+currentIndex)
	case DiskTargetBusIDE:
		return fmt.Sprintf("hd%c", 'a'+currentIndex)
	case "nvme":
		return fmt.Sprintf("nvme%dn1", currentIndex)
	case "xen":
		return fmt.Sprintf("xvd%c", 'a'+currentIndex)
	default:
		return fmt.Sprintf("vd%c", 'a'+currentIndex)
	}
}

func RandomMacAddress() string {
	randStr := utils.RandStringRunes(12)
	out := make([]rune, 17)
	for i := range 6 {
		out[i*2] = rune(randStr[i*2])
		out[(i*2)+1] = rune(randStr[(i*2)+1])
		if i < 5 {
			out[(i*2)+2] = ':'
		}
	}
	return string(out)
}
