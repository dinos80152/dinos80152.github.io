# AWS Athena Create Table Manually

1. Create Table

    ```sql
    CREATE EXTERNAL TABLE `20210121_test_log`(
      `uuid` string,
      `hostid` string,
      `hostname` string,
      `osname` string)
    PARTITIONED BY (
      `x` string,
      `year` string,
      `month` string,
      `day` string)
    ROW FORMAT SERDE
      'org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe'
    STORED AS INPUTFORMAT
      'org.apache.hadoop.mapred.TextInputFormat'
    OUTPUTFORMAT
      'org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat'
    LOCATION
      's3://dinolai-test-log/'
    TBLPROPERTIES (
      'has_encrypted_data'='false',
      'transient_lastDdlTime'='1611144510')
    ```

1. Add Partition

    ```sql
    ALTER TABLE `20210121_test_log` ADD
      PARTITION (x = '0', year = '2021', month = '01', day = '20')
      PARTITION (x = '1', year = '2021', month = '01', day = '20')
      PARTITION (x = '2', year = '2021', month = '01', day = '20')
      PARTITION (x = '3', year = '2021', month = '01', day = '20')
      PARTITION (x = '4', year = '2021', month = '01', day = '20')
      PARTITION (x = '5', year = '2021', month = '01', day = '20')
      PARTITION (x = '6', year = '2021', month = '01', day = '20')
      PARTITION (x = '7', year = '2021', month = '01', day = '20')
      PARTITION (x = '8', year = '2021', month = '01', day = '20')
      PARTITION (x = '9', year = '2021', month = '01', day = '20')
      PARTITION (x = 'a', year = '2021', month = '01', day = '20')
      PARTITION (x = 'b', year = '2021', month = '01', day = '20')
      PARTITION (x = 'c', year = '2021', month = '01', day = '20')
      PARTITION (x = 'd', year = '2021', month = '01', day = '20')
      PARTITION (x = 'e', year = '2021', month = '01', day = '20')
      PARTITION (x = 'f', year = '2021', month = '01', day = '20')
    ;
    ```

1. Query

    ```sql
    SELECT osname, hostname, count(distinct(hostid)) as host_count, count(*) as log_count
    FROM "default"."20210121_test_log"
    WHERE osname = 'linux'
    GROUP BY osname;
    ```
