# go-validation-util

Package validation exposes some utility functions to validate strings and
numbers. Validation error messages are formatted in the form of
`fieldName,errorCode`. `errorCode` may contain additional information, such as
min/max length, etc.

## TODO

- [ ] Implement checking of `m` in `format` (`m,n`) of `NumberFormat(field string, value float64, format string)`
- [ ] Unit tests
- [ ] Implement validation with struct tags

