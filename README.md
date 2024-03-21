# Greta <img alt="Greta Logo" src="docs/favicon.svg" height="96" align="right"/>

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

[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![VectorLogoZone](https://www.vectorlogo.zone/logos/vectorlogozone/vectorlogozone-ar21.svg)](https://www.vectorlogo.zone/ "Logos in credits")

* docker
* Fxemoji
