source ../../../support/dev/hurtlocker-base-plan.sh

pkg_name=gateway
pkg_origin=chuckleheads
pkg_description="Gateway server to proxy http requests to a gRPC service"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_build_deps=(core/go core/git core/dep)
pkg_svc_run="${pkg_name} start"
