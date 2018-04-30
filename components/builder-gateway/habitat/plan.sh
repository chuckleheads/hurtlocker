pkg_name=builder-gateway
pkg_origin=chuckleheads
pkg_version=0.0.1
pkg_description="Gateway server to proxy http requests to a gRPC service"
pkg_upstream_url="https://github.com/chuckleheads/hurtlocker"
pkg_license=('Apache-2.0')
pkg_maintainer="The Habitat Maintainers <humans@habitat.sh>"
pkg_bin_dirs=(bin)
pkg_build_deps=(core/go core/git core/dep)
pkg_svc_run="${pkg_name}"

export GOPATH="${HAB_CACHE_SRC_PATH}/go"
export workspace_src="${GOPATH}/src"
export base_path="github.com/chuckleheads/hurtlocker"
export pkg_cache_path="${workspace_src}/${base_path}"

do_before() {
  mkdir -p "$pkg_cache_path"
}

do_download() {
  cp -r "${PLAN_CONTEXT}/../../../"* "$pkg_cache_path"
}

do_build() {
  pushd "${pkg_cache_path}/components/${pkg_name}/cmd/${pkg_name}" >/dev/null
  GOOS=linux go build -o "${GOPATH}/bin/${pkg_name}" .
  popd >/dev/null
}

do_install() {
  cp -r "${GOPATH}/bin" "${pkg_prefix}/${bin}"
}
