source ../../../support/dev/hurtlocker-base-plan.sh

pkg_name=tasker
pkg_origin=chuckleheads
pkg_description="Tasker runs tasks and streams the results back to the api"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_deps=(core/hab)
pkg_build_deps=(core/go core/git core/dep)
