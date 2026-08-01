# LitePan-fpk

Auto-build [LitePan](https://github.com/Ponphil/LitePan) into fnOS .fpk packages for x86 and arm64.

## Structure

`
.github/workflows/build-fpk.yml   GitHub Action workflow
internal/fusemount/constants.go   Fork: MountRoot as env var (LITEPAN_MOUNT_ROOT)
LitePan/                           App store listing assets (icon, preview, README)
LitePan-arm/                       fnOS app template (arm64)
LitePan-x86/                       fnOS app template (x86)
fnpack/                            fnpack packaging tool
fnpack.json                        App store metadata (auto-updated by CI)
`

## How it works

1. Clones original LitePan source
2. Overrides `constants.go` with this fork's version (env var support for MountRoot)
3. Cross-compiles linux/amd64 + linux/arm64 binaries (`-tags fuse`, CGO disabled)
4. Copies binaries into `LitePan-{arm,x86}/app/bin/`
5. Builds .fpk packages with fnpack

## Trigger

Push to `main` triggers the build. Manual dispatch also available.
Artifacts are uploaded; manual runs also create a GitHub release.
