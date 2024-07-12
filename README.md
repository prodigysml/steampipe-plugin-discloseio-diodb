# Steampipe Plugin Disclose.io Diodb

This is a steampipe plugin to parse the JSON structure within disclose.io's diodb project.

## Installation Instructions
1. Ensure `steampipe` is installed
2. Ensure `golang` is installed
3. Clone this repo and `cd` into the directory
4. Run `make install`. This will copy the go binary into the right place and the appropriate config file
5. Run `steampipe query` to open the steampipe console
6. Run your queries as you normally would! Below is an example:

```sql
select * from diodb where contact_url is not NULL;
```

## Clean
1. Run `make clean`