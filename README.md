# Greta [<img alt="Greta Logo" src="docs/favicon.svg" height="96" align="right"/>](https://greta.marcuse.info/)

[![build](https://github.com/fileformat/greta/actions/workflows/build.yaml/badge.svg)](https://github.com/fileformat/greta/actions/workflows/build.yaml)
[![example](https://github.com/fileformat/greta/actions/workflows/example.yaml/badge.svg)](https://github.com/fileformat/greta/blob/main/.github/workflows/example.yaml)

**G**o **RE**nder **T**emplate **A**ction: A Github Action that renders a Go template.

## Action Inputs

| Name       | Description          | Notes                             |
| ---------- | -------------------- | --------------------------------- |
| `template` | The path to the Go template file |  |
| `input`     | A `.json` file (or `-` for stdin) | default is `-` (stdin) |
| `mode`     | `html` to use [html/template](https://pkg.go.dev/html/template) and `text` for [text/template](https://pkg.go.dev/text/template) | default is `html` |
| `output`   | The path to the output file.  | stdout if not specified |

## License

[MIT](LICENSE.txt)

## Credits

[![Docker](https://www.vectorlogo.zone/logos/docker/docker-ar21.svg)](https://www.docker.com/ "Deployment")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting and CI")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Mozilla](https://www.vectorlogo.zone/logos/mozilla/mozilla-ar21.svg)](https://github.com/mozilla/fxemoji "Logo")
[![VectorLogoZone](https://www.vectorlogo.zone/logos/vectorlogozone/vectorlogozone-ar21.svg)](https://www.vectorlogo.zone/ "Logos in credits")
