
MEMORY
{
    FLASH_TEXT (rw) : ORIGIN = 0x00808000, LENGTH = 96K
    RAM (xrw)       : ORIGIN = 0x20020000, LENGTH = 128K
}

_stack_size = 8K;

INCLUDE "targets/arm.ld"
