# Vendor test

Testing vendoring in combination with a `cmd` directory.

Our project root is a library which provides a function that needs a logrus logger.

We also provide a command/binary `vtest` that vendors its dependencies. ...**But** it doesn't vendor the parent library, because that would be a bit crazy.

The end result is that the `github.com/Sirupsen/logrus` dependency in the library doesn't resolve to the vendored dependency for the command. So we either get a "cannot find package", or a type mismatch if logrus is in the GOPATH:

```
# github.com/Sydsvenskan/vendortest/cmd/vtest
./vtest.go:9: cannot use "github.com/Sydsvenskan/vendortest/cmd/vtest/vendor/github.com/Sirupsen/logrus".StandardLogger() (type *"github.com/Sydsvenskan/vendortest/cmd/vtest/vendor/github.com/Sirupsen/logrus".Logger) as type *"github.com/Sirupsen/logrus".Logger in argument to vendortest.StartWithLogger
```
