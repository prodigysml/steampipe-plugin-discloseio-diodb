# Steampipe Plugin Disclose.io Diodb

This is a steampipe plugin to parse the JSON structure within disclose.io's diodb project.

## Requirements
1. `steampipe`
2. `golang`

## Installation Instructions
1. Clone this repo (`git clone`) and `cd` into the directory
2. Run `make install`. This will copy the go binary into the right place and the appropriate config file
3. Run `steampipe query` to open the steampipe console
4. Run your queries as you normally would! Below is an example:

```sql
select * from diodb where contact_url is not NULL;
```

## Clean / Uninstall
1. Run `make clean`