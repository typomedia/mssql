# `mssql` - Microsoft SQL Server CLI

This is a simple CLI to execute queries and backup/restore databases on Microsoft SQL Server. It is written in [Go](https://go.dev/) 
and uses the [Cobra](https://cobra.dev/) library. 

## Motivation

It's a single binary and drop-in replacement for the `sqlcmd` command line tool without the need of installing the SQL Server client tools.

## Usage

    mssql [command]

    backup      Backup a database
    completion  Generate the autocompletion script for the specified shell
    exec        Execute a statement
    help        Help about any command
    info        Show server information
    list        List all databases
    query       Execute a query
    restore     Restore a database
    version     Show the current version

## Flags

    -h, --help   help for mssql

## Examples

### `mssql info`
```
mssql info -u sa -w password
Version                    Edition                     
Microsoft SQL Server 2022  Developer Edition (64-bit)
```

### `mssql list`
```
mssql list -u sa -w password
Name        File                                
demodb      /var/opt/mssql/data/demodb.mdf  
testdb      /var/opt/mssql/data/testdb.mdf
```

### `mssql query`
```
mssql query -u sa -w password -q "SELECT cpu_count FROM sys.dm_os_sys_info"
cpu_count  
4     
```

### `mssql exec`
```
mssql exec -u sa -w password -q "CREATE DATABASE test"
```

### `mssql backup`
```
mssql backup -u sa -w password -d testdb /tmp/testdb.bak
```

### `mssql restore`
```
mssql restore -u sa -w password -d testdb /tmp/testdb.bak
```

## Help

    mssql -h

## Build
    
    make build

---
Copyright © 2024 Typomedia Foundation. All rights reserved.