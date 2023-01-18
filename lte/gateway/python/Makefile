include $(MAGMA_ROOT)/lte/gateway/python/defs.mk
# Contains common targets for gateway python development
include $(MAGMA_ROOT)/orc8r/gateway/python/python.mk

define DEPRECATION_WARNING

Warning: As part of the Make to Bazel switch-over of the AGW this Makefile
will be deprecated soon, see https://github.com/magma/magma/issues/14971.
Do not add any new Make targets for the AGW!
All new targets should be integrated with the Bazel build system,
see https://magma.github.io/magma/docs/next/bazel/agw_with_bazel.

endef
export DEPRECATION_WARNING

# Set up some pattern rules
define setup_rules
	$(eval PSRC := $1)
	CLEAN_LIST += $(PSRC)_clean
	BUILD_LIST += $(PSRC)_build
	TEST_LIST += $(PSRC)_test
endef
_ := $(foreach python_src, $(PYTHON_SRCS), $(eval $(call setup_rules,$(python_src))))

snowflake:
	@echo "$$DEPRECATION_WARNING"
	sudo env PATH=${PYTHON_BUILD}/bin:${PATH} snowflake --make-snowflake

buildenv: setupenv protos swagger $(BUILD_LIST) py_patches snowflake
$(BUILD_LIST): %_build:
	@echo "$$DEPRECATION_WARNING"
	make -C $* install_egg

$(BIN)/pytest: install_virtualenv
	$(VIRT_ENV_PIP_INSTALL) pytest==7.1.2

$(BIN)/pylint: install_virtualenv
	$(VIRT_ENV_PIP_INSTALL) pylint==2.14.0

# Disable W0511: todo warnings
# Disable R0903: Too few public methods

check:
	./precommit.py --format --diff

clean: $(CLEAN_LIST)
	@echo "$$DEPRECATION_WARNING"
	sudo rm -rf $(PYTHON_BUILD)/
	sudo find . -name '*.pyc' -o -name '__pycache__' -prune -exec rm -rf {} \;
$(CLEAN_LIST): %_clean:
	make -C $* remove_egg
