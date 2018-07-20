# LW

## What is this

*lw* is a package to enable selective logging.  While it is not zero-cost, it should do a little better than setting the log.Writer to ioutil.Discard.

### Todo items

- [ ] INFO, WARNING, ERROR, FATAL, SUCCESS, REQUEST
- [ ] look at standard log formats
- [ ] design to hookup easily to `https://goaccess.io`
- [ ] consider an approach where message format can be specified (call location/line)

Now update as 'static' implementation

Private internal structure of enabled/disabled logging-types
- LWEnable(withLoc)
- LWDisable()
- InfoEnable(bool)
- WarningEnable(bool)
- TraceEnable(bool)
- DebugEnable(bool)
- ErrorEnable(bool)
- FatalEnable(bool)