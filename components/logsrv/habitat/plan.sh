source ../../../support/dev/hurtlocker-base-plan.sh

pkg_name=logsrv
pkg_origin=chuckleheads
pkg_description="Log server receives log streams over gRPC"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_build_deps=(core/go core/git core/dep)
pkg_exports=(
  [port]=port
)
pkg_exposes=(port)
