
One db will be record in single file on disk, which contains a list of `pages`    
<img src="https://user-images.githubusercontent.com/16873751/98424592-f844f380-2046-11eb-99f4-021003288f93.png" alt="botdb_bucket" width="600"/>
<br/>


When load page into memory, there is conversion from page -> node, then then node is organized as two dimension tree.  Meta is a spacial page as entrance to all data to certain db.  
<img src="https://user-images.githubusercontent.com/16873751/98425036-98e7e300-2048-11eb-9105-3e428d00d256.png" alt="botdb_bucket" width="600"/>
<br/>

When operate on the tree, read or write, an explore process which is abstract as `bucket`, with the help of `cursor` to find specific location
<img src="https://user-images.githubusercontent.com/16873751/98408842-aee4ac00-2026-11eb-9d67-d807ddb8104a.png" alt="botdb_bucket" width="600"/>
<br/>

## db
Please go to [boltdb_internal->db](./boltdb_internal.md/#db) for more information.


