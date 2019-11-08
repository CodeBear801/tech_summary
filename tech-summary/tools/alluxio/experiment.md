- [Prerequisite](#prerequisite)
- [Task0: Login to the AMI instance](#task0-login-to-the-ami-instance)
  - [MacOS or Linux Users](#macos-or-linux-users)
  - [Windows Putty Users](#windows-putty-users)
- [Task1: Mount a S3 bucket to Alluxio](#task1-mount-a-s3-bucket-to-alluxio)
- [Task2: Create a simple table and scan it in Presto](#task2-create-a-simple-table-and-scan-it-in-presto)
- [Task3: Run a more complicated query twice](#task3-run-a-more-complicated-query-twice)
- [(Bonus) Task4: Alluxio specific operations](#bonus-task4-alluxio-specific-operations)
  - [Check Alluxio Total Capacity Usage](#check-alluxio-total-capacity-usage)
  - [Check Mount Table in Alluxio](#check-mount-table-in-alluxio)
  - [Check How Much Data Cached Under a Directory](#check-how-much-data-cached-under-a-directory)
  - [Free Cache Capacity](#free-cache-capacity)
  - [Reference](#reference)

# Prerequisite 

* Address of the assigned instances. If you were not assigned with the address during your check-in, please talk to a TA.
* A working SSH client installed on your local laptop.
* You could create your own instance from aws and search for "Alluxio presto sandbox"

# Task0: Login to the AMI instance

## MacOS or Linux Users

  * Download a passphrase-protected [private key (pem file)](https://alluxio-training-1.s3-ap-southeast-1.amazonaws.com/summit.pem)
  * Update the permission of the downloaded private key
 
  ```console
  $ chmod 600 ~/Downloads/summit.pem
  ```     
  * login to the instance, entering the passphrase (asking a TA if you don't know the passphrase) when prompted.

  ```console
  $ ssh -i ~/Downloads/summit.pem ec2-user@<instance-address>
  Enter passphrase for key 'summit.pem':
  ```
  
  * By completing task0, you are expected to see the welcome message similar to the following:

```

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
29 package(s) needed for security, out of 55 available
Run "sudo yum update" to apply all updates.
[ec2-user@ip-172-31-31-241 ~]$ 
```
  
## Windows Putty Users

  * Download the [private key (ppk file for Putty)](https://alluxio-training-1.s3-ap-southeast-1.amazonaws.com/summit.ppk)
  * Login to the instance following [instructions](https://docs.aws.amazon.com/en_pv/AWSEC2/latest/UserGuide/putty.html#putty-ssh). Remember to fill in "Host Name (or IP Address)" field in "Session" with `ec2-user@<instance-address>`, and select the downloaded .ppk file in "Connection->SSH->Auth"
  * By completing task0, you are expected to see the welcome message similar to the following:

```

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
29 package(s) needed for security, out of 55 available
Run "sudo yum update" to apply all updates.
[ec2-user@ip-172-31-31-241 ~]$ 
```

# Task1: Mount a S3 bucket to Alluxio 

Mount an S3 bucket which has example input data for this tutorial to Alluxio file system. Note that this S3 bucket is a public bucket accessible to all AWS users.

```console
$ alluxio fs mount --readonly /example s3://apc999/presto-tutorial/example-reason/
Mounted s3://apc999/presto-tutorial/example-reason at /example
```

After mounting this bucket, you are expected to see files inside this directory `/example` after running the following command:

```console
$ alluxio fs ls /example
-r-x------  alluxio-trainingalluxio-training           1712       PERSISTED 10-22-2019 21:46:15:000   0% /example/part-00000-e4f39f27-4060-4bef-a778-c19f65e3a9ea-c000.snappy.parquet
```

Alternatively, you can also browse the directory from Alluxio WebUI at `http://<instance-address>:19999/browse?path=%2Fexample`.

# Task2: Create a simple table and scan it in Presto

This task aims to start a Prest CLI, and create a new table through the CLI.

First, start Preso CLI

```console
$ presto --catalog hive
```

Use "default" schema

```sql
presto> use default;
```

Create an external table using files in Alluxio:

```sql
presto:default> create table reason (
  r_reason_sk bigint,
  r_reason_id varchar,
  r_reason_desc varchar
) WITH (
  external_location = 'alluxio://localhost:19998/example',
  format = 'PARQUET'
);
```

Check the table information:

```sql
presto:default> describe reason;
    Column     |  Type   | Extra | Comment 
---------------+---------+-------+---------
 r_reason_sk   | bigint  |       |         
 r_reason_id   | varchar |       |         
 r_reason_desc | varchar |       |         
(3 rows)

Query 20191105_063932_00009_fbxgy, FINISHED, 1 node
Splits: 19 total, 19 done (100.00%)
0:00 [3 rows, 211B] [14 rows/s, 1.02KB/s]
```

Scan the newly created table using the following command:

```sql
presto:default> select * from reason limit 20;
r_reason_sk |   r_reason_id    |                r_reason_desc                
-------------+------------------+---------------------------------------------
           1 | AAAAAAAABAAAAAAA | Package was damaged                         
           2 | AAAAAAAACAAAAAAA | Stopped working                             
           3 | AAAAAAAADAAAAAAA | Did not get it on time                      
           4 | AAAAAAAAEAAAAAAA | Not the product that was ordred             
           5 | AAAAAAAAFAAAAAAA | Parts missing                               
           6 | AAAAAAAAGAAAAAAA | Does not work with a product that I have    
           7 | AAAAAAAAHAAAAAAA | Gift exchange                               
           8 | AAAAAAAAIAAAAAAA | Did not like the color                      
           9 | AAAAAAAAJAAAAAAA | Did not like the model                      
          10 | AAAAAAAAKAAAAAAA | Did not like the make                       
          11 | AAAAAAAALAAAAAAA | Did not like the warranty                   
          12 | AAAAAAAAMAAAAAAA | No service location in my area              
          13 | AAAAAAAANAAAAAAA | Found a better price in a store             
          14 | AAAAAAAAOAAAAAAA | Found a better extended warranty in a store 
          15 | AAAAAAAAPAAAAAAA | Not working any more                        
          16 | AAAAAAAAABAAAAAA | Did not fit                                 
          17 | AAAAAAAABBAAAAAA | Wrong size                                  
          18 | AAAAAAAACBAAAAAA | Lost my job                                 
          19 | AAAAAAAADBAAAAAA | unauthoized purchase                        
          20 | AAAAAAAAEBAAAAAA | duplicate purchase                          
(20 rows)

Query 20191105_064426_00011_fbxgy, FINISHED, 1 node
Splits: 18 total, 18 done (100.00%)
0:00 [31 rows, 1002B] [105 rows/s, 3.32KB/s]
```

Since thie table is tiny (about 1KB), this query is supposed to complete fast.


# Task3: Run a more complicated query twice

For this task, let us use an existing schema “alluxio” which has a few pre-definied tables on S3:

```sql
presto> use alluxio;
```

Run the a more complicated query in the following:

```sql
presto:alluxio> with ssr as
(select ss_store_sk as store_sk,
     sum(ss_ext_sales_price) as sales
 from store_sales left outer join store_returns on
     (ss_item_sk = sr_item_sk and ss_ticket_number = sr_ticket_number)
     , promotion
 where ss_promo_sk = p_promo_sk
 group by ss_store_sk
 limit 100)
 select channel, sk, sum(sales) as sales
 from (select 'store channel' as channel, store_sk as sk, sales from ssr) x
 group by rollup (channel, sk) limit 100;
```

My reverse engineering on this query:
- What's the total amount of return money based on different channel of promotion
- Table store_sales
```
create external table store_sales
(
    ...
    ss_store_sk               int,
    ss_promo_sk               int,
    ss_ticket_number          int,
    ...
)
```
- Table store_returns
(
    ...
    sr_store_sk               int,
    sr_ticket_number          int,
    ...
)
- Table promotion
(
    ...
    p_promo_sk                int,
    p_channel_dmail           string,
    p_channel_email           string,
    p_channel_catalog         string,
    p_channel_tv              string,
    p_channel_radio           string,
    p_channel_press           string,
    p_channel_event           string,
    p_channel_demo            string,
    p_channel_details         string,
    ...
)
- The query in [`with`](https://teradata.github.io/presto/docs/141t/sql/select.html) defines named relations for use within a query, it sums all sold items with certain promotion channel
- Then next query list all result by (channel, channel-id, sum)
- [`ROLLUP`](https://prestodb.github.io/docs/current/sql/select.html) operator generates all possible subtotals for a given set of columns.

You are expected to see running information like the following indicating the status:

```
Query 20191105_064959_00015_fbxgy, RUNNING, 1 node, 390 splits
0:09 [2.71M rows,   41MB] [ 291K rows/s, 4.39MB/s] [         <=>                              ]

     STAGES   ROWS  ROWS/s  BYTES  BYTES/s  QUEUED    RUN   DONE
0.........R      0       0     0B       0B       0     66      0
  1.......R      0       0     0B       0B       0     32      0
    2.....R   1000     107  15.7K    1.68K       0     32     16
      3...R   248K   26.6K  5.68M     623K       0     48      0
        4.S  1.39M    149K  30.6M    3.28M      31     57     11
        5.S  1.32M    141K  10.4M    1.11M      46      7     43
      6...F   1000       0  3.95K       0B       0      0      1
```

When the query completes, you should press `q` to quite the output and see following summary.

```
Query 20191105_064959_00015_fbxgy, FINISHED, 1 node
Splits: 4,039 total, 4,039 done (100.00%)
6:51 [317M rows, 3.31GB] [770K rows/s, 8.24MB/s]
```

The first run of this query took 6min 51sec with 3.31GB data processed. The query time can vary depending on the network traffic and how many users are querying the same S3 bucket concurrently.

Re-run the same query and compare the duration to finish the queries. The second time is supposed to finish in about 2 minutes because the data is cached in Alluxio now. You can run this queries multiple times and see that having data cached in Alluxio leads to much more consistent performance.

My testing result
```
store channel |   34 |  2.5149941141199903E9
 store channel |   64 |   2.533709799259994E9
 store channel |  190 |   2.522522362150005E9
 store channel |  115 |   2.526016496579988E9
 store channel |  328 |  2.5132327335100026E9
 store channel |   20 |  2.5206218145399942E9
 store channel |  188 |  2.5187698155200047E9
 store channel |  196 |   2.525871516370006E9
 store channel |  400 |   2.518581968089997E9
...
 NULL          | NULL | 2.5244082774835992E11

Query 20191107_211955_00030_amcp7, FINISHED, 1 node
Splits: 4,039 total, 4,039 done (100.00%)
3:00 [317M rows, 3.31GB] [1.76M rows/s, 18.8MB/s]
```

# (Bonus) Task4: Alluxio specific operations

Press `ctrl+d` to exit from the Presto CLI process and try out Alluxio specific operations

## Check Alluxio Total Capacity Usage

Use Alluxio `fsadmin` CLI to checkout basic cluster information including the fraction of space used:

```console
$ alluxio fsadmin report
Alluxio cluster summary: 
    Master Address: localhost:19998
    Web Port: 19999
    Rpc Port: 19998
    Started: 11-05-2019 06:05:27:236
    Uptime: 0 day(s), 0 hour(s), 56 minute(s), and 51 second(s)
    Version: 2.0.1
    Safe Mode: false
    Zookeeper Enabled: false
    Live Workers: 1
    Lost Workers: 0
    Total Capacity: 40.00GB
        Tier: MEM  Size: 40.00GB
    Used Capacity: 16.46GB
        Tier: MEM  Size: 16.46GB
    Free Capacity: 23.54GB
```

## Check Mount Table in Alluxio

Use Alluxio `fs mount` CLI to display the existing mount points

```console
$ alluxio fs mount
s3://apc999/presto-tutorial/example-reason            on  /example  (s3, capacity=-1B, used=-1B, read-only, not shared, properties={})
/opt/alluxio/underFSStorage                           on  /         (local, capacity=15.99GB, used=-1B(0%), not read-only, not shared, properties={})
s3a://alluxio-public-http-ufs/tpcds/scale100-parquet  on  /s3       (s3, capacity=-1B, used=-1B, read-only, not shared, properties={aws.secretKey=******, aws.accessKeyId=******})
```

## Check How Much Data Cached Under a Directory

Alluxio `fs du` CLI calculates how much data is currently cached in `/s3`, 16.46GB out of total 37.36GB files under `/s3` has been loaded to Alluxio as shown:

```console
$ alluxio fs du -s -h /s3
File Size     In Alluxio       Path
37.36GB       16.46GB (44%)    /s3
```

## Free Cache Capacity

Alluxio `fs free` CLI reclaims data from Alluxio space in a given directory:

```console
$ alluxio fs free /s3
/s3 was successfully freed from memory.
$ alluxio fs du -s -h /s3
File Size     In Alluxio       Path
37.36GB       0B (0%)          /s3
```

## Reference
- The experiment information is coming from: https://gist.github.com/apc999/9c73439e66defdb09628f40bf8bdb822

- SQL table information could be found here: https://docs.qubole.com/en/latest/user-guide/hive/example-datasets.html
