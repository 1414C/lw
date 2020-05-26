# LW

<a href="https://github.com/1414C/lw/actions?query=workflow%3Abuild-test" alt="BuildStatus">
        <img src="https://github.com/1414C/lw/workflows/build-test/badge.svg" /></a>

<a href="https://github.com/1414C/lw/releases" alt="Releases">
        <img src="https://img.shields.io/github/v/release/1414C/lw" /></a>

<a href="https://golang.org/dl/" alt="GoVersion">
        <img src="https://img.shields.io/github/go-mod/go-version/1414C/lw" /></a>

## What is this

*lw* is a package to enable selective logging.  While it is not zero-cost, it should do a little better than setting the log.Writer to ioutil.Discard.

### Todo items

- [x] INFO, WARNING, ERROR, FATAL, CONSOLE
- [x] look at standard log formats
- [ ] design to hookup easily to `https://goaccess.io`
- [x] consider an approach where message format can be specified (call location/line)
