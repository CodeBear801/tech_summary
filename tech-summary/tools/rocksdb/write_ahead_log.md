# Write ahead log

<img src="https://user-images.githubusercontent.com/16873751/97507573-b0351b00-193a-11eb-8e08-df65d43cb89e.png" alt="write_ahead_log" width="600"/>
<br/>
Image from: https://xixymyciqecu.schmidt-grafikdesign.com/write-ahead-log-postgresql-database-32348ng.html

<img src="https://user-images.githubusercontent.com/16873751/96522058-7ae84900-1227-11eb-93ae-a18eec4e1b76.png" alt="write_ahead_log" width="600"/>

more info: https://github.com/google/leveldb/blob/master/doc/log_format.md  

wal对应当前memtable, 当memtable变成immutable-memtable的时候，wal也会被最终保存下来
