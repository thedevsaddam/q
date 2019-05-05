# q

### Installation
```sh
Coming soon!
```


### Examples

**Find**
```sh
cat data.json | q -pretty -find="items.[0].name"
````

**Get**
```sh
cat data.json | q -from="items" -pretty
```

```sh
cat data.json | q -from="items.[0].name"
```

**Command**
```sh
cat data.json | q -from="prices" -command="first" #first,last,count,avg etc
```

**Aggregate functions**
```sh
cat data.json | q -from="prices" -command="avg" #sum,min,max,avg,count
```

```sh
cat data.json | q --from="items" -command="min:price"
```

**Offset**
```sh
cat data.json | q -pretty -from="items" -offset="2"
```

**Limit**
```sh
cat data.json | q -pretty -from="items" -limit="4"
```

**Distinct**
```sh
cat data.json | q -from="items" -pretty -distinct="price"
```

**Pluck**
```sh
cat data.json | q -pretty -from="items" -command="pluck:name"
```

**Select/Columns**
```sh
cat data.json | q -from="items" -pretty -columns="name,price"
```

**Where**
```sh
cat data.json | q --from="items" --where="name=Fujitsu"
```

**OrWhere**
```sh
cat data.json | q --from="items" --where="name=Fujitsu" --orWhere="id=int:1"
```

**GroupBy**
```sh
cat data.json | q -pretty -from="items" -groupBy="price"
```

**Sort [array]**
```sh
cat data.json | q -pretty -from="prices" -sort="desc"
```

**SortBy [array of objects]**
```sh
cat data.json | q -pretty -from="items" -sortBy="price:desc"
```

**Type [JSON, XML, YAML, CSV]**
```sh
cat data.csv| q -type="csv" -pretty -where="ISBN=9781846053443" -columns="ASP,ISBN,RRP" # type can be "csv", "xml", "yaml", "json"; default is "json"
```