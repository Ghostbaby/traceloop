SHELL=/bin/bash -o pipefail
DEST_DIR?=/dist
LINUX_HEADERS=$(shell rpm -q kernel-devel --last | head -n 1 | awk -F'kernel-devel-' '{print "/usr/src/kernels/"$$2}' | cut -d " " -f 1)

build:
	@mkdir -p "$(DEST_DIR)"
	# main
	clang -D__KERNEL__ -D__ASM_SYSREG_H \
		-DCIRCLE_BUILD_URL=\"$(CIRCLE_BUILD_URL)\" \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Wunused \
		-Wall \
		-Werror \
		-O2 -emit-llvm -c bpf/straceback-main-bpf.c \
		$(foreach path,$(LINUX_HEADERS), -I $(path)/arch/x86/include -I $(path)/arch/x86/include/generated -I $(path)/include -I $(path)/include/generated/uapi -I $(path)/arch/x86/include/uapi -I $(path)/include/uapi) \
		-o - | llc -march=bpf -filetype=obj -o "${DEST_DIR}/straceback-main-bpf.o"
	# tailcall
	clang -D__KERNEL__ -D__ASM_SYSREG_H \
		-DCIRCLE_BUILD_URL=\"$(CIRCLE_BUILD_URL)\" \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Wunused \
		-Wall \
		-Werror \
		-O2 -emit-llvm -c bpf/straceback-tailcall-bpf.c \
		$(foreach path,$(LINUX_HEADERS), -I $(path)/arch/x86/include -I $(path)/arch/x86/include/generated -I $(path)/include -I $(path)/include/generated/uapi -I $(path)/arch/x86/include/uapi -I $(path)/include/uapi) \
		-o - | llc -march=bpf -filetype=obj -o "${DEST_DIR}/straceback-tailcall-bpf.o"
	# guess
	clang -D__KERNEL__ -D__ASM_SYSREG_H \
		-DCIRCLE_BUILD_URL=\"$(CIRCLE_BUILD_URL)\" \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Wunused \
		-Wall \
		-Werror \
		-O2 -emit-llvm -c bpf/straceback-guess-bpf.c \
		$(foreach path,$(LINUX_HEADERS), -I $(path)/arch/x86/include -I $(path)/arch/x86/include/generated -I $(path)/include -I $(path)/include/generated/uapi -I $(path)/arch/x86/include/uapi -I $(path)/include/uapi) \
		-o - | llc -march=bpf -filetype=obj -o "${DEST_DIR}/straceback-guess-bpf.o"
	# bindata
	go-bindata -pkg straceback -prefix "${DEST_DIR}/" -modtime 1 -o "${DEST_DIR}/straceback-assets-bpf.go" \
		"${DEST_DIR}/straceback-guess-bpf.o" \
		"${DEST_DIR}/straceback-main-bpf.o" \
		"${DEST_DIR}/straceback-tailcall-bpf.o"
