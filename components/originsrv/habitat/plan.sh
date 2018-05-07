source ../../../support/dev/hurtlocker-base-plan.sh

pkg_name=originsrv
pkg_origin=chuckleheads
pkg_description="Origin Server is a gRPC server that deals with origin requests"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_build_deps=(core/go core/git core/dep)
pkg_deps=(core/cockroach)
pkg_svc_run="${pkg_name}"
