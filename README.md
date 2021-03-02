# NBP CLI tool

CLI tool for calling API of National Bank of Poland (NBP)

It gets exchange rate for given currency from tables (A, B, C) published by [NBP](https://www.nbp.pl/home.aspx?f=/kursy/kursy_archiwum.html) by using [NBP Web API](http://api.nbp.pl/).

## Installation

Download binary from [latest release](https://github.com/jfilipczyk/nbp-cli/releases/latest), extract it and add to your $PATH.

## How to use it

To get latest average rate for EUR (table A):
```bash
nbp ex eur
```

To get buy/sell rate for EUR (table C) for given date:
```bash
nbp ex eur -t c -d 2020-01-10
```

**NBP does not publish exchange rates every day**, only on working days (tables A, C) and on working Wednesdays (table B).
If you query for a date when rate was not published in a result you will get latest published rate before given date.

### Output format

Table format (default):
```bash
[kuba@localhost ~]$ nbp ex eur
code            EUR           
currency        euro          
effectiveDate   2021-03-04    
mid             4.554         
no              043/A/NBP/2021
table           A
```

JSON format:
```bash
[kuba@localhost ~]$ nbp ex eur -o json
{"code":"EUR","currency":"euro","effectiveDate":"2021-03-04","mid":4.554,"no":"043/A/NBP/2021","table":"A"}
```

To get single field with [jq](https://github.com/stedolan/jq):
```bash
[kuba@localhost ~]$ nbp ex eur -o json | jq .mid
4.554
```

## License
[MIT](./LICENSE)