# see also https://gist.github.com/blockpane/fe03eb0839fac417b92cd7eb98cdf356

OUT_DIR=output

APP_NAME="${APP_NAME:-peloche}"
APP_VERSION="${APP_VER:-v0.0.0}"
APP_VERSION_XYZ=${APP_VERSION#v} # remove the prefix 'v' if given (in the tag name)

echo "Packaging app: $APP_NAME"
echo "Version: $APP_VERSION"
echo "Output directory: $OUT_DIR"

mkdir -p "$OUT_DIR"
rm -rf "$OUT_DIR/$APP_NAME-*.app"
rm -f "$OUT_DIR/$APP_NAME-*.dmg"

fyne package \
  -name "$APP_NAME" \
  --app-version $APP_VERSION_XYZ \
  --os darwin \
  --icon etc/icon.png

APP_PATH="$OUT_DIR/$APP_NAME-$APP_VERSION-macos.app"
DMG_PATH="$OUT_DIR/$APP_NAME-$APP_VERSION-macos.dmg"

mv $APP_NAME.app "$APP_PATH"
xattr -rc "$APP_PATH"

echo ".app packaged in '$APP_PATH'"

hdiutil create -srcfolder "$APP_PATH" "$DMG_PATH"

echo ".dmg packaged in '$DMG_PATH'"
