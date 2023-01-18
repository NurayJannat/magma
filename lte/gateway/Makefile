format_all:
	find $(MAGMA_ROOT)/orc8r/gateway/c/ \( -iname "*.c" -o -iname "*.cpp" -o -iname "*.h" -o -iname "*.hpp" \) -exec \
	/usr/bin/clang-format-11 -i {} \;
	find $(MAGMA_ROOT)/lte/gateway/c/ \( -iname "*.c" -o -iname "*.cpp" -o -iname "*.h"  -o -iname "*.hpp" \) -exec \
	/usr/bin/clang-format-11 -i {} \;
	find $(MAGMA_ROOT)/lte/gateway/python/ \( -iname "*.c" -o -iname "*.cpp" -o -iname "*.h"  -o -iname "*.hpp" \) -exec \
	/usr/bin/clang-format-11 -i {} \;
