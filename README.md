# LW

## What is this

*lw* is a package to enable selective logging.  While it is not zero-cost, it should do a little better than setting the log.Writer to ioutil.Discard.

### Todo items

- [x] INFO, WARNING, ERROR, FATAL, CONSOLE
- [x] look at standard log formats
- [ ] design to hookup easily to `https://goaccess.io`
- [x] consider an approach where message format can be specified (call location/line)
