source ../../../support/dev/hurtlocker-dev-base-plan.sh

do_dev_prepare() {
  # Order matters here
  do_default_prepare
  PLAN_CONTEXT="../habitat"
}

do_prepare() {
  do_dev_prepare
}

do_hurtlocker_install() {
  local pkg_path
  pkg_path=$(hab pkg path chuckleheads/"$pkg_name")

  build_line "Linking new binary into package"
  ln -sfv "${GOPATH}/bin/$pkg_name" "$pkg_path/bin/$pkg_name"

  # build_line "Copying new config into package"
  # cp -v "$PLAN_CONTEXT/default.toml" "$pkg_path/default.toml"
  # cp -v "$PLAN_CONTEXT/config/config.toml" "$pkg_path/config/config.toml"

  # build_line "Copying run hooks into package"
  # for hook in "$PLAN_CONTEXT"/hooks/*; do
  #   cp -v "$hook" "$pkg_path/hooks/$(basename "$hook")"
  # done
}

do_install_wrapper() {
  do_install
}
