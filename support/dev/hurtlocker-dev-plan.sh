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

  build_line "Copying new config into package"
  if [ -f "$PLAN_CONTEXT/default.toml" ]; then
    cp -v "$PLAN_CONTEXT/default.toml" "$pkg_path/default.toml"
  fi
  if [ -f "$PLAN_CONTEXT/config/config.toml" ]; then
    cp -v "$PLAN_CONTEXT/config/config.toml" "$pkg_path/config/config.toml"
  fi
	if [ -d "$PLAN_CONTEXT/hooks" ]; then
		build_line "Copying run hooks into package"
		for hook in "$PLAN_CONTEXT"/hooks/*; do
			cp -v "$hook" "$pkg_path/hooks/$(basename "$hook")"
		done
	fi
  if [ -d "$PLAN_CONTEXT/../migrations" ]; then
		build_line "Copying migrations into package"
    cp -r "$PLAN_CONTEXT/../migrations" "${pkg_path}/"
  fi
}

do_install_wrapper() {
  do_install
}
