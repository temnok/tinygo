
MEMORY
{
    FLASH_TEXT (rw) : ORIGIN = 0x00808000, LENGTH = 96K
    RAM (xrw)       : ORIGIN = 0x20020000, LENGTH = 64K
    SERIAL (xrw)    : ORIGIN = 0x20030000, LENGTH = 64K
}

_stack_size = 4K;
_serial_start = ORIGIN(SERIAL);

INCLUDE "targets/arm.ld"
