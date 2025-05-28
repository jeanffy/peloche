# Péloche

Pronounced "pɪlɒʃ".

"Péloche" is a french slang word, that means "reel" or "film roll".

This is a (very) lightweight image organization software. This project was initially created for me to learn the Go language.

The project is in prototype stage only, not suited for production use.

## Principle

This program is based upon the directory tree hierachy of photos/images. No database is created.

Only a sidecar .xmp file is created when notes or keywords are added/updated for an image.

## Photo editing

Photo editing is done in an external program that needs to be set up. By default the "PhotoScape X" editor is chosen (only suitable for macOS for now).

When editing a photo, the original file is backuped with a "-(original)" suffix. The edited photo can then be restored at any time without loosing any information.

## Files supported

Only file types "jpg/jpeg", "png" and "heic" are currently supported (based on the file extension).

## How to run

Clone this repo and run `go mod tidy` then `go run .`.

## Package

The app can be packaged with dedicated scripts for each platform:

- macOS: `sh etc/scripts/package-macos.sh`

The resulting app can be found in the `output` directory.

## Developer

The architecture is inspired by:

- the Clean Architecture principles, with two layers "domain" and "infra"
- the Ports & Adapters architecture

### Resources

If any resource is added to the `etc/resources` directory, the command `go generate` needs to be run to update the `infra/ui/assets/assets.go` file.

### External libraries used

- [Fyne](https://fyne.io) for the GUI
- https://github.com/adrium/goheif for HEIC decoding
- https://github.com/nfnt/resize for image resizing (TODO: find an alternative as the project is now archived)
- https://github.com/sqweek/dialog for native dialogs

### Publish a release

To trigger the packaging pipeline, when the main branch is stable, create and publish a tag:

- push all your commits to remote
- `sh etc/scripts/new-release.sh vX.Y.Z`
