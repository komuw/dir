# Release Notes

Most recent version is listed first.  


## v0.0.3
- Update CI: https://github.com/komuw/kama/pull/17   
- Dump more information about variables/types: https://github.com/komuw/kama/pull/18      
                                             : https://github.com/komuw/kama/pull/21       
- Implement own `dump` functionality: https://github.com/komuw/kama/pull/22     
  We used to use `sanity-io/litter` to do dumping.      
  We however, decided to implement our own dump functionality.       
  The main reason precipitating we are doing this is because sanity-io/litter has no way to compact       
  arrays/slices/maps that are inside structs.        

## v0.0.2
- add test example: https://github.com/komuw/kama/pull/13
- add types to the fields of a struct: https://github.com/komuw/kama/pull/16

## v0.0.1
- pretty print variables and packages: https://github.com/komuw/kama/pull/10
- add cli: https://github.com/komuw/kama/pull/11
- add pretty printing for data structures: https://github.com/komuw/kama/pull/12
