//go:build cortexm && !nxp && !qemu && abort.bkpt
// +build cortexm,!nxp,!qemu,abort.bkpt

package runtime

import (
	"device/arm"
)

func exit(code int) {
	abort()
}

func abort() {
	for {
		arm.Asm("bkpt")
	}
}
