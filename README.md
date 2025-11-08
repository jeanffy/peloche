# Péloche

Pronounced "pɪlɒʃ".

"Péloche" is a french slang word, that means "reel" or "film roll".

This is a (very) lightweight image organization software. This project was initially created for me to learn the Go language.

The project is in prototype stage only, not suited for production use.

## Principle

This program is based upon the directory tree hierachy of photos/images. No database is created.

Only a sidecar .xmp file is created when notes or keywords are added/updated for an image.

## Commands

- install dependencies: `moon run install`
- run application in dev mode: `moon run dev`
- run application in prod mode: `moon run preview`
- lint application: `moon run lint`
  - auto-fix lint errors: `moon run lint-fix`
- cleanup: `moon run clean`
- build application for distribution:
  - windows: `npm run build-win`
  - linux: `npm run build-linux`
  - macOS: `npm run build-mac`
