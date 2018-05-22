pkg_name=router
pkg_origin=chuckleheads
pkg_version=0.0.1
pkg_description="envoy proxy for builder chuckleheads"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_deps=(core/envoy)
pkg_svc_user="root"
pkg_svc_group="root"
# Since Habitat doesn't have a way to declare a bind *for* another service
# We have to list all permutations here 
pkg_binds_optional=(
  [originsrv]="port"
  [sessionsrv]="port"
)
# pkg_exports=(
#   [port]=port
# )
# pkg_exposes=(port)

do_build() {
  return 0
}

do_install() {
  return 0
}
