
One db will be record in single file on disk, which contains a list of `pages`    
When load page into memory, there is conversion from page -> node, then then node is organized as two dimension tree.  Meta is a spacial page as entrance to all data to certain db.  
When operate on the tree, read or write, an explore process which is abstract as `bucket`, with the help of `cursor` to find specific location


## db
Please go to [boltdb_internal->db](./boltdb_internal.md/#db) for more information.


