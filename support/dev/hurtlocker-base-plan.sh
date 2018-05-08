#!/bin/bash

export GOPATH="${HAB_CACHE_SRC_PATH}/go"
export workspace_src="${GOPATH}/src"
export base_path="github.com/chuckleheads/hurtlocker"
export pkg_cache_path="${workspace_src}/${base_path}"

pkg_version() {
  echo "$(($(git rev-list master --count)))"
}

do_before() {
  do_hurtlocker_before
}

do_download() {
  do_hurtlocker_download
}

do_build() {
  do_hurtlocker_build
}

do_install() {
  do_hurtlocker_install
}

do_strip() {
  return 0
}

do_clean() {
  do_hurtlocker_clean
}

do_hurtlocker_before() {
  mkdir -p "$pkg_cache_path"
  update_pkg_version
}

do_hurtlocker_download() {
  cp -r "${PLAN_CONTEXT}/../../../"* "$pkg_cache_path"
}

do_hurtlocker_build() {
  pushd "${pkg_cache_path}/components/${pkg_name}" >/dev/null
  GOOS=linux go build -o "${GOPATH}/bin/${pkg_name}" .
  popd >/dev/null
}

do_hurtlocker_install() {
  cp -r "${GOPATH}/bin" "${pkg_prefix}/${bin}"
}

do_hurtlocker_clean() {
  do_default_clean
  rm -rf "${GO_PATH}"
}
