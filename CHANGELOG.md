# Changelog

## [1.1.0](https://github.com/imnotjames/kube-state-healthz/compare/v1.1.0-rc.1...v1.1.0) (2024-07-09)


### Bug Fixes

* write to stderr ([b8f034a](https://github.com/imnotjames/kube-state-healthz/commit/b8f034a7a1ade7d0cfccaf29ea7d1719aa89d126))


### Miscellaneous Chores

* release 1.1.0 ([8855bc1](https://github.com/imnotjames/kube-state-healthz/commit/8855bc1f5c4064fa5598c9690783384dba3955b0))

## [1.1.0-rc.1](https://github.com/imnotjames/kube-state-healthz/compare/v1.0.1...v1.1.0-rc.1) (2024-07-08)


### Features

* support env vars for all flags ([#9](https://github.com/imnotjames/kube-state-healthz/issues/9)) ([98b298a](https://github.com/imnotjames/kube-state-healthz/commit/98b298a98830cf4143f82f3e377450f22fe058ac))


### Bug Fixes

* proper print to stdout ([2fa6f41](https://github.com/imnotjames/kube-state-healthz/commit/2fa6f41fc85368a37b317c55de05bda60b1b2732))
* proper print to stdout ([27c0064](https://github.com/imnotjames/kube-state-healthz/commit/27c00642aba29b8d1a42195cd182ff07b0069db8))
* use cobra stdout ([#11](https://github.com/imnotjames/kube-state-healthz/issues/11)) ([9285314](https://github.com/imnotjames/kube-state-healthz/commit/9285314100f433b5aa614f024965199ed6fb01f5))
* use correct var for error ([2c0a279](https://github.com/imnotjames/kube-state-healthz/commit/2c0a2793ccbed40ec969e7b443bf8f2485b9aa71))

## [1.0.1](https://github.com/imnotjames/kube-state-healthz/compare/v1.0.0...v1.0.1) (2024-07-08)


### Bug Fixes

* handle in-cluster config without crashing ([72eb5fe](https://github.com/imnotjames/kube-state-healthz/commit/72eb5fed15a4510208506d5b3383a1dd25d6c57e))

## 1.0.0 (2024-07-08)


### Features

* add dockerfile ([4458874](https://github.com/imnotjames/kube-state-healthz/commit/44588743a10dacf2242b4ab831171d74ed4bb96f))
* add label selector support ([e1d9e98](https://github.com/imnotjames/kube-state-healthz/commit/e1d9e986a2f2e2a023b41ee85f1a31015e453f6b))
* log requests ([a727807](https://github.com/imnotjames/kube-state-healthz/commit/a72780730cdc58151a6697f0b4c34183a63b293e))
* use a markdown table in `check` command ([75cef6b](https://github.com/imnotjames/kube-state-healthz/commit/75cef6b17b702c932539d2d079fca4e81d2a5947))
* use cli option for kubeconfig when specified ([3e9cf04](https://github.com/imnotjames/kube-state-healthz/commit/3e9cf04d89162cd3e8024d29f4eea8a3d64be47d))


### Bug Fixes

* consistency in flag descriptions ([67285e9](https://github.com/imnotjames/kube-state-healthz/commit/67285e973b29cfe5293dfc746e0add8905799510))
* correct the dockerfile to properly run the bin & serve ([c1e089b](https://github.com/imnotjames/kube-state-healthz/commit/c1e089bf65a12b5004c8d4876b5967568ec172a6))
* handle empty state for default kube config ([da6ba5d](https://github.com/imnotjames/kube-state-healthz/commit/da6ba5db8a0cb9eb4be8f4ed3c27fda2a90ee63c))
* use http status code for no content success ([5b24f70](https://github.com/imnotjames/kube-state-healthz/commit/5b24f704467a56018bda2d7c2fafe7d6d43bfaf6))
