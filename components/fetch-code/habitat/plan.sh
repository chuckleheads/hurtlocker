source ../../../support/dev/hurtlocker-base-plan.sh

pkg_name=fetch-code
pkg_origin=chuckleheads
pkg_description="fetch code downloads git repos"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_build_deps=(core/go core/git core/dep)
