pkg_name=graph-datastore-ratel
pkg_origin=chuckleheads
pkg_version=0.0.1
pkg_description="dgraph UI database for builder chuckleheads"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_deps=(core/dgraph-ratel)
pkg_svc_run="dgraph-ratel"

do_build() {
  return 0
}

do_install() {
  return 0
}
